package rerd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/hkdnet/rerd/parser"
	"github.com/jinzhu/inflection"

	"github.com/pkg/errors"
)

func Run(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "cannot read file")
	}
	s := string(b)
	tables := parser.ParseTables(s)
	buildReferences(tables)
	printErd(os.Stdout, tables)
	return nil
}

func buildReferences(tables []*parser.Table) {
	m := make(map[string]*parser.Table)
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

func printErd(w io.Writer, tables []*parser.Table) {
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
