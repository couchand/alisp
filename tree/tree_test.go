package tree

import "testing"

func TestAtom(t *testing.T) {
    a := Atom()

    if len(a.Children) != 0 {
        t.Errorf("Expecting Atom")
    }
}

func TestList(t *testing.T) {
    a, b, c := Atom(), Atom(), Atom()
    sl := []SyntaxTree{ a, b, c }
    l := List(sl)

    if len(l.Children) != 3 {
        t.Errorf("Expecting List")
    }
}
