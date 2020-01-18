package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
)

const Usage = "rerd FILENAME"

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, Usage)
		os.Exit(1)
	}
	err := run(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type Lexer struct {
	idx    int
	tokens []string

	result []Table
}

func NewLexer(src string) *Lexer {
	s := strings.ReplaceAll(src, ";", " ;")
	return &Lexer{
		idx:    0,
		tokens: strings.Fields(s),
	}
}

type Token struct {
	identifier string

	lBrace bool
	rBrace bool
}
type Table struct {
	Name    string
	Columns []Column
}
type Column struct {
	Name string
	// type string

	Reference string
}

func (l *Lexer) Lex(lval *yySymType) int {
	if l.idx >= len(l.tokens) {
		return 0
	}
	s := l.tokens[l.idx]
	l.idx++

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

func run(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "cannot read file")
	}
	s := string(b)
	p := yyNewParser()
	l := NewLexer(s)
	i := p.Parse(l)
	fmt.Printf("i = %d\n", i)
	fmt.Printf("%#v\n", l.result)
	return nil
}
