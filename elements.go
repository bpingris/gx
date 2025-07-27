package gx

import (
	"fmt"
	"io"
)

type Element struct {
	tag      string
	children []Node
}

func (e *Element) Render(c *Context, w io.Writer) error {
	if _, err := w.Write([]byte("<" + e.tag)); err != nil {
		return err
	}

	attrs := make(map[string]string)
	var contentChildren []Node

	for i := range e.children {
		if attr, ok := e.children[i].(*attrNode); ok {
			attrs[attr.key] = attr.value
		} else {
			contentChildren = append(contentChildren, e.children[i])
		}
	}

	for attr, value := range attrs {
		if _, err := fmt.Fprintf(w, ` %s="%s"`, attr, value); err != nil {
			return err
		}
	}

	if len(contentChildren) == 0 {
		_, err := w.Write([]byte("/>"))
		return err
	}

	if _, err := w.Write([]byte(">")); err != nil {
		return err
	}

	for i := range contentChildren {
		if err := contentChildren[i].Render(c, w); err != nil {
			return err
		}
	}

	_, err := w.Write([]byte("</" + e.tag + ">"))
	return err

}

func Html(children ...Node) Node {
	return &Element{"html", children}
}

func Head(children ...Node) Node {
	return &Element{"head", children}
}

func Body(children ...Node) Node {
	return &Element{"body", children}
}

func Title(children ...Node) Node {
	return &Element{"title", children}
}

func Meta(children ...Node) Node {
	return &Element{"meta", children}
}

func Link(children ...Node) Node {
	return &Element{"link", children}
}

func Script(children ...Node) Node {
	return &Element{"script", children}
}

func Style(children ...Node) Node {
	return &Element{"style", children}
}

func Section(children ...Node) Node {
	return &Element{"section", children}
}

func Article(children ...Node) Node {
	return &Element{"article", children}
}

func Header(children ...Node) Node {
	return &Element{"header", children}
}

func Footer(children ...Node) Node {
	return &Element{"footer", children}
}

func Nav(children ...Node) Node {
	return &Element{"nav", children}
}

func Aside(children ...Node) Node {
	return &Element{"aside", children}
}

func Main(children ...Node) Node {
	return &Element{"main", children}
}

func Div(children ...Node) Node {
	return &Element{"div", children}
}

func Span(children ...Node) Node {
	return &Element{"span", children}
}

func P(children ...Node) Node {
	return &Element{"p", children}
}

func H1(children ...Node) Node {
	return &Element{"h1", children}
}

func H2(children ...Node) Node {
	return &Element{"h2", children}
}

func H3(children ...Node) Node {
	return &Element{"h3", children}
}

func H4(children ...Node) Node {
	return &Element{"h4", children}
}

func H5(children ...Node) Node {
	return &Element{"h5", children}
}

func H6(children ...Node) Node {
	return &Element{"h6", children}
}

func Strong(children ...Node) Node {
	return &Element{"strong", children}
}

func Em(children ...Node) Node {
	return &Element{"em", children}
}

func B(children ...Node) Node {
	return &Element{"b", children}
}

func I(children ...Node) Node {
	return &Element{"i", children}
}

func Small(children ...Node) Node {
	return &Element{"small", children}
}

func Code(children ...Node) Node {
	return &Element{"code", children}
}

func Pre(children ...Node) Node {
	return &Element{"pre", children}
}

func Blockquote(children ...Node) Node {
	return &Element{"blockquote", children}
}

func A(children ...Node) Node {
	return &Element{"a", children}
}

func Ul(children ...Node) Node {
	return &Element{"ul", children}
}

func Ol(children ...Node) Node {
	return &Element{"ol", children}
}

func Li(children ...Node) Node {
	return &Element{"li", children}
}

func Table(children ...Node) Node {
	return &Element{"table", children}
}

func Tr(children ...Node) Node {
	return &Element{"tr", children}
}

func Td(children ...Node) Node {
	return &Element{"td", children}
}

func Th(children ...Node) Node {
	return &Element{"th", children}
}

func Thead(children ...Node) Node {
	return &Element{"thead", children}
}

func Tbody(children ...Node) Node {
	return &Element{"tbody", children}
}

func Tfoot(children ...Node) Node {
	return &Element{"tfoot", children}
}

func Form(children ...Node) Node {
	return &Element{"form", children}
}

func Input(children ...Node) Node {
	return &Element{"input", children}
}

func Button(children ...Node) Node {
	return &Element{"button", children}
}

func Label(children ...Node) Node {
	return &Element{"label", children}
}

func Select(children ...Node) Node {
	return &Element{"select", children}
}

func Option(children ...Node) Node {
	return &Element{"option", children}
}

func Textarea(children ...Node) Node {
	return &Element{"textarea", children}
}

func Img(children ...Node) Node {
	return &Element{"img", children}
}

func Video(children ...Node) Node {
	return &Element{"video", children}
}

func Audio(children ...Node) Node {
	return &Element{"audio", children}
}

func Canvas(children ...Node) Node {
	return &Element{"canvas", children}
}

func Svg(children ...Node) Node {
	return &Element{"svg", children}
}

func Br(children ...Node) Node {
	return &Element{"br", children}
}

func Hr(children ...Node) Node {
	return &Element{"hr", children}
}

func CSSLink(url string) Node {
	return &Element{
		"link",
		[]Node{Attr("rel", "stylesheet"), Attr("href", url)},
	}
}

func JSScript(url string) Node {
	return &Element{
		"script",
		[]Node{Attr("src", url)},
	}
}

func InlineCSS(css string) Node {
	return Style(Type("text/css"), Raw(css))
}

func InlineJS(js string) Node {
	return Script(Type("text/javascript"), Raw(js))
}

func Viewport(content string) Node {
	return Meta(Name("viewport"), Attr("content", content))
}

func ResponsiveViewport() Node {
	return Viewport("width=device-width, initial-scale=1.0")
}

func Charset(charset string) Node {
	return Meta(Name("charset"), Attr("charset", charset))
}

func UTF8Charset() Node {
	return Charset("utf-8")
}

func Favicon(url string) Node {
	return Link(Rel("icon"), Href(url))
}

func Description(content string) Node {
	return Meta(Name("description"), Attr("content", content))
}

func Keywords(content string) Node {
	return Meta(Name("keywords"), Attr("content", content))
}

func Author(name string) Node {
	return Meta(Name("author"), Attr("content", name))
}

type rawNode struct {
	text string
}

func (r *rawNode) Render(c *Context, w io.Writer) error {
	_, err := w.Write([]byte(r.text))
	return err
}

func Raw(text string) Node {
	return &rawNode{text}
}

type textNode struct {
	text string
}

func (t *textNode) Render(c *Context, w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s", t.text)
	return err
}

func Text(text string) Node {
	return &textNode{text}
}

func Textf(format string, args ...any) Node {
	return &textNode{fmt.Sprintf(format, args...)}
}

type fragmentNode struct {
	children []Node
}

func (f *fragmentNode) Render(c *Context, w io.Writer) error {
	for i := range f.children {
		if err := f.children[i].Render(c, w); err != nil {
			return err
		}
	}
	return nil
}

func Fragment(children ...Node) Node {
	return &fragmentNode{children}
}

type doctypeNode struct {
	doctype string
}

func (d *doctypeNode) Render(c *Context, w io.Writer) error {
	_, err := fmt.Fprintf(w, "<!DOCTYPE %s>", d.doctype)
	return err
}

func Doctype(doctype string) Node {
	return &doctypeNode{doctype}
}

func DoctypeHTML5() Node {
	return Doctype("html")
}
