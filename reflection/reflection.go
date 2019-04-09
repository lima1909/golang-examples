package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type timer struct {
	dur time.Duration
}

func (t *timer) createStopTimeFunc(fn interface{}) interface{} {
	ft := reflect.TypeOf(fn)
	if reflect.Func != ft.Kind() {
		panic(fmt.Sprintf("f ist not a func, %v", ft))
	}

	fv := reflect.ValueOf(fn)

	wrapperFunc := reflect.MakeFunc(ft, func(args []reflect.Value) []reflect.Value {
		start := time.Now()
		ret := fv.Call(args)
		time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
		t.dur = time.Now().Sub(start)
		return ret
	})

	return wrapperFunc.Interface()
}

func lenS(s string) int {
	return len(s)
}

func ret42() int {
	return 42
}

func main() {
	printTypeExamples()

	printValueExamples()
	showExecExample()

	t := new(timer)

	f := t.createStopTimeFunc(lenS).(func(string) int)
	fmt.Printf("Result: %v\n", f("Buhhh"))
	fmt.Printf("Time: %v\n", t.dur)

	v42 := t.createStopTimeFunc(ret42).(func() int)
	fmt.Printf("Result: %v\n", v42())
	fmt.Printf("Time: %v\n", t.dur)

	fmt.Println("-------")
	printPublicValues()
	printPrivateValues()
}
