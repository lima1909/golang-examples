package main

import (
	"log"
	"os"
)

func main() {
	log := log.New(os.Stdout, "mylog-entry: ", log.LstdFlags|log.Lshortfile)
	log.Println("the log message")

}
