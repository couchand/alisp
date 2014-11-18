package tree

import (
    "fmt"
    "strings"
)

type SyntaxTree struct {
    Children []SyntaxTree
    Text string
}

func Atom(s string) SyntaxTree {
    return SyntaxTree{Text: s}
}

func List(elements []SyntaxTree) SyntaxTree {
    return SyntaxTree{Children: elements}
}

func (t SyntaxTree) String() string {
    if len(t.Text) != 0 {
        return t.Text
    }

    children := make([]string, len(t.Children))
    for i, child := range t.Children {
        children[i] = child.String()
    }
    return fmt.Sprintf("(%s)", strings.Join(children, " "))
}
