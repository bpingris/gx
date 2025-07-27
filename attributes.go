package gx

import (
	"fmt"
	"io"
)

type attrNode struct {
	key   string
	value string
}

func (a *attrNode) Render(c *Context, w io.Writer) error {
	_, err := fmt.Fprintf(w, ` %s="%s"`, a.key, a.value)
	return err
}

func Type(t string) Node {
	return &attrNode{"type", t}
}

func Rel(url string) Node {
	return &attrNode{"rel", url}
}

func Href(url string) Node {
	return &attrNode{"href", url}
}

func Class(class string) Node {
	return &attrNode{"class", class}
}

func ID(id string) Node {
	return &attrNode{"id", id}
}

func Attr(attr, value string) Node {
	return &attrNode{attr, value}
}

func Name(name string) Node {
	return &attrNode{"name", name}
}

func Src(url string) Node {
	return &attrNode{"src", url}
}

func Placeholder(text string) Node {
	return &attrNode{"placeholder", text}
}

func Min(value string) Node {
	return &attrNode{"min", value}
}

func Max(value string) Node {
	return &attrNode{"max", value}
}

func Data(key, value string) Node {
	return &attrNode{"data-" + key, value}
}

func For(id string) Node {
	return &attrNode{"for", id}
}

func Action(url string) Node {
	return &attrNode{"action", url}
}

func Method(method string) Node {
	return &attrNode{"method", method}
}

func Target(target string) Node {
	return &attrNode{"target", target}
}

func Title_(title string) Node {
	return &attrNode{"title", title}
}

func Style_(style string) Node {
	return &attrNode{"style", style}
}

func Lang(lang string) Node {
	return &attrNode{"lang", lang}
}

func Dir(direction string) Node {
	return &attrNode{"dir", direction}
}

func TabIndex(index string) Node {
	return &attrNode{"tabindex", index}
}

func Role(role string) Node {
	return &attrNode{"role", role}
}

func AriaLabel(label string) Node {
	return &attrNode{"aria-label", label}
}

func AriaHidden() Node {
	return &attrNode{"aria-hidden", "true"}
}

func Disabled() Node {
	return &attrNode{"disabled", "disabled"}
}

func Required() Node {
	return &attrNode{"required", "required"}
}

func Readonly() Node {
	return &attrNode{"readonly", "readonly"}
}

func Multiple() Node {
	return &attrNode{"multiple", "multiple"}
}

func Checked() Node {
	return &attrNode{"checked", "checked"}
}

func Autofocus() Node {
	return &attrNode{"autofocus", "autofocus"}
}

func Hidden() Node {
	return &attrNode{"hidden", "hidden"}
}

func Selected() Node {
	return &attrNode{"selected", "selected"}
}
