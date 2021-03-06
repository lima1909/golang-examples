package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/lima1909/golang-examples/base"
)

// URL for test
const url = "https://jsonplaceholder.typicode.com/users/2"

// create a client (not use http.Get direct), so I can config my calls
var client = &http.Client{
	Timeout: time.Second * 10,
}

func getURL() string {
	var urlFlag string
	flag.StringVar(&urlFlag, "url", url, "the URL for remote call to get the User")
	flag.Parse()
	return urlFlag
}

// GetJSONBytes get the Json by a User ID
func GetJSONBytes() ([]byte, error) {
	response, err := client.Get(getURL())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

// GetUserByJSONBytes get the User structur by the JSON bytes
func GetUserByJSONBytes(jsonBytes []byte) base.User {
	var user base.User
	json.Unmarshal(jsonBytes, &user)
	return user
}

func main() {
	jsonBytes, err := GetJSONBytes()
	if err != nil {
		log.Fatal(err)
	}
	user := GetUserByJSONBytes(jsonBytes)

	log.Println(user.String())
}
