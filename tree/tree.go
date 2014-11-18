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

func (t SyntaxTree) IsAtom() bool {
    return len(t.Text) != 0
}

func (t SyntaxTree) IsList() bool {
    return !t.IsAtom()
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
