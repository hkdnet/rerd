package parser

import "fmt"

type Token struct {
	identifier string

	lBrace bool
	rBrace bool
}
type LexResult struct {
	val string

	eof bool
	// TODO: col, line
}
type Table struct {
	Name       string
	Columns    []*Column
	References []*Table
}
type Column struct {
	Name string
	// type string

	Reference *Table
}

func (l *Lexer) Lex(lval *yySymType) int {
	res := l.lex()
	fmt.Printf("DEBUG: lex is finished: %#v\n", res)
	if res.eof {
		return 0
	}
	s := res.val
	if s == "{" {
		return int('{')
	} else if s == "}" {
		return int('}')
	} else if s == ";" {
		return int(';')
	} else {
		lval.token = Token{identifier: s}
		return identifier
	}
}
func (l *Lexer) Error(s string) {
	panic(s)
}

func ParseTables(s string) []*Table {
	p := yyNewParser()
	l := &Lexer{src: s}
	l.next()
	p.Parse(l)
	return l.result
}
