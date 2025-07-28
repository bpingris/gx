# GX

HTML generation library for Go.

## Installation

```bash
go get github.com/bpingris/gx
```

## Quick Start

### Basic HTML Generation

```go
package main

import (
    "os"
    "github.com/bpingris/gx"
)

func main() {
    page := gx.Html(
        gx.Head(
            gx.Title(gx.Text("My Website")),
            gx.UTF8Charset(),
            gx.ResponsiveViewport(),
        ),
        gx.Body(
            gx.H1(gx.Text("Hello, World!")),
            gx.P(
                gx.Class("intro"),
                gx.Text("Welcome to my website built with GX!"),
            ),
        ),
    )

    ctx := gx.NewContext()
    page.Render(ctx, os.Stdout)
}
```

### Using Context

```go
type User struct {
    Name  string
    Email string
}

func UserProfile() gx.Node {
    return gx.WithContext(func(c *gx.Context) gx.Node {
        user := gx.Use[User](c)
        
        return gx.Div(
            gx.Class("user-profile"),
            gx.H2(gx.Text(user.Name)),
            gx.P(gx.Text(user.Email)),
        )
    })
}

func main() {
    user := User{Name: "John Doe", Email: "john@example.com"}
    
    page := gx.Div(
        gx.Provide(user, UserProfile()),
    )

    ctx := gx.NewContext()
    page.Render(ctx, os.Stdout)
}
```

## Template Compilation

For better performance with large layouts, you can pre-compile templates:

### Basic Layout Compilation

```go
// Define your layout with a slot for dynamic content
func Layout(children ...gx.Node) gx.Node {
    return gx.Html(
        gx.Head(
            gx.Title(gx.Text("My App")),
            gx.UTF8Charset(),
            gx.ResponsiveViewport(),
            gx.CSSLink("/styles.css"),
        ),
        gx.Body(
            gx.Header(
                gx.H1(gx.Text("My Website")),
                gx.Nav(
                    gx.A(gx.Href("/"), gx.Text("Home")),
                    gx.A(gx.Href("/about"), gx.Text("About")),
                ),
            ),
            gx.Main(
                gx.Class("container"),
                children, // Dynamic content goes here
            ),
            gx.Footer(
                gx.P(gx.Text("© 2025 My Company")),
            ),
        ),
    )
}

// Compile the layout once (typically at app startup)
var compiledLayout, _ = gx.Compile(Layout(gx.Slot()))

// Use the compiled layout in your pages
func HomePage() gx.Node {
    return compiledLayout.Render(
        gx.H2(gx.Text("Welcome Home")),
        gx.P(gx.Text("This is the home page content.")),
        gx.Div(
            gx.Class("home-content"),
            gx.P(gx.Text("Latest news and updates...")),
        ),
    )
}

func AboutPage() gx.Node {
    return compiledLayout.Render(
        gx.H2(gx.Text("About Us")),
        gx.P(gx.Text("We are a great company.")),
        gx.Ul(
            gx.Li(gx.Text("Founded in 2020")),
            gx.Li(gx.Text("100+ happy customers")),
        ),
    )
}
```

## API Reference

### Core Types

- `Node` - Interface for all renderable elements
- `Context` - Passes data through the component tree
- `CompiledTemplate` - Pre-rendered template with slots

### HTML Elements

All standard HTML elements are supported:

```go
// Structure
gx.Html(), gx.Head(), gx.Body(), gx.Div(), gx.Span()

// Text content  
gx.H1(), gx.H2(), gx.P(), gx.Text(), gx.Raw()

// Forms
gx.Form(), gx.Input(), gx.Button(), gx.Select(), gx.Option()

// Lists
gx.Ul(), gx.Ol(), gx.Li()

// Tables
gx.Table(), gx.Tr(), gx.Td(), gx.Th()

// Media
gx.Img(), gx.Video(), gx.Audio()
```

### Attributes

```go
// Common attributes
gx.Class("my-class")
gx.ID("my-id")  
gx.Href("/path")
gx.Src("/image.jpg")

// Form attributes
gx.Type("text")
gx.Name("username")
gx.Placeholder("Enter username")

// Data attributes
gx.Data("toggle", "modal")

// Boolean attributes
gx.Disabled()
gx.Required()
gx.Checked()

// Custom attributes
gx.Attr("custom-attr", "value")
```

### Utility Functions

```go
// Text and Raw content
gx.Text("Hello")                    // Escaped text
gx.Textf("Hello %s", name)         // Formatted text
gx.Raw("<script>...</script>")     // Unescaped HTML

// Fragments
gx.Fragment(node1, node2, node3)   // Group nodes without wrapper

// Document helpers
gx.DoctypeHTML5()                  // <!DOCTYPE html>
gx.UTF8Charset()                   // <meta charset="utf-8">
gx.ResponsiveViewport()            // Responsive viewport meta
gx.CSSLink("/styles.css")          // CSS link
gx.JSScript("/script.js")          // JavaScript script
```

### Context Functions

```go
// Create context
ctx := gx.NewContext()

// Push values
ctx.Push(user)
ctx.Push("some string")

// Use values in components
user := gx.Use[User](ctx)
str := gx.Use[string](ctx)

// Safe usage (with existence check)
user, exists := gx.SafeUse[User](ctx)

// Provide values to children
gx.Provide(user, 
    gx.Div(/* children can access user via context */),
)

// Create context-aware components
gx.WithContext(func(c *gx.Context) gx.Node {
    user := gx.Use[User](c)
    return gx.Text(user.Name)
})
```

### Template Compilation

```go
// Create a slot for dynamic content
gx.Slot()

// Compile a template
compiled, err := gx.Compile(templateWithSlot)

// Use compiled template
page := compiled.Render(dynamicContent...)
```

## Examples

### Simple Blog Post

```go
type Post struct {
    Title   string
    Content string
    Author  string
}

func BlogPost() gx.Node {
    return gx.WithContext(func(c *gx.Context) gx.Node {
        post := gx.Use[Post](c)

        return gx.Article(
            gx.Class("blog-post"),
            gx.Header(
                gx.H1(gx.Text(post.Title)),
                gx.P(
                    gx.Class("author"),
                    gx.Text("By "), 
                    gx.Strong(gx.Text(post.Author)),
                ),
            ),
            gx.Div(
                gx.Class("content"),
                gx.Raw(post.Content), // Assuming HTML content
            ),
        )
    })
}
```

### Popular Alternatives

- [gomponents](https://github.com/maragudk/gomponents) (**might be better**)
