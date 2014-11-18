// lexer

package lexer

import "regexp"
//import "fmt"
import "github.com/couchand/alisp/token"

type Lexer struct {
    s string
    pos int
}

func (l *Lexer) GetToken() token.Token {
    //fmt.Println("getting token from '", l.s[l.pos:], "'")

    if l.pos == len(l.s) {
        //fmt.Println("EOF")
        return token.EOF
    }

    for {
        if m, _ := regexp.MatchString(" ", l.s[l.pos:l.pos + 1]); m {
            l.pos += 1
        } else {
            break
        }
    }

    //fmt.Println("getting")
    //fmt.Println("next: '", l.s[l.pos : l.pos+1], "'")
    //fmt.Println("remain: '", l.s[l.pos:], "'")

    cur := l.s[l.pos]

    //fmt.Println("cur: '", cur, "'")

    if cur == '(' {
        l.pos = l.pos + 1
        //fmt.Println("open paren")
        return token.PAREN_OPEN
    }
    if cur == ')' {
        l.pos = l.pos + 1
        //fmt.Println("close paren")
        return token.PAREN_CLOSE
    }

    nextRE := regexp.MustCompile("[ ()]")
    next := nextRE.FindStringIndex(l.s[l.pos:])

    if next == nil {
        panic("Illegal input")
    }

    //fmt.Println("l.pos = ", l.pos, ", next[0] = ", next[0], ", next[1] = ", next[1])

    l.pos = l.pos + next[0]

    for {
        if m, _ := regexp.MatchString(" ", l.s[l.pos:l.pos + 1]); m {
            l.pos += 1
        } else {
            break
        }
    }

    //fmt.Println("atom!")
    //fmt.Println("next: ", l.s[l.pos : l.pos+1])
    //fmt.Println("remain: ", l.s[l.pos:])

    return token.ATOM
}

func MakeLexer(str string) *Lexer {
    return &Lexer{ s: str }
}
