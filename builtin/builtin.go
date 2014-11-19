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
    "set": set,
}
