package types

import (
    "fmt"
    "strings"
)

type typeTag int

const typeNil     = 0
const typeInt     = 1
const typeAtom    = 2
const typeCons    = 3
const typeLambda  = 4

type Value struct {
    tag typeTag
    val interface{}
}

type consCell struct {
    head Value
    tail Value
}

func Nil() Value {
    return Value{tag: typeNil}
}

func (v Value) IsNil() bool {
    return v.tag == typeNil
}

func Int(val int64) Value {
    return Value{tag: typeInt, val: val}
}

func (v Value) IsInt() bool {
    return v.tag == typeInt
}

func (v Value) IntVal() int64 {
    if !v.IsInt() {
        msg := fmt.Sprintf("Value not an int: %v", v)
        panic(msg)
    }
    return v.val.(int64)
}

func Atom(val string) Value {
    return Value{tag: typeAtom, val: val}
}

func (v Value) IsAtom() bool {
    return v.tag == typeAtom
}

func (v Value) AtomVal() string {
    if !v.IsAtom() {
        msg := fmt.Sprintf("Value is not an atom: %v", v)
        panic(msg)
    }
    return v.val.(string)
}

func Cons(left, right Value) Value {
    return Value{tag: typeCons, val: consCell{left, right}}
}

func (v Value) IsCons() bool {
    return v.tag == typeCons
}

func (v Value) IsProperList() bool {
    if !v.IsCons() {
        return false
    }
    tail := v.CdrVal()
    return tail.IsNil() || tail.IsProperList()
}

func (v Value) Len() int64 {
    if !v.IsCons() {
        return 0
    }
    return 1 + v.CdrVal().Len()
}

func (v Value) CarVal() Value {
    if !v.IsCons() {
        msg := fmt.Sprintf("Value not a cons: %v", v)
        panic(msg)
    }
    return v.val.(consCell).head;
}

func (v Value) CdrVal() Value {
    if !v.IsCons() {
        msg := fmt.Sprintf("Value not a cons: %v", v)
        panic(msg)
    }
    return v.val.(consCell).tail;
}

func (v Value) String() string {
    if v.IsNil() {
        return "nil"
    }
    if v.IsInt() {
        return fmt.Sprintf("%v", v.val)
    }
    if v.IsAtom() {
        return fmt.Sprintf("%v", v.val)
    }
    if v.IsProperList() {
        l := v.Len()
        sl := make([]string, l)
        i := 0
        c := v
        for {
            if !c.IsCons() {
                break
            }
            sl[i] = c.CarVal().String()
            i += 1
            c = c.CdrVal()
        }
        return fmt.Sprintf("(%s)", strings.Join(sl, " "))
    }
    if v.IsCons() {
        return fmt.Sprintf("(%v . %v)", v.CarVal().String(), v.CdrVal().String())
    }
    panic("don't know what to do")
}
