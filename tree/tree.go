package tree

import (
    "fmt"
    "strings"
    "regexp"
    "strconv"

    "github.com/couchand/alisp/types"
)

var numberRE = regexp.MustCompile("^[0-9]+$")

type SyntaxTree struct {
    Children []SyntaxTree
    Text string
}

func Atom(s string) SyntaxTree {
    return SyntaxTree{Text: s}
}

func Quote(element SyntaxTree) SyntaxTree {
    return SyntaxTree{Children: []SyntaxTree{ element }, Text: "'"}
}

func List(elements []SyntaxTree) SyntaxTree {
    return SyntaxTree{Children: elements}
}

func (t SyntaxTree) IsAtom() bool {
    return !t.IsQuote() && len(t.Text) != 0
}

func (t SyntaxTree) IsQuote() bool {
    return t.Text == "'"
}

func (t SyntaxTree) IsList() bool {
    return !t.IsAtom() && !t.IsQuote()
}

func (t SyntaxTree) String() string {
    if t.IsAtom() {
        return t.Text
    }

    children := make([]string, len(t.Children))
    for i, child := range t.Children {
        children[i] = child.String()
    }
    return fmt.Sprintf("(%s)", strings.Join(children, " "))
}

func list(vs []types.Value) types.Value {
    if len(vs) == 0 {
        return types.Nil()
    }
    return types.Cons(vs[0], list(vs[1:]))
}

func (t SyntaxTree) Val() types.Value {
    if t.IsAtom() {
        //fmt.Printf("Atom: %v\n", t.Text)
        if numberRE.MatchString(t.Text) {
            res, err := strconv.ParseInt(t.Text, 10, 64)
            if err != nil {
                msg := fmt.Sprintf("Error converting int: %v", err)
                panic(msg)
            }
            return types.Int(res)
        }

        return types.Atom(t.Text)
    }
    //fmt.Println("List")
    children := make([]types.Value, len(t.Children))
    for i, child := range t.Children {
        //fmt.Printf("  %v\n", child)
        children[i] = child.Val()
        //fmt.Println("k")
    }
    return list(children)
}
