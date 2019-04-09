package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pw := "my-password"
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), 10)
	if err != nil {
		log.Fatalf("err by generate hash from password: %v", err)
	}
	log.Printf("hash: %v", string(hash))

	falsePw := "dummy"
	err = bcrypt.CompareHashAndPassword(hash, []byte(falsePw))
	log.Printf("%v", err)

	err = bcrypt.CompareHashAndPassword(hash, []byte(pw))
	log.Printf("password is correct. err is %v", err)

}
