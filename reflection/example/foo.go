package example

// Foo ...
type Foo struct {
	s string
	I int
}

// New ...
func New() Foo {
	return Foo{s: "foo", I: 42}
}
