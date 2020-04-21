package main

import (
	"fmt"
	"unsafe"
)

// doc: https://www.pixelstech.net/article/1584241521-Understand-unsafe-in-GoLang
//

// Programmer ...
type Programmer struct {
	name     string
	language string
}

func main() {
	p := Programmer{"stefno", "go"}
	fmt.Println(p)

	name := (*string)(unsafe.Pointer(&p))
	*name = "qcrao"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	*lang = "Golang"

	fmt.Println(p)
}
