package list

import (
	"fmt"
	"io/ioutil"
	"log"
)

// example for an Enum
const (
	file int = iota + 1
	directory
)

// Dir List of all files and directories from a start directory
func Dir(startDirectory string, maxFiles int) {
	files, err := ioutil.ReadDir(startDirectory)
	if err != nil {
		log.Fatal("Error by list files ", err)
	}
	for i, f := range files {
		if i < maxFiles {
			fmt.Println(getFilePrefix(f.IsDir()), f.Name())
		}
	}
}

func getFilePrefix(isDir bool) string {
	if isDir == true {
		return "d "
	}
	return "- "
}
