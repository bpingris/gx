package gx_test

import (
	"strings"
	"testing"

	"github.com/bpingris/gx"
)

func TestIfTrue(t *testing.T) {
	ctx := gx.NewContext()
	var buf strings.Builder

	authenticated := true
	node := gx.Div(
		gx.If(authenticated,
			gx.Div(gx.Class("user-panel"), gx.Text("Welcome, User!")),
		),
		gx.Div(gx.Text("Always visible content")),
	)

	node.Render(ctx, &buf)
	result := buf.String()

	if !strings.Contains(result, "Welcome, User!") {
		t.Error("Expected conditional content to be rendered when condition is true")
	}
	if !strings.Contains(result, "Always visible content") {
		t.Error("Expected non-conditional content to always be rendered")
	}
}

func TestIfFalse(t *testing.T) {
	ctx := gx.NewContext()
	var buf strings.Builder

	authenticated := false
	node := gx.Div(
		gx.If(authenticated,
			gx.Div(gx.Class("user-panel"), gx.Text("Welcome, User!")),
		),
		gx.Div(gx.Text("Always visible content")),
	)

	node.Render(ctx, &buf)
	result := buf.String()

	if strings.Contains(result, "Welcome, User!") {
		t.Error("Expected conditional content to be rendered when condition is true")
	}
	if !strings.Contains(result, "Always visible content") {
		t.Error("Expected non-conditional content to always be rendered")
	}
}

func TestIfElseTrue(t *testing.T) {
	ctx := gx.NewContext()
	var buf strings.Builder

	authenticated := true
	node := gx.Div(
		gx.IfElse(authenticated,
			gx.Div(gx.Class("user-panel"), gx.Text("Welcome, User!")),
			gx.Div(gx.Class("login-panel"), gx.Text("Login")),
		),
		gx.Div(gx.Text("Always visible content")),
	)

	node.Render(ctx, &buf)
	result := buf.String()

	if !strings.Contains(result, "Welcome, User!") {
		t.Error("Expected conditional content to be rendered when condition is true")
	}
	if strings.Contains(result, "Login") {
		t.Error("Not expecting else content to be rendered when condition is true")
	}
	if !strings.Contains(result, "Always visible content") {
		t.Error("Expected non-conditional content to always be rendered")
	}
}

func TestIfElseFalse(t *testing.T) {
	ctx := gx.NewContext()
	var buf strings.Builder

	authenticated := false
	node := gx.Div(
		gx.IfElse(authenticated,
			gx.Div(gx.Class("user-panel"), gx.Text("Welcome, User!")),
			gx.Div(gx.Class("login-panel"), gx.Text("Login")),
		),
		gx.Div(gx.Text("Always visible content")),
	)

	node.Render(ctx, &buf)
	result := buf.String()

	if strings.Contains(result, "Welcome, User!") {
		t.Error("Not expecting if content to be rendered when condition is false")
	}
	if !strings.Contains(result, "Login") {
		t.Error("Expected else content to be rendered when condition is false")
	}
	if !strings.Contains(result, "Always visible content") {
		t.Error("Expected non-conditional content to always be rendered")
	}
}

func TestIff(t *testing.T) {
	var buf strings.Builder

	type Foo struct {
		Bar struct {
			Baz string
		}
	}
	var foo *Foo

	node := gx.Div(
		gx.Iff(foo != nil,
			func() gx.Node {
				return gx.Div(gx.Text(foo.Bar.Baz))
			},
		),
		gx.Div(gx.Text("Always visible content")),
	)

	node.Render(gx.NewContext(), &buf)
	result := buf.String()

	if !strings.Contains(result, "Always visible content") {
		t.Error("Expected non-conditional content to always be rendered")
	}
	if strings.Contains(result, "Foo") {
		t.Error("Not expecting conditional content to be rendered when condition is false")
	}
}
