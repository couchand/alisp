// parser

package parser

import "github.com/couchand/alisp/tree"
import "github.com/couchand/alisp/token"
import "github.com/couchand/alisp/lexer"

func Parse(l *lexer.Lexer) tree.SyntaxTree {
    initial := l.GetToken()

    if initial == token.EOF {
        panic("No input found!")
    }
    if initial == token.PAREN_OPEN {
        return parseKernel(l)
    }
    if initial == token.ATOM {
        return tree.Atom()
    }
    panic("Illegal input found!")
}

func parseKernel(l *lexer.Lexer) tree.SyntaxTree {
    sl := []tree.SyntaxTree{}

    for {
        t := l.GetToken()
        if t == token.EOF {
            panic("Early end of input found, expecting ')'!")
        }
        if t == token.PAREN_OPEN {
            sl = append(sl, parseKernel(l))
        }
        if t == token.PAREN_CLOSE {
            return tree.List(sl)
        }
        if t == token.ATOM {
            sl = append(sl, tree.Atom())
        }
    }
}
