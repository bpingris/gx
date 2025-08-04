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

var (
	_ Node = (*ifNode)(nil)
)
