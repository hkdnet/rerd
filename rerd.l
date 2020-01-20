%{
package parser

import (
	"fmt"
)

type Lexer struct {
	c    byte
	src  string
	in   []byte
	un   []byte
	mark int

	result []*Table // for parser
}
func (l*Lexer) lex() LexResult {
%}

%yyn l.next()
%yyc l.c
%yym fmt.Printf("\tstate accepts: %q\n", l.in); l.mark = len(l.in)

%%
	l.in = l.in[:0]
	l.mark = -1

\0
	return LexResult{eof: true}

[a-zA-Z0-9_-]+
	return LexResult{val: string(l.in) }

[ \t\n\r]+          // Ignore whitespace

%%
	if l.mark >= 0 {
		if len(l.in) > l.mark {
			l.unget(l.c)
			for i := len(l.in)-1; i >= l.mark; i-- {
				l.unget(l.in[i])
			}
			l.next()
		}
		l.in = l.in[:l.mark]
		goto yyAction // Hook: Execute the semantic action of the last matched rule.
	}

    s := string(l.c)
    l.next()
	return LexResult{val: s}
}

func (l*Lexer) next() {
	if len(l.un) != 0 {
		l.c = l.un[len(l.un)-1]
		l.un = l.un[:len(l.un)-1]
		return
	}

	l.in = append(l.in, l.c)
	if len(l.src) == 0 {
		l.c = 0
		return
	}

	l.c = l.src[0]
	fmt.Printf("\tnext: %q\n", l.c)
	l.src = l.src[1:]
}

func (l*Lexer) unget(b byte) {
	l.un = append(l.un, b)
}
