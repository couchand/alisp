package lexer

import "testing"
//import "fmt"
import "github.com/couchand/alisp/token"

func TestLex(t *testing.T) {
    l := MakeLexer("(123 456 foobar)")

    if l.GetToken() != token.PAREN_OPEN {
        t.Errorf("Expected open paren")
    }

    if l.GetToken() != token.ATOM {
        t.Errorf("Expected atom")
    }
    if l.GetToken() != token.ATOM {
        t.Errorf("Expected atom")
    }
    if l.GetToken() != token.ATOM {
        t.Errorf("Expected atom")
    }

    if l.GetToken() != token.PAREN_CLOSE {
        t.Errorf("Expected close paren")
    }

    if l.GetToken() != token.EOF {
        t.Errorf("Expected EOF")
    }
    if l.GetToken() != token.EOF {
        t.Errorf("Expected EOF")
    }
    if l.GetToken() != token.EOF {
        t.Errorf("Expected EOF")
    }
}

func TestNested(t *testing.T) {
    l := MakeLexer("((123) 456 foobar)")

    if l.GetToken() != token.PAREN_OPEN {
        t.Errorf("Expected open paren")
    }

    if l.GetToken() != token.PAREN_OPEN {
        t.Errorf("Expected open paren")
    }
    if l.GetToken() != token.ATOM {
        t.Errorf("Expected atom")
    }
    if l.GetToken() != token.PAREN_CLOSE {
        t.Errorf("Expected close paren")
    }

    if l.GetToken() != token.ATOM {
        t.Errorf("Expected atom")
    }
    if l.GetToken() != token.ATOM {
        t.Errorf("Expected atom")
    }

//    fmt.Println("l: ", l)

    if l.GetToken() != token.PAREN_CLOSE {
        t.Errorf("Expected close paren")
    }

    if l.GetToken() != token.EOF {
        t.Errorf("Expected EOF")
    }
    if l.GetToken() != token.EOF {
        t.Errorf("Expected EOF")
    }
    if l.GetToken() != token.EOF {
        t.Errorf("Expected EOF")
    }
}
