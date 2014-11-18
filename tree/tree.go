package tree

type SyntaxTree struct {
    Children []SyntaxTree
}

func Atom() SyntaxTree {
    return SyntaxTree{}
}

func List(elements []SyntaxTree) SyntaxTree {
    return SyntaxTree{elements}
}
