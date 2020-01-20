package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/jinzhu/inflection"

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

func run(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "cannot read file")
	}
	s := string(b)
	tables := parseTables(s)
	buildReferences(tables)
	printErd(os.Stdout, tables)
	return nil
}

func parseTables(s string) []*Table {
	p := yyNewParser()
	l := &Lexer{src: s}
	p.Parse(l)
	return l.result
}

func buildReferences(tables []*Table) {
	m := make(map[string]*Table)
	for _, table := range tables {
		singularName := inflection.Singular(table.Name)
		m[singularName] = table
	}

	for _, table := range tables {
		for _, col := range table.Columns {
			if endWith(col.Name, "_id") {
				if t, ok := m[col.Name[0:len(col.Name)-3]]; ok {
					col.Reference = t
				} else {
					fmt.Printf("Found reference-ish column but not found table `%s`.`%s`\n", table.Name, col.Name)
				}
			}
		}
	}
	for _, table := range tables {
		for _, col := range table.Columns {
			if col.Reference != nil {
				table.References = append(table.References, col.Reference)
			}
		}
	}
}

func printErd(w io.Writer, tables []*Table) {
	fmt.Fprintln(w, "@startuml \"erd\"")
	for _, table := range tables {
		fmt.Fprintf(w, "entity \"%s\" {\n", table.Name)
		fmt.Fprintln(w, "\t+ id [PK]")
		fmt.Fprintln(w, "\t==")
		for _, column := range table.Columns {
			fmt.Fprintf(w, "\t# %s\n", column.Name)
		}
		fmt.Fprintln(w, "}")
		for _, ref := range table.References {
			fmt.Fprintf(w, "%s --o{ %s\n", ref.Name, table.Name)
		}
	}
	fmt.Fprintln(w, `@enduml`)
}
