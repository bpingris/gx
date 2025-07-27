package gx_test

import (
	"strings"
	"testing"

	"github.com/bpingris/gx"
)

func normalizeHTML(html string) string {
	html = strings.ReplaceAll(html, "\n", "")
	html = strings.ReplaceAll(html, "\t", "")
	return html
}

func TestTextNode(t *testing.T) {
	ctx := gx.NewContext()
	var buf strings.Builder

	text := gx.Text("hello world")

	text.Render(ctx, &buf)
	if buf.String() != "hello world" {
		t.Errorf("expected 'hello world', get '%q'", buf.String())
	}
}

func TestRawNode(t *testing.T) {
	ctx := gx.NewContext()
	var buf strings.Builder

	raw := gx.Raw(`<div><span>hello world</span><script>alert('hello world');</script></div>`)

	raw.Render(ctx, &buf)
	if buf.String() != "<div><span>hello world</span><script>alert('hello world');</script></div>" {
		t.Errorf("expected 'hello world', get '%q'", buf.String())
	}
}

func TestNestedElements(t *testing.T) {
	ctx := gx.NewContext()
	var buf strings.Builder

	div := gx.Div(
		gx.Class("foo"),
		gx.ID("foo"),
		gx.Div(
			gx.Nav(
				gx.Ul(
					gx.Li(gx.Text("foo")),
					gx.Li(gx.Text("bar")),
					gx.Li(gx.Text("baz")),
				),
			),
		),
	)

	div.Render(ctx, &buf)
	expected := `<div class="foo" id="foo">
		<div>
			<nav>
				<ul>
					<li>foo</li>
					<li>bar</li>
					<li>baz</li>
				</ul>
			</nav>
		</div>
	</div>`
	if buf.String() != normalizeHTML(expected) {
		t.Errorf("expected '%q', got '%q'", expected, buf.String())
	}
}
