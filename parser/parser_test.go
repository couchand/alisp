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

    if len(r.Children) != 1 {
        t.Errorf("Expecting quoted atom, got %v", r)
    }

    c := r.Children[0]
    if len(c.Children) != 3 {
        t.Errorf("Expecting quoted value, got %v", c)
    }

    if c.Children[0].Text != "1" {
        t.Errorf("Expecting atom to have quoted value")
    }
    if c.Children[1].Text != "2" {
        t.Errorf("Expecting atom to have quoted value")
    }
    if c.Children[2].Text != "3" {
        t.Errorf("Expecting atom to have quoted value")
    }
}

//func TestQuote2(t *testing.T) {
//    l := lexer.MakeLexer("(list 'foo 'bar)")
//    r := Parse(l)
//
//    if len(r.Children) != 3 {
//        t.Errorf("Expecting parsed tree")
//    }
//
//    fooAtom := r.Children[1]
//    if !fooAtom.IsQuote() {
//        t.Errorf("Expecting quoted value, got %v", fooAtom)
//    }
//    if fooAtom.Children[0].Text != "'foo" {
//        t.Errorf("Expecting atom to have quoted value, got '%s'", fooAtom.Text)
//    }
//
//    barAtom := r.Children[2]
//    if !barAtom.IsQuote() {
//        t.Errorf("Expecting quoted value, got %v", barAtom)
//    }
//    if barAtom.Text != "'bar" {
//        t.Errorf("Expecting atom to have quoted value, got '%s'", barAtom.Text)
//    }
//}
