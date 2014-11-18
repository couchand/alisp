// token

package token

type Token struct {
    int
    Text string
}

var EOF         = Token{0, ""}
var PAREN_OPEN  = Token{1, "("}
var PAREN_CLOSE = Token{2, ")"}

func (t Token) IsAtom() bool {
    return t.int == 3
}

func Atom(s string) Token {
    return Token{3, s}
}
