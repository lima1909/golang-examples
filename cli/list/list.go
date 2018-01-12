package list

import (
	"fmt"
	"io/ioutil"
	"log"
)

// ListCurrentDir List of the current directory
func ListCurrentDir(maxFiles int) {
	files, err := ioutil.ReadDir("./")
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
