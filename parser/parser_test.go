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
