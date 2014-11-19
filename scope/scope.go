package scope

import "fmt"
import "github.com/couchand/alisp/types"

type Scope struct {
    elements map[string]types.Value
    parent *Scope
}

func (s *Scope) Set(key string, val types.Value) types.Value {
    s.elements[key] = val
    return val
}

func (s *Scope) Has(key string) bool {
    _, exists := s.elements[key]
    return exists || s.parent != nil && s.parent.Has(key)
}

func (s *Scope) Get(key string) types.Value {
    val, exists := s.elements[key]
    if exists {
        return val
    }
    if s.parent == nil {
        msg := fmt.Sprintf("Value not in scope: '%s'", key)
        panic(msg)
    }
    return s.parent.Get(key)
}

func makeElements() map[string]types.Value {
    return make(map[string]types.Value)
}

func (s *Scope) ChildScope() *Scope {
    return &Scope{makeElements(), s}
}

func RootScope() *Scope {
    return &Scope{elements: makeElements()}
}
