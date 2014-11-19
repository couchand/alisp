package scope

import "github.com/couchand/alisp/types"

import "testing"

func TestAdd(t *testing.T) {
    s := RootScope()

    setval := s.Set("foo", types.Int(1337))

    if !s.Has("foo") {
        t.Errorf("Expected scope to have key 'foo'")
    }

    if setval.IntVal() != 1337 {
        t.Errorf("Expected set to return set value")
    }

    v := s.Get("foo")
    if v.IntVal() != 1337 {
        t.Errorf("Expected 1337, got %v", v.IntVal())
    }
}

func TestChild(t *testing.T) {
    parent := RootScope()

    parent.Set("foo", types.Int(42))

    child := parent.ChildScope()

    if child.Get("foo").IntVal() != 42 {
        t.Errorf("Expected child to get parent scope")
    }

    child.Set("foo", types.Int(1337))

    if parent.Get("foo").IntVal() == 1337 {
        t.Errorf("Expected parent value to be shadowed")
    }
}
