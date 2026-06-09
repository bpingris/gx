package gx

import (
	"bytes"
	"fmt"
	"html"
	"io"
	"strings"
)

var voidElements = map[string]bool{
	"area":   true,
	"base":   true,
	"br":     true,
	"col":    true,
	"embed":  true,
	"hr":     true,
	"img":    true,
	"input":  true,
	"link":   true,
	"meta":   true,
	"param":  true,
	"source": true,
	"track":  true,
	"wbr":    true,
}

type Element struct {
	tag     string
	attrs   map[string]string
	content []Node
}

func newElement(tag string, children []Node) *Element {
	attrs := make(map[string]string)
	var content []Node
	for i := range children {
		if attr, ok := children[i].(*attrNode); ok {
			attrs[attr.key] = attr.value
		} else {
			content = append(content, children[i])
		}
	}
	return &Element{tag, attrs, content}
}

func (e *Element) Render(c *Context, w io.Writer) error {
	if _, err := io.WriteString(w, "<"); err != nil {
		return err
	}
	if _, err := io.WriteString(w, e.tag); err != nil {
		return err
	}

	for attr, value := range e.attrs {
		if _, err := fmt.Fprintf(w, ` %s="%s"`, html.EscapeString(attr), html.EscapeString(value)); err != nil {
			return err
		}
	}

	if voidElements[e.tag] {
		_, err := io.WriteString(w, ">")
		return err
	}

	if len(e.content) == 0 {
		if _, err := io.WriteString(w, "></"); err != nil {
			return err
		}
		if _, err := io.WriteString(w, e.tag); err != nil {
			return err
		}
		_, err := io.WriteString(w, ">")
		return err
	}

	if _, err := io.WriteString(w, ">"); err != nil {
		return err
	}

	for i := range e.content {
		if err := e.content[i].Render(c, w); err != nil {
			return err
		}
	}

	if _, err := io.WriteString(w, "</"); err != nil {
		return err
	}
	if _, err := io.WriteString(w, e.tag); err != nil {
		return err
	}
	_, err := io.WriteString(w, ">")
	return err

}

func Html(children ...Node) Node {
	return newElement("html", children)
}

func Head(children ...Node) Node {
	return newElement("head", children)
}

func Body(children ...Node) Node {
	return newElement("body", children)
}

func Title(children ...Node) Node {
	return newElement("title", children)
}

func Meta(children ...Node) Node {
	return newElement("meta", children)
}

func Link(children ...Node) Node {
	return newElement("link", children)
}

func Script(children ...Node) Node {
	return newElement("script", children)
}

func Style(children ...Node) Node {
	return newElement("style", children)
}

func Section(children ...Node) Node {
	return newElement("section", children)
}

func Article(children ...Node) Node {
	return newElement("article", children)
}

func Header(children ...Node) Node {
	return newElement("header", children)
}

func Footer(children ...Node) Node {
	return newElement("footer", children)
}

func Nav(children ...Node) Node {
	return newElement("nav", children)
}

func Aside(children ...Node) Node {
	return newElement("aside", children)
}

func Main(children ...Node) Node {
	return newElement("main", children)
}

func Div(children ...Node) Node {
	return newElement("div", children)
}

func Span(children ...Node) Node {
	return newElement("span", children)
}

func P(children ...Node) Node {
	return newElement("p", children)
}

func H1(children ...Node) Node {
	return newElement("h1", children)
}

func H2(children ...Node) Node {
	return newElement("h2", children)
}

func H3(children ...Node) Node {
	return newElement("h3", children)
}

func H4(children ...Node) Node {
	return newElement("h4", children)
}

func H5(children ...Node) Node {
	return newElement("h5", children)
}

func H6(children ...Node) Node {
	return newElement("h6", children)
}

func Strong(children ...Node) Node {
	return newElement("strong", children)
}

func Em(children ...Node) Node {
	return newElement("em", children)
}

func B(children ...Node) Node {
	return newElement("b", children)
}

func I(children ...Node) Node {
	return newElement("i", children)
}

func Small(children ...Node) Node {
	return newElement("small", children)
}

