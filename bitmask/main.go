package main

import (
	"fmt"
)

const (
	bold       = 1
	italic     = 2
	underlined = 4
)

func format(conf byte) string {
	r := ""
	if (conf & bold) != 0 {
		r = r + "b"
	}
	if (conf & italic) != 0 {
		r = r + "i"
	}
	if (conf & underlined) != 0 {
		r = r + "u"
	}

	return r
}

func main() {
	fmt.Println(format(bold))
	fmt.Println(format(bold | italic | underlined))
	fmt.Println(format(italic | underlined))
}
