package parser

import "testing"
import "github.com/couchand/alisp/lexer"

func TestParse(t *testing.T) {
    l := lexer.MakeLexer("(1 2 3)")
    r := Parse(l)

    if len(r.Children) != 3 {
        t.Errorf("Expecting parsed tree")
    }

    for _, c := range r.Children {
        if len(c.Children) != 0 {
            t.Errorf("Expecting atom")
        }
    }
}

func TestQuote(t *testing.T) {
    l := lexer.MakeLexer("'(1 2 3)")
    r := Parse(l)

    if len(r.Children) != 0 {
        t.Errorf("Expecting quoted atom")
    }

    if r.Text != "'(1 2 3)" {
        t.Errorf("Expecting atom to have quoted value, got '%s'", r.Text)
    }
}

func TestQuote2(t *testing.T) {
    l := lexer.MakeLexer("(list 'foo 'bar)")
    r := Parse(l)

    if len(r.Children) != 3 {
        t.Errorf("Expecting parsed tree")
    }

    if r.Children[1].Text != "'foo" {
        t.Errorf("Expecting atom to have quoted value, got '%s'", r.Text)
    }

    if r.Children[2].Text != "'bar" {
        t.Errorf("Expecting atom to have quoted value, got '%s'", r.Text)
    }
}
