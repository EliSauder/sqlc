package mssql

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/sqlc-dev/teesql/parser"

	"github.com/sqlc-dev/sqlc/internal/source"
	"github.com/sqlc-dev/sqlc/internal/sql/ast"
	"github.com/sqlc-dev/sqlc/internal/sql/sqlerr"
)

func NewParser() *Parser {
	return &Parser{}
}

type Parser struct{}

func (p *Parser) Parse(r io.Reader) ([]ast.Statement, error) {
	blob, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	script, err := parser.Parse(ctx, bytes.NewReader(blob))
	if err != nil {
		return nil, err
	}

	toByte := byteOffsets(blob)

	var stmts []ast.Statement
	loc := 0
	for _, batch := range script.Batches {
		for _, stmt := range batch.Statements {
			start := loc
			if frag, ok := stmt.(fragmented); ok && frag.Frag().HasSpan() {
				if s := toByte(frag.Frag().StartOffset); s > start {
					start = s
				}
			}
			end := statementEnd(blob, start)

			converter := &cc{toByte: toByte}
			out := converter.convert(stmt)
			if _, ok := out.(*ast.TODO); ok {
				// A schema holds statements sqlc has no use for, which are
				// skipped. A named query sqlc cannot analyze is an error,
				// rather than a query missing from the generated code.
				if namedQuery.Match(blob[loc:end]) {
					kind := strings.TrimSuffix(strings.TrimPrefix(fmt.Sprintf("%T", stmt), "*ast."), "Statement")
					return nil, &sqlerr.Error{
						Message:  fmt.Sprintf("%s statements are not supported", kind),
						Location: start,
					}
				}
				for _, ddl := range converter.blockSchemaStatements(stmt) {
					stmts = append(stmts, ast.Statement{
						Raw: &ast.RawStmt{Stmt: ddl, StmtLocation: loc, StmtLen: end - loc},
					})
				}
				loc = end
				continue
			}

			if namedQuery.Match(blob[loc:end]) {
				if err := checkPlaceholders(blob, start, end, converter.namedParams); err != nil {
					return nil, err
				}
			}

			stmts = append(stmts, ast.Statement{
				Raw: &ast.RawStmt{
					Stmt:         out,
					StmtLocation: loc,
					StmtLen:      end - loc,
				},
			})
			loc = end
		}
	}

	return stmts, nil
}

// namedQuery matches the comment that names a query, in either comment
// style.
var namedQuery = regexp.MustCompile(`(?m)^\s*(--|/\*)\s*name:\s*\S+`)

// checkPlaceholders reports a placeholder in a query's text that did not
// become one of its parameters, which a construct the converter does not
// follow drops, leaving generated code that passes the query too few
// arguments.
func checkPlaceholders(blob []byte, start, end int, params map[string]namedParam) error {
	for _, p := range placeholders(blob[start:end]) {
		if _, ok := params[strings.ToLower(p.name)]; !ok {
			return &sqlerr.Error{
				Message:  fmt.Sprintf("placeholder @%s is not supported where it is used", p.name),
				Location: start + p.offset,
			}
		}
	}
	return nil
}

type placeholder struct {
	name   string
	offset int
}

// placeholders are the @name placeholders in sql, leaving out @@ system
// variables and anything inside a string, a quoted or bracketed
// identifier or a comment.
func placeholders(sql []byte) []placeholder {
	var out []placeholder
	skipTo := func(i int, close string) int {
		if j := bytes.Index(sql[i:], []byte(close)); j >= 0 {
			return i + j + len(close)
		}
		return len(sql)
	}
	for i := 0; i < len(sql); {
		switch c := sql[i]; {
		case c == '\'' || c == '"':
			// A doubled quote inside is an escaped one, which the scan
			// steps over as the end of one string and the start of the
			// next.
			i = skipTo(i+1, string(c))
		case c == '[':
			i = skipTo(i+1, "]")
		case c == '-' && i+1 < len(sql) && sql[i+1] == '-':
			i = skipTo(i, "\n")
		case c == '/' && i+1 < len(sql) && sql[i+1] == '*':
			i = skipTo(i+2, "*/")
		case c == '@':
			j := i + 1
			for j < len(sql) && isNameByte(sql[j]) {
				j++
			}
			if i+1 < len(sql) && sql[i+1] == '@' {
				i = j
				continue
			}
			if j > i+1 {
				out = append(out, placeholder{name: string(sql[i+1 : j]), offset: i})
			}
			i = j
		default:
			i++
		}
	}
	return out
}

