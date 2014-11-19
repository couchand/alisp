package eval

import (
    "fmt"
    "regexp"
    "strconv"
    "github.com/couchand/alisp/tree"
    "github.com/couchand/alisp/types"
    "github.com/couchand/alisp/builtin"
    "github.com/couchand/alisp/scope"
)

var atomRE = regexp.MustCompile("^'")
var numberRE = regexp.MustCompile("^[0-9]+$")
var nilRE = regexp.MustCompile("^nil$")

func quote(arg tree.SyntaxTree) types.Value {
    //fmt.Printf("Quoting %v\n", arg)
    if arg.IsList() && len(arg.Children) == 0 {
        //fmt.Println("Found an empty list")
        return types.Nil()
    }
    //fmt.Println("Calling valify")
    v := arg.Val()
    //fmt.Printf("Quoted %v\n", v)
    return v
}

func Eval(t tree.SyntaxTree) types.Value {
    //fmt.Printf("Evaluating %v", t)

    return EvalScope(t, scope.RootScope())
}

func EvalScope(t tree.SyntaxTree, s *scope.Scope) types.Value {
    //fmt.Println("%v", s)

    if t.IsAtom() {
        //fmt.Printf("Atomizing %v", t)

        if numberRE.MatchString(t.Text) {
            res, err := strconv.ParseInt(t.Text, 10, 64)
            if err != nil {
                msg := fmt.Sprintf("Error converting int: %v", err)
                panic(msg)
            }
            return types.Int(res)
        } else if nilRE.MatchString(t.Text) {
            return types.Nil()
        } else if atomRE.MatchString(t.Text) {
            return types.Atom(t.Text[1:])
        } else if s.Has(t.Text) {
            return s.Get(t.Text)
        } else {
            msg := fmt.Sprintf("Unknown value '%s'", t.Text)
            panic(msg)
        }
    } else if t.IsQuote() {
        return quote(t.Children[0])
    } else {
        if len(t.Children) == 0 {
            panic("Illegal unit")
        }
        name := t.Children[0].Text

        if name == "quote" {
            if len(t.Children) != 2 {
                panic("Expected a single argument to quote")
            }
            return quote(t.Children[1])
        }

        if name == "define" {
            if len(t.Children) != 3 {
                panic("Expected exactly two arguments to define")
            }
            key := t.Children[1]
            value := t.Children[2]
            if !key.IsAtom() {
                panic("Expected name for define")
            }
            return s.Set(key.Text, EvalScope(value, s))
        }

        fn, ok := builtin.Builtins[name]
        if !ok {
            msg := fmt.Sprintf("Unknown function '%s'", name)
            panic(msg)
        }

        params := make([]types.Value, len(t.Children) - 1)
        for i, c := range t.Children {
            if i != 0 {
                params[i - 1] = EvalScope(c, s.ChildScope())
            }
        }

        return fn(params)
    }
}
