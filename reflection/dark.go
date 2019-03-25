package main

import (
	"fmt"
	"reflect"

	_ "unsafe"

	"github.com/lima1909/golang-examples/reflection/example"
)

func printPublicValues() {
	f := example.New()
	v := reflect.ValueOf(&f).Elem()
	fmt.Printf("- Public Fields: %#v \n", v)

	for i := 0; i < v.NumField(); i++ {
		// v.Field(i).Interface() throws PANIC by private/unexported field (field: s)
		fmt.Printf("%v: %#v (%T)\n", v.Type().Field(i).Name, v.Field(i), v.Field(i))
	}
}

// https://t.co/YNZ3xofNlv
// https://docs.google.com/presentation/d/1OCwIPVffnKsFWuIW-jH2EpihQwODGH7ljsTSoSo14Qw/edit#slide=id.p

//go:linkname valueInterface reflect.valueInterface
func valueInterface(v reflect.Value, safe bool) interface{}

func printPrivateValues() {
	f := example.New()
	v := reflect.ValueOf(&f).Elem()
	fmt.Printf("- Private Fields: %#v \n", v)

	for i := 0; i < v.NumField(); i++ {
		intf := valueInterface(v.Field(i), false)
		fmt.Printf("%v: %#v (%T)\n", v.Type().Field(i).Name, intf, intf)

	}
}
