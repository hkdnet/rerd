%{
package main
%}
%start program
%union {
	program []*Table
	token Token
	tables []*Table
	table *Table
	columns []*Column
	column *Column
}
%type<program> program
%type<tables> tables
%type<table> table
%type<columns> columns
%type<column> column
%token<token> identifier
%token<token> '{'
%token<token> '}'
%token<token> ';'
%%
program: tables {
	$$ = $1
	yylex.(*Lexer).result = $$
}

;
tables : { $$ = make([]*Table, 0) }
  | table tables {
	$$ = append($2, $1)
  }
;
table : identifier '{' columns '}' {
	$$ = &Table { Name: $1.identifier, Columns: $3 }
}
;
columns : { $$ = make([]*Column, 0) }
  | column columns {
	$$ = append($2, $1)
  }
;
column : identifier ';' {
	$$ = &Column{Name: $1.identifier}
}
;
%%
