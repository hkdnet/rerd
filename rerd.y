%{
package main
%}
%start program
%token a A
%union {
i int32;
}

%%
program:
  | A
;
