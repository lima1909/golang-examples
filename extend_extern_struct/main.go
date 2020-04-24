package main

import (
	"fmt"

	"github.com/lima1909/golang-examples/extend_extern_struct/foo"
)

// Stringer ...
type Stringer interface{ ToString() string }

// MyFoo type of foo.F
type MyFoo foo.F

// ToString ...
func (mf MyFoo) ToString() string {
	return mf.S
}

func printType(f *foo.F) {
	mf := (*MyFoo)(f)
	fmt.Println(mf.ToString())
}

func printGeneric(i interface{}) {
	if o, ok := i.(Stringer); ok {
		fmt.Println(o.ToString())
	} else {
		fmt.Println("nö")
	}
}

func convert(i interface{}) interface{} {
	switch o := i.(type) {
	case foo.F:
		return (MyFoo)(o)
	case *foo.F:
		return (*MyFoo)(o)
	}
	return i
}

func createFoo() *foo.F {
	mf := &MyFoo{S: "My Foo"}
	return (*foo.F)(mf)
}

func main() {
	printType(createFoo())
	printGeneric(convert(createFoo()))
}
