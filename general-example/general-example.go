package main

import (
	"errors"
	"fmt"
)

// Weekday example for an Enum
type Weekday int

// FileTypeEnum Enum with File = 1 and Dir = 2
const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func doSomethingWithError(returnError bool) error {
	if returnError {
		return errors.New("i'm a error")
	}
	return nil
}

func typeSwitch(i interface{}) {
	switch value := i.(type) {
	case string:
		fmt.Printf("you are a string with value: %v\n", value)
	case int:
		fmt.Printf("you are a int with value: %v\n", value)
	default:
		fmt.Printf("you are an unknown type %T with value: %v\n", value, value)

	}
}

func string2byteArray(s string) []byte {
	return []byte(s)
}

func byteArray2string(b []byte) string {
	return string(b)
}

func main() {
	err := doSomethingWithError(true)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	typeSwitch("a string")
	typeSwitch(Monday)

	fmt.Println("byteArray2string and back: ", byteArray2string(string2byteArray("I'm a string")))
}
