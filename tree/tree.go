package tree

import (
    "fmt"
    "strings"
)

type SyntaxTree struct {
    Children []SyntaxTree
}

func Atom() SyntaxTree {
    return SyntaxTree{}
}

func List(elements []SyntaxTree) SyntaxTree {
    return SyntaxTree{elements}
}

func (t SyntaxTree) String() string {
    children := make([]string, len(t.Children))
    for i, child := range t.Children {
        children[i] = child.String()
    }
    return fmt.Sprintf("(%s)", strings.Join(children, " "))
}
