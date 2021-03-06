// parser

package parser

import "github.com/couchand/alisp/tree"
import "github.com/couchand/alisp/token"
import "github.com/couchand/alisp/lexer"

func quote(v tree.SyntaxTree) tree.SyntaxTree {
    return tree.Quote(v)
}

func Parse(l lexer.Lexer) tree.SyntaxTree {
    initial := l.GetToken()

    if initial == token.EOF {
        return tree.SyntaxTree{}
    }
    if initial == token.PAREN_OPEN {
        return parseKernel(l)
    }
    if initial == token.QUOTE {
        return quote(Parse(l))
    }
    if initial.IsAtom() {
        return tree.Atom(initial.Text)
    }
    panic("Illegal input found!")
}

func parseKernel(l lexer.Lexer) tree.SyntaxTree {
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
        if t == token.QUOTE {
            sl = append(sl, quote(Parse(l)))
        }
        if t.IsAtom() {
            sl = append(sl, tree.Atom(t.Text))
        }
        //panic("Illegal input found!")
    }
}
