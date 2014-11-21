package builtin

import (
    "fmt"
    "regexp"
    "strconv"

    "github.com/couchand/alisp/types"
)

var numberRE = regexp.MustCompile("^[0-9]+$")
var nilRE = regexp.MustCompile("^nil$")

func sum(xs []types.Value) types.Value {
    var s int64 = 0
    for _, x := range xs {
        s += x.IntVal()
    }
    return types.Int(s)
}

func mul(xs []types.Value) types.Value {
    var p int64 = 1
    for _, x := range xs {
        p *= x.IntVal()
    }
    return types.Int(p)
}

func cons(ps []types.Value) types.Value {
    if len(ps) != 2 {
        panic("cons expects two parameters only")
    }
    return types.Cons(ps[0], ps[1])
}

func car(ps []types.Value) types.Value {
    if len(ps) != 1 {
        panic("car expects one parameter only")
    }
    return ps[0].CarVal()
}

func cdr(ps []types.Value) types.Value {
    if len(ps) != 1 {
        panic("cdr expects one parameter only")
    }
    return ps[0].CdrVal()
}

func list(ps []types.Value) types.Value {
    if len(ps) == 0 {
        return types.Nil()
    }
    return types.Cons(ps[0], list(ps[1:]))
}

func length(ps []types.Value) types.Value {
    if len(ps) != 1 {
        panic("len expectes one parameter only")
    }
    return types.Int(ps[0].Len())
}

func mapp(ps []types.Value) types.Value {
    if len(ps) % 2 != 0 {
        panic("map expects even number of params")
    }
    pairs := make([]types.Value, len(ps) / 2)
    for i := 0; i < len(ps); i += 2 {
        if !ps[i].IsAtom() {
            panic("map expects atoms for keys")
        }

        pairs[i/2] = types.Cons(ps[i], ps[i+1])
    }
    return list(pairs)
}

func get(ps []types.Value) types.Value {
    if len(ps) != 2 {
        panic("get expects two parameters")
    }
    if !ps[1].IsAtom() {
        panic("get expects atom for key")
    }
    if ps[0].IsNil() {
        return types.Nil()
    }
    if !ps[0].IsProperList() {
        panic("get expects map")
    }
    pair := ps[0].CarVal()
    if pair.CarVal().AtomVal() == ps[1].AtomVal() {
        return pair.CdrVal()
    }
    return get([]types.Value{ ps[0].CdrVal(), ps[1] })
}

func has(ps []types.Value) types.Value {
    if len(ps) != 2 {
        panic("get expects two parameters")
    }
    needle := ps[1]
    if !needle.IsAtom() {
        panic("get expects atom for key")
    }
    mappings := ps[0]
    if mappings.IsNil() {
        //fmt.Println("Empty.")
        return types.Int(0)
    }
    if !mappings.IsProperList() {
        panic("get expects map")
    }
    pair := mappings.CarVal()
    hay := pair.CarVal()
    //fmt.Printf("Checking '%s' against '%s'.", needle, hay)
    if hay.AtomVal() == needle.AtomVal() {
        return types.Int(1)
    }
    return has([]types.Value{ mappings.CdrVal(), needle })
}

func set(ps []types.Value) types.Value {
    if len(ps) != 3 {
        panic("set expects two parameters")
    }
    if !ps[1].IsAtom() {
        panic("set expects atom for key")
    }
    if ps[0].IsNil() {
        return types.Cons(types.Cons(ps[1], ps[2]), types.Nil())
    }
    if !ps[0].IsProperList() {
        panic("set expects map")
    }
    pair := ps[0].CarVal()
    if pair.CarVal().AtomVal() == ps[1].AtomVal() {
        return types.Cons(types.Cons(ps[1], ps[2]), ps[0].CdrVal())
    }
    return types.Cons(pair, set([]types.Value{ ps[0].CdrVal(), ps[1], ps[2] }))
}

func unbound(ps []types.Value) types.Value {
    if len(ps) != 1 {
        panic("unbound expects one parameter")
    }
    if !ps[0].IsCons() {
        panic("unbound expects lambda")
    }
    params := ps[0].CarVal()
    if !params.IsCons() {
        panic("unbound expects lambda")
    }
    u := params.CarVal()
    if !u.IsProperList() {
        panic("unbound expects lambda")
    }
    return u
}

func bound(ps []types.Value) types.Value {
    if len(ps) != 1 {
        panic("bound expects one parameter")
    }
    if !ps[0].IsCons() {
        panic("bound expects lambda")
    }
    params := ps[0].CarVal()
    if !params.IsCons() {
        panic("bound expects lambda")
    }
    b := params.CdrVal()
    if !b.IsProperList() {
        panic("bound expects lambda")
    }
    return b
}

func closure(ps []types.Value) types.Value {
    if len(ps) != 3 {
        panic("closure expects three parameters")
    }
    if !(ps[0].IsNil() || ps[0].IsProperList()) {
        panic("closure expects unbound params")
    }
    if !(ps[1].IsNil() || ps[1].IsProperList()) {
        panic("closure expects bound params")
    }
    if !ps[2].IsProperList() {
        panic("closure expects expression")
    }
    return types.Cons(types.Cons(ps[0], ps[1]), ps[2])
}

func nilQues(ps []types.Value) types.Value {
    if len(ps) != 1 {
        panic("nil? expects one argument")
    }
    if ps[0].IsNil() {
        return types.Int(1)
    }
    return types.Int(0)
}