func Code(children ...Node) Node {
	return newElement("code", children)
}

func Pre(children ...Node) Node {
	return newElement("pre", children)
}

func Blockquote(children ...Node) Node {
	return newElement("blockquote", children)
}

func A(children ...Node) Node {
	return newElement("a", children)
}

func Ul(children ...Node) Node {
	return newElement("ul", children)
}

func Ol(children ...Node) Node {
	return newElement("ol", children)
}

func Li(children ...Node) Node {
	return newElement("li", children)
}

func Table(children ...Node) Node {
	return newElement("table", children)
}

func Tr(children ...Node) Node {
	return newElement("tr", children)
}

func Td(children ...Node) Node {
	return newElement("td", children)
}

func Th(children ...Node) Node {
	return newElement("th", children)
}

func Thead(children ...Node) Node {
	return newElement("thead", children)
}

func Tbody(children ...Node) Node {
	return newElement("tbody", children)
}

func Tfoot(children ...Node) Node {
	return newElement("tfoot", children)
}

func Form(children ...Node) Node {
	return newElement("form", children)
}

func Input(children ...Node) Node {
	return newElement("input", children)
}

func Fieldset(children ...Node) Node {
	return newElement("fieldset", children)
}

func Button(children ...Node) Node {
	return newElement("button", children)
}

func Label(children ...Node) Node {
	return newElement("label", children)
}

func Select(children ...Node) Node {
	return newElement("select", children)
}

func Option(children ...Node) Node {
	return newElement("option", children)
}

func Textarea(children ...Node) Node {
	return newElement("textarea", children)
}

func Img(children ...Node) Node {
	return newElement("img", children)
}

func Video(children ...Node) Node {
	return newElement("video", children)
}

func Audio(children ...Node) Node {
	return newElement("audio", children)
}

func Canvas(children ...Node) Node {
	return newElement("canvas", children)
}

func Svg(children ...Node) Node {
	return newElement("svg", children)
}

func Br() Node {
	return newElement("br", nil)
}

func Hr() Node {
	return newElement("hr", nil)
}

func CSSLink(url string) Node {
	return newElement("link", []Node{Attr("rel", "stylesheet"), Attr("href", url)})
}

func JSScript(url string) Node {
	return newElement("script", []Node{Attr("src", url)})
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
	return Meta(Attr("charset", charset))
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
	_, err := io.WriteString(w, r.text)
	return err
}

func Raw(text string) Node {
	return &rawNode{text}
}

type textNode struct {
	text string
}

func (t *textNode) Render(c *Context, w io.Writer) error {
	_, err := io.WriteString(w, html.EscapeString(t.text))
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

const slotPlaceholder = "<!-- slot -->"

type slotNode struct{}

func (s *slotNode) Render(c *Context, w io.Writer) error {
	_, err := io.WriteString(w, slotPlaceholder)
	return err
}

func Slot() Node {
	return &slotNode{}
}

type CompiledTemplate struct {
	beforeSlot string
	afterSlot  string
}

func (t *CompiledTemplate) Render(children ...Node) Node {
	return &compiledNode{
		template: t,
		children: children,
	}
}

type compiledNode struct {
	template *CompiledTemplate
	children []Node
}

func (cn *compiledNode) Render(c *Context, w io.Writer) error {
	if _, err := io.WriteString(w, cn.template.beforeSlot); err != nil {
		return err
	}

	for i := range cn.children {
		if err := cn.children[i].Render(c, w); err != nil {
			return err
		}
	}

	_, err := io.WriteString(w, cn.template.afterSlot)
	return err
}

func Compile(template Node) (*CompiledTemplate, error) {
	ctx := NewContext()
	var buf bytes.Buffer

	if err := template.Render(ctx, &buf); err != nil {
		return nil, err
	}

	html := buf.String()

	parts := strings.Split(html, slotPlaceholder)
	if len(parts) != 2 {
		return &CompiledTemplate{
			beforeSlot: html,
			afterSlot:  "",
		}, nil
	}

	return &CompiledTemplate{
		beforeSlot: parts[0],
		afterSlot:  parts[1],
	}, nil
}
