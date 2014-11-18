package builtin

import "github.com/couchand/alisp/types"

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

var Builtins = map[string](func([]types.Value)types.Value){
    "+": sum,
    "*": mul,
    "cons": cons,
    "car": car,
    "cdr": cdr,
    "list": list,
}
