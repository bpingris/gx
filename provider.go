package gx

import "io"

type provideNode struct {
	value    any
	children []Node
}

func (p *provideNode) Render(c *Context, w io.Writer) error {
	restore := c.PushScoped(p.value)
	defer restore()

	for i := range p.children {
		if err := p.children[i].Render(c, w); err != nil {
			return err
		}
	}
	return nil
}

func Provide(value any, children ...Node) Node {
	return &provideNode{value, children}
}
