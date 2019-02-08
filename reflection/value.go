package main

import (
	"fmt"
	"reflect"
)

// a good description can you find her:
// https://medium.com/capital-one-tech/learning-to-use-go-reflection-822a0aed74b7

func printValueExamples() {
	// if you use a pointer, than you get the Value: reflect.ValueOf(..).Elem().Interface()
	// if you use a var, than you get the Value: reflect.ValueOf(..).Interface() (without Elem())
	a := "a"
	fmt.Printf("%v - %T \n", reflect.ValueOf(a).Interface(), reflect.ValueOf(a).Interface())
	// you need a pointer for set a new value
	ap := reflect.ValueOf(&a)
	fmt.Printf("Type: %v Kind: %v Pointer: %v Value %v\n", ap.Type(), ap.Type().Kind(), ap.Interface(), ap.Elem().Interface())
	ap.Elem().SetString("aa")
	fmt.Printf("New Value: %v\n", ap.Elem().Interface())
	// or an other way to set a new value
	ap.Elem().Set(reflect.ValueOf("bb"))
	fmt.Printf("New-new Value: %v - %T\n", ap.Elem().Interface(), ap.Elem().Interface())

	fmt.Println("----------------------")
}

// Exec is a dummy struct
type Exec struct{}

// Do print a string (must be public: capital letter!!!)
func (e Exec) Do(s string) {
	fmt.Println("print exec string:", s)
}

// call a method per reflection
func showExecExample() {
	et := reflect.TypeOf(Exec{})
	ev := reflect.New(et)

	m := ev.MethodByName("Do")
	if reflect.Func == m.Type().Kind() {
		m.Call([]reflect.Value{reflect.ValueOf("Yehhh ...")})
	}

	fmt.Println("----------------------")
}
