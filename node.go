package gx

import "io"

type Node interface {
	Render(c *Context, w io.Writer) error
}
