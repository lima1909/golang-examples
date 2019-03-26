package main

import "fmt"

func main() {
	mi := map[string]int{}
	mi["foo"]++
	fmt.Printf("foo: %v \n", mi["foo"])

	type mys struct{ name string }
	ms := map[string]mys{}
	ms["foo"] = mys{name: "foo"}
	// compile error
	// ms["foo"].name = "bar"
	fmt.Printf("name: %s\n", ms["foo"].name)

}
