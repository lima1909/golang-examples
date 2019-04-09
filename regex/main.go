package main

import (
	"fmt"
	"regexp"
)

// doc: https://yourbasic.org/golang/regexp-cheat-sheet/

func main() {
	matched, err := regexp.MatchString(`a.b`, "aaxbb")
	fmt.Println(matched) // true
	fmt.Println(err)     // nil (regexp is valid)
	fmt.Println()

	// ^ (start with) a one char $ (end with) b
	matched, _ = regexp.MatchString(`^a.b$`, "aaxbb")
	fmt.Println(matched) // false
	matched, _ = regexp.MatchString(`^a.b$`, "axb")
	fmt.Println(matched) // true
}
