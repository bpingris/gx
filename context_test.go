package gx_test

import (
	"testing"

	"github.com/bpingris/gx"
)

func TestContextPushAndUse(t *testing.T) {
	ctx := gx.NewContext()

	testString := "hello world"
	ctx.Push(testString)

	result := gx.Use[string](ctx)
	if result != testString {
		t.Errorf("Expected %q, got %q", testString, result)
	}

	testInt := 42
	ctx.Push(testInt)

	resultInt := gx.Use[int](ctx)
	if resultInt != testInt {
		t.Errorf("Expected %d, got %d", testInt, resultInt)
	}

	type User struct {
		Name string
		Age  int
	}
	testUser := User{Name: "John", Age: 30}
	ctx.Push(testUser)

	resultUser := gx.Use[User](ctx)
	if resultUser != testUser {
		t.Errorf("Expected %+v, got %+v", testUser, resultUser)
	}
}

func TestContextUseNonExistent(t *testing.T) {
	ctx := gx.NewContext()

	result := gx.Use[string](ctx)
	if result != "" {
		t.Errorf("Expected empty string for non-existent type, got %q", result)
	}
}

func TestContextOverwriteValue(t *testing.T) {
	ctx := gx.NewContext()

	ctx.Push("first")
	result := gx.Use[string](ctx)
	if result != "first" {
		t.Errorf("Expected 'first', got %q", result)
	}

	ctx.Push("second")
	result = gx.Use[string](ctx)
	if result != "second" {
		t.Errorf("Expected 'second', got %q", result)
	}
}
