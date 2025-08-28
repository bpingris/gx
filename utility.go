package gx

import "io"

type ifNode struct {
	condition  bool
	trueChild  Node
	falseChild Node
}

// Render implements Node.
func (in *ifNode) Render(c *Context, w io.Writer) error {
	if in.condition && in.trueChild != nil {
		return in.trueChild.Render(c, w)
	}
	if !in.condition && in.falseChild != nil {
		return in.falseChild.Render(c, w)
	}
	return nil
}

func If(condition bool, child Node) Node {
	return &ifNode{condition, child, nil}
}

func IfElse(condition bool, trueChild, falseChild Node) Node {
	return &ifNode{condition, trueChild, falseChild}
}

func Iff(condition bool, fn func() Node) Node {
	var child Node
	if condition {
		child = fn()
	}
	return &ifNode{condition, child, nil}
}

var (
	_ Node = (*ifNode)(nil)
)

type mapNode[T any] struct {
	items []T
	fn    func(item T, index int) Node
}

func (m *mapNode[T]) Render(c *Context, w io.Writer) error {
	for i, item := range m.items {
		if err := m.fn(item, i).Render(c, w); err != nil {
			return err
		}
	}
	return nil
}

func Map[T any](items []T, fn func(item T, index int) Node) Node { return &mapNode[T]{items, fn} }

var _ Node = (*mapNode[any])(nil)

type repeatNode struct {
	times int
	child Node
}

func (r *repeatNode) Render(c *Context, w io.Writer) error {
	for range r.times {
		if err := r.child.Render(c, w); err != nil {
			return err
		}
	}
	return nil
}

func Repeat(times int, child Node) Node { return &repeatNode{times, child} }

var _ Node = (*repeatNode)(nil)

// type switchNode[T comparable] struct {
// 	value    T
// 	cases    map[T]Node
// 	fallback Node
// }
//
// func (s *switchNode[T]) Render(c *Context, w io.Writer) error {
// 	var child Node
// 	if cchild, ok := s.cases[s.value]; ok {
// 		child = cchild
// 	}
// }
