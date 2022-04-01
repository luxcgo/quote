package main

import (
	"fmt"
)

type Fooer interface {
	Foo()
}

type A struct {
	f Fooer
}

func (a *A) Bar() {
	a.f.Foo()
}

func NewA(f Fooer) *A {
	return &A{f: f}
}

type B struct {
}

func (b *B) Foo() {
	fmt.Println("B.Foo()")
}

type C struct {
}

func (c *C) Foo() {
	fmt.Println("C.Foo()")
}

func main() {
	a := NewA(new(B))
	a.Bar()

	a.f = &C{}
	a.Bar()
}
