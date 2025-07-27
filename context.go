package gx

import (
	"io"
	"reflect"
)

type Context struct {
	values map[reflect.Type]any
}

func NewContext() *Context {
	return &Context{
		values: make(map[reflect.Type]any),
	}
}

func (c *Context) Push(value any) {
	typ := reflect.TypeOf(value)
	c.values[typ] = value
}

func Use[T any](c *Context) T {
	var zero T
	typ := reflect.TypeOf(zero)
	if value, exists := c.values[typ]; exists {
		return value.(T)
	}

	return zero
}

func SafeUse[T any](c *Context) (T, bool) {
	var zero T
	typ := reflect.TypeOf(zero)
	if value, exists := c.values[typ]; exists {
		return value.(T), true
	}
	return zero, false
}

type componentNode struct {
	fn func(c *Context) Node
}

func (n *componentNode) Render(c *Context, w io.Writer) error {
	return n.fn(c).Render(c, w)
}

func WithContext(fn func(c *Context) Node) Node {
	return &componentNode{fn}
}