func consQues(ps []types.Value) types.Value {
    if len(ps) != 1 {
        panic("cons? expects one argument")
    }
    if ps[0].IsCons() {
        return types.Int(1)
    }
    return types.Int(0)
}

func listQues(ps []types.Value) types.Value {
    if len(ps) != 1 {
        panic("list? expects one argument")
    }
    if ps[0].IsProperList() {
        return types.Int(1)
    }
    return types.Int(0)
}

func iff(ps []types.Value) types.Value {
    if len(ps) != 3 {
        panic("if expects three arguments")
    }
    condition := ps[0]
    consequent := ps[1]
    alternate := ps[2]
    if condition.IsNil() || condition.IsInt() && condition.IntVal() == 0 {
        return evalScope([]types.Value{ alternate, types.Nil() })
    } else {
        return evalScope([]types.Value{ consequent, types.Nil() })
    }
}

var Builtins = map[string](func([]types.Value)types.Value){
    "+": sum,
    "*": mul,
    "cons": cons,
    "car": car,
    "cdr": cdr,
    "list": list,
    "len": length,
    "map": mapp,
    "get": get,
    "has": has,
    "set": set,
    "nil?": nilQues,
    "cons?": consQues,
    "list?": listQues,
//    "if": iff,
    "unbound": unbound,
    "bound": bound,
    "closure": closure,
}

func eval(ps []types.Value) types.Value {
    if len(ps) < 1 {
        panic("expected function lambda")
    }
    lambda := ps[0]
    args := ps[1:]
    return executeCall(lambda, args)
}

func addBinding(bindingMap, key, val types.Value) types.Value {
    return types.Cons(types.Cons(key, val), bindingMap)
}

func executeCall(lambda types.Value, arguments []types.Value) types.Value {
    if !lambda.IsCons() {
        panic("Value is not a lambda!")
    }
    values := lambda.CarVal()
    expr := lambda.CdrVal()
    if !values.IsCons() {
        panic("Value is not a lambda!")
    }
    if !expr.IsProperList() {
        panic("Value is not a lambda!")
    }

    unbound := values.CarVal()
    bound := values.CdrVal()

    ul := unbound.Len()
    al := int64(len(arguments))
    if ul != al {
        msg := fmt.Sprintf("Expecting %v arguments, got %v", ul, al)
        panic(msg)
    }

    for i := range arguments {
        //fmt.Printf("Binding %s to %v\n", unbound.CarVal(), arguments[i])
        bound = addBinding(bound, unbound.CarVal(), arguments[i])
        unbound = unbound.CdrVal()
    }

    return evalScope([]types.Value{ expr, bound })
}

func evalScope(ps []types.Value) types.Value {
    if len(ps) != 2 {
        panic("eval-scope expects two parameters")
    }
    expr := ps[0]
    if expr.IsInt() {
        return expr
    }
    if !(expr.IsAtom() || expr.IsProperList()) {
        panic("eval-scope expects expression")
    }
    scope := ps[1]
    if !(scope.IsNil() || scope.IsProperList()) {
        panic("eval-scope expects scope")
    }

    //fmt.Printf("scope: %v\n", scope)
    //fmt.Printf("expr:  %v\n", expr)

    if expr.IsAtom() {
        name := expr.AtomVal()
        if numberRE.MatchString(name) {
            res, err := strconv.ParseInt(name, 10, 64)
            if err != nil {
                msg := fmt.Sprintf("Error converting int: %v", err)
                panic(msg)
            }
            //fmt.Printf("Got literal number %v\n", res)
            return types.Int(res)
        }
        if nilRE.MatchString(name) {
            return types.Nil()
        }
        if has([]types.Value{ scope, expr }).IntVal() == 1 {
            val := get([]types.Value{ scope, expr })
            //fmt.Printf("Got %v for %s\n", val, expr)
            return val
        }
        msg := fmt.Sprintf("Unknown value '%s'", expr.AtomVal())
        panic(msg)
    } else {
        if expr.Len() == 0 {
            panic("illegal unit")
        }

        head := expr.CarVal()
        if head.IsAtom() {
            name := head.AtomVal()

            if name == "lambda" {
                panic("nothing yet")
            }
            if name == "quote" {
                panic("nothing yet")
            }
            if name == "define" {
                panic("nothing yet")
            }

            fn, exists := Builtins[name]

            if exists {
                var params []types.Value
                l := expr.CdrVal().Len()
                if l == 0 {
                    params = make([]types.Value, 0)
                } else {
                    params = make([]types.Value, l)
                    i := 0
                    cursor := expr.CdrVal()
                    for cursor.IsCons() {
                        params[i] = evalScope([]types.Value{cursor.CarVal(), scope})
                        i += 1
                        cursor = cursor.CdrVal()
                    }
                }
                return fn(params)
            } else if has([]types.Value{ scope, head }).IntVal() == 1 {
                lambda := get([]types.Value{ scope, head })

                arguments := expr.CdrVal()
                args := make([]types.Value, arguments.Len())
                i := 0
                for arguments.IsCons() {
                    args[i] = evalScope([]types.Value{arguments.CarVal(), scope})
                    i += 1
                    arguments = arguments.CdrVal()
                }

                return executeCall(lambda, args)
            } else {
                msg := fmt.Sprintf("Unknown function '%s'", name)
                panic(msg)
            }
        }
        msg := fmt.Sprintf("Don't know what to do with %v", head)
        panic(msg)
    }
}

func init() {
    Builtins["eval"] = eval
    Builtins["eval-scope"] = evalScope
}
