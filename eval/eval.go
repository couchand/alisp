package eval

import (
    "fmt"
    "regexp"
    "strconv"
    "github.com/couchand/alisp/tree"
    "github.com/couchand/alisp/types"
    "github.com/couchand/alisp/builtin"
)

var numberRE = regexp.MustCompile("^[0-9]+$")
var nilRE = regexp.MustCompile("^nil$")

func Eval(t tree.SyntaxTree) types.Value {
    //fmt.Printf("Evaluating %v", t)

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
        } else {
            msg := fmt.Sprintf("Unknown value '%s'", t.Text)
            panic(msg)
        }
    } else {
        if len(t.Children) == 0 {
            panic("Illegal unit")
        }
        name := t.Children[0].Text

        if name == "quote" {
            if len(t.Children) != 2 {
                panic("Expected a single argument to quote")
            }
            arg := t.Children[1]
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

        fn, ok := builtin.Builtins[name]
        if !ok {
            msg := fmt.Sprintf("Unknown function '%s'", name)
            panic(msg)
        }

        params := make([]types.Value, len(t.Children) - 1)
        for i, c := range t.Children {
            if i != 0 {
                params[i - 1] = Eval(c)
            }
        }

        return fn(params)
    }
}
