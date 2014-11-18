package tree

import "testing"

func TestAtom(t *testing.T) {
    a := Atom("a")

    if !a.IsAtom() {
        t.Errorf("Expecting Atom")
    }
    if a.Text != "a" {
        t.Errorf("Expecting text value")
    }
}

func TestList(t *testing.T) {
    a, b, c := Atom("a"), Atom("b"), Atom("c")
    sl := []SyntaxTree{ a, b, c }
    l := List(sl)

    if !l.IsList() {
        t.Errorf("Expecting List")
    }
    if len(l.Children) != 3 {
        t.Errorf("Expecting List")
    }
}

func TestString(t *testing.T) {
    a := Atom("a")
    sl := []SyntaxTree{ a }
    l := List(sl)

    str := l.String()

    if str != "(a)" {
        t.Errorf("Incorrect string representation")
    }
}
