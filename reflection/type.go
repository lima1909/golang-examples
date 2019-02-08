package main

import (
	"fmt"
	"reflect"
)

type foo struct {
	name string `mytag:"foo"`
}

func printTypeExamples() {

	fmt.Printf("String/Kind: %v/%v\n", reflect.TypeOf(""), reflect.TypeOf("").Kind())
	a := "a"
	fmt.Printf("String Pointer/Kind: %v/%v\n", reflect.TypeOf(&a), reflect.TypeOf(&a).Kind())
	fmt.Printf("int/Kind: %v/%v\n\n", reflect.TypeOf(0), reflect.TypeOf(0).Kind())

	// slice
	sl := []string{"a", "b"}
	slt := reflect.TypeOf(sl)
	fmt.Printf("String-Slice/Kind/Elm: %v/%v/%v\n\n", slt, slt.Kind(), slt.Elem())
	// !! PANIC - Slice - Type has no Len
	// slt.Len()

	// array
	arr := [2]string{"a", "b"}
	fmt.Printf("String-Array %v/%v\n", reflect.TypeOf(arr), reflect.TypeOf(arr).Kind())
	fmt.Printf("Elem/Len %v/%v\n\n", reflect.TypeOf(arr).Elem(), reflect.TypeOf(arr).Len())

	// map
	m := map[string]int32{"a": 1, "b": 2}
	mt := reflect.TypeOf(m)
	fmt.Printf("Map %v/%v\n", mt, mt.Kind())
	fmt.Printf("Key/Value(Elem) %v - %v\n\n", mt.Key(), mt.Elem())

	// struct
	f := foo{"blub"}
	ft := reflect.TypeOf(f)
	fmt.Printf("Sruct Type/Kind %v/%v\n", ft, ft.Kind())
	fmt.Printf("Fields/Name %v (%v)/%v\n", ft.NumField(), ft.Field(0), ft.Name())
	fmt.Printf("Tag %v/%v \n", ft.Field(0).Tag, ft.Field(0).Tag.Get("mytag"))

	fmt.Println("----------------------")
}
