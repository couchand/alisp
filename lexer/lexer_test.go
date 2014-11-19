package lexer

import "testing"
//import "fmt"
import "github.com/couchand/alisp/token"

func TestSpaces(t *testing.T) {
    l := MakeLexer("  (  foo  bar   baz    )   ")

    if l.GetToken() != token.PAREN_OPEN {
        t.Errorf("Expected open paren")
    }

    s := l.GetToken()
    if s.Text != "foo" {
        t.Errorf("Expected 'foo', got '%s'", s.Text)
    }
    s = l.GetToken()
    if s.Text != "bar" {
        t.Errorf("Expected 'bar', got '%s'", s.Text)
    }
    s = l.GetToken()
    if s.Text != "baz" {
        t.Errorf("Expected 'baz', got '%s'", s.Text)
    }

    if l.GetToken() != token.PAREN_CLOSE {
        t.Errorf("Expected close paren")
    }

    if l.GetToken() != token.EOF {
        t.Errorf("Expected EOF")
    }
}

func TestQuote(t *testing.T) {
    l := MakeLexer("'bar'(1 2 3)")

    if l.GetToken() != token.QUOTE {
        t.Errorf("Expected quote")
    }

    s := l.GetToken()
    if s.Text != "bar" {
        t.Errorf("Expected 'bar', got '%s'", s.Text)
    }

    if l.GetToken() != token.QUOTE {
        t.Errorf("Expected quote")
    }

    if l.GetToken() != token.PAREN_OPEN {
        t.Errorf("Expected open paren")
    }
}

func TestLex(t *testing.T) {
    l := MakeLexer("(123 456 foobar)")

    if l.GetToken() != token.PAREN_OPEN {
        t.Errorf("Expected open paren")
    }

    if !l.GetToken().IsAtom() {
        t.Errorf("Expected atom")
    }
    if !l.GetToken().IsAtom() {
        t.Errorf("Expected atom")
    }
    if !l.GetToken().IsAtom() {
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
    if !l.GetToken().IsAtom() {
        t.Errorf("Expected atom")
    }
    if l.GetToken() != token.PAREN_CLOSE {
        t.Errorf("Expected close paren")
    }

    if !l.GetToken().IsAtom() {
        t.Errorf("Expected atom")
    }
    if !l.GetToken().IsAtom() {
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