func isNameByte(b byte) bool {
	return b == '_' || b == '#' || b == '$' || b == '@' ||
		'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || '0' <= b && b <= '9' || b >= utf8.RuneSelf
}

// byteOffsets returns a function mapping a UTF-16 code-unit offset — how
// teesql records source spans, mirroring ScriptDom — to a byte offset in
// blob. For ASCII input the two are identical and the identity is returned.
func byteOffsets(blob []byte) func(int) int {
	ascii := true
	for _, b := range blob {
		if b >= utf8.RuneSelf {
			ascii = false
			break
		}
	}
	if ascii {
		return func(off int) int { return off }
	}
	// Walk the blob once, recording the byte index at which each UTF-16
	// offset begins.
	byteAt := make([]int, 0, len(blob)+1)
	for i := 0; i < len(blob); {
		r, size := utf8.DecodeRune(blob[i:])
		units := 1
		if r > 0xFFFF {
			units = 2
		}
		for u := 0; u < units; u++ {
			byteAt = append(byteAt, i)
		}
		i += size
	}
	byteAt = append(byteAt, len(blob))
	return func(off int) int {
		if off < 0 {
			return 0
		}
		if off >= len(byteAt) {
			return len(blob)
		}
		return byteAt[off]
	}
}

// statementEnd scans from start for the semicolon ending the statement,
// skipping string literals, quoted and bracketed identifiers, and comments.
func statementEnd(blob []byte, start int) int {
	for i := start; i < len(blob); i++ {
		switch blob[i] {
		case '\'', '"':
			i = skipQuoted(blob, i)
		case '[':
			i = skipBracketed(blob, i)
		case '-':
			if i+1 < len(blob) && blob[i+1] == '-' {
				i = skipLineComment(blob, i)
			}
		case '/':
			if i+1 < len(blob) && blob[i+1] == '*' {
				i = skipBlockComment(blob, i)
			}
		case ';':
			return i + 1
		}
	}
	return len(blob)
}

// skipQuoted skips a string literal or quoted identifier; T-SQL escapes the
// quote character by doubling it.
func skipQuoted(blob []byte, i int) int {
	q := blob[i]
	for j := i + 1; j < len(blob); j++ {
		if blob[j] == q {
			if j+1 < len(blob) && blob[j+1] == q {
				j++
				continue
			}
			return j
		}
	}
	return len(blob) - 1
}

// skipBracketed skips a [bracketed identifier]; a closing bracket is escaped
// by doubling it.
func skipBracketed(blob []byte, i int) int {
	for j := i + 1; j < len(blob); j++ {
		if blob[j] == ']' {
			if j+1 < len(blob) && blob[j+1] == ']' {
				j++
				continue
			}
			return j
		}
	}
	return len(blob) - 1
}

func skipLineComment(blob []byte, i int) int {
	for j := i; j < len(blob); j++ {
		if blob[j] == '\n' {
			return j
		}
	}
	return len(blob) - 1
}

func skipBlockComment(blob []byte, i int) int {
	for j := i + 2; j < len(blob); j++ {
		if blob[j] == '*' && j+1 < len(blob) && blob[j+1] == '/' {
			return j + 1
		}
	}
	return len(blob) - 1
}

// https://learn.microsoft.com/en-us/sql/t-sql/language-elements/comment-transact-sql
func (p *Parser) CommentSyntax() source.CommentSyntax {
	return source.CommentSyntax{
		Dash:      true, // -- comment
		SlashStar: true, // /* comment */
	}
}
