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

func contains(haystack tree.SyntaxTree, needle tree.SyntaxTree) bool {
    for _, candidate := range haystack.Children {
        if candidate.Text == needle.Text {
            return true
        }
    }
    return false
}

func list(vs []types.Value) types.Value {
    if len(vs) == 0 {
        return types.Nil()
    }
    return types.Cons(vs[0], list(vs[1:]))
}

//            defs := buildBindingMap(t.Children[1], t.Children[2], s)
func buildBindingMap(s *scope.Scope, params tree.SyntaxTree, expression tree.SyntaxTree) types.Value {
    if !params.IsList() {
        panic("expected param list to lambda")
    }
    if expression.IsAtom() {
        bindings := buildBindingMapKernel(s, params, expression)
        return types.Cons(bindings[0], types.Nil())
    } else {
        bindings := buildBindingMapKernel(s, params, expression)
        return list(bindings)
    }
}

func buildBindingMapKernel(s *scope.Scope, params tree.SyntaxTree, expression tree.SyntaxTree) []types.Value {
    if expression.IsAtom() {
        name := expression.Text
        if numberRE.MatchString(name) {
            return []types.Value{}
        } else if nilRE.MatchString(name) {
            return []types.Value{}
        }
        if name == "lambda" || name == "define" || name == "quote" || name == "if" {
            return []types.Value{}
        }
        _, builtin := builtin.Builtins[name]
        if !builtin && !contains(params, expression) {
            key := types.Atom(name)
            val := s.Get(name)
            return []types.Value{ types.Cons(key, val) }
        } else {
            return []types.Value{}
        }
    } else {
        mappings := make([]types.Value, 0, len(expression.Children))
        for _, child := range expression.Children {
            ms := buildBindingMapKernel(s, params, child)
            for _, mapping := range ms {
                mappings = append(mappings, mapping)
            }
        }
        return mappings
    }
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

        if !t.Children[0].IsAtom() {
            panic("not yet")
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

        if name == "lambda" {
            if len(t.Children) != 3 {
                panic("Expected exactly two arguments to lambda")
            }
            args := quote(t.Children[1])
            expr := quote(t.Children[2])
            defs := buildBindingMap(s, t.Children[1], t.Children[2])
            return types.Cons(types.Cons(args, defs), expr)
        }

        fn, ok := builtin.Builtins[name]
        if !ok {
            if s.Has(name) {
                lambda := s.Get(name)

                params := make([]types.Value, len(t.Children))
                params[0] = lambda
                for i, c := range t.Children[1:] {
                    params[i + 1] = EvalScope(c, s.ChildScope())
                }

                return builtin.Builtins["eval"](params)
            }

            msg := fmt.Sprintf("Unknown function '%s'", name)
            panic(msg)
        }

        params := make([]types.Value, len(t.Children) - 1)
        for i, c := range t.Children[1:] {
            params[i] = EvalScope(c, s.ChildScope())
        }

        return fn(params)
    }
}
