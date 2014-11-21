// lexer

package lexer

import "regexp"
//import "fmt"
import "github.com/couchand/alisp/token"

type Lexer interface {
    GetToken() token.Token
}

type lexer struct {
    s string
    pos int
}

func (l *lexer) GetToken() token.Token {
    //fmt.Println("getting token from '", l.s[l.pos:], "'")

    for {
        if l.pos == len(l.s) {
            //fmt.Println("EOF")
            return token.EOF
        }

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
    if cur == '\'' {
        l.pos = l.pos + 1
        return token.QUOTE
    }

    nextRE := regexp.MustCompile("[ ()']")
    next := nextRE.FindStringIndex(l.s[l.pos:])

    if next == nil {
        x := len(l.s) - l.pos
        next = []int{ x, x }
    }

    //fmt.Println("l.pos = ", l.pos, ", next[0] = ", next[0], ", next[1] = ", next[1])

    oldpos := l.pos
    l.pos = l.pos + next[0]
    newpos := l.pos

    for {
        if l.pos == len(l.s) {
            break
        }
        if m, _ := regexp.MatchString(" ", l.s[l.pos:l.pos + 1]); m {
            l.pos += 1
        } else {
            break
        }
    }

    //fmt.Println("atom!")
    //fmt.Println("next: ", l.s[l.pos : l.pos+1])
    //fmt.Println("remain: ", l.s[l.pos:])

    return token.Atom(l.s[oldpos:newpos])
}

func MakeLexer(str string) Lexer {
    return &lexer{ s: str }
}
