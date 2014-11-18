package eval

import (
    "fmt"
    "regexp"
    "strconv"
    "github.com/couchand/alisp/tree"
)

func sum(xs []int64) int64 {
    var s int64 = 0
    for _, x := range xs {
        s += x
    }
    return s
}

func mul(xs []int64) int64 {
    var p int64 = 1
    for _, x := range xs {
        p *= x
    }
    return p
}

var builtins = map[string](func([]int64)int64){
    "+": sum,
    "*": mul,
}

var numberRE = regexp.MustCompile("[0-9]+")

func Eval(t tree.SyntaxTree) int64 {
    //fmt.Printf("Evaluating %v", t)

    if t.IsAtom() {
        //fmt.Printf("Atomizing %v", t)

        if numberRE.MatchString(t.Text) {
            res, err := strconv.ParseInt(t.Text, 10, 64)
            if err != nil {
                msg := fmt.Sprintf("Error converting int: %v", err)
                panic(msg)
            }
            return res
        } else {
            msg := fmt.Sprintf("Unknown value '%s'", t.Text)
            panic(msg)
        }
    } else {
        if len(t.Children) == 0 {
            panic("Illegal unit")
        }
        name := t.Children[0].Text
        fn, ok := builtins[name]
        if !ok {
            msg := fmt.Sprintf("Unknown function '%s'", name)
            panic(msg)
        }

        params := make([]int64, len(t.Children) - 1)
        for i, c := range t.Children {
            if i != 0 {
                params[i - 1] = Eval(c)
            }
        }

        return fn(params)
    }
}
