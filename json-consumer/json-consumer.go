package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// User is the result structur
type User struct {
	UserID   int    `json:"id"`
	Name     string `json:"name"`
	UserName string
	EMail    string
}

func (user User) String() string {
	return fmt.Sprintf("User [%s%d%s%s%s%s]",
		"ID: ", user.UserID,
		", Name: ", user.Name,
		", EMail: ", user.EMail)
}

// URL for test
const url = "https://jsonplaceholder.typicode.com/users/2"

func getURL() string {
	var urlFlag string
	flag.StringVar(&urlFlag, "url", url, "the URL for remote call to get the User")
	flag.Parse()
	return urlFlag
}

// GetJSONBytes get the Json by a User ID
func GetJSONBytes(id string) ([]byte, error) {
	response, err := http.Get(getURL())
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
func GetUserByJSONBytes(jsonBytes []byte) User {
	var user User
	json.Unmarshal(jsonBytes, &user)
	return user
}

func main() {
	jsonBytes, err := GetJSONBytes("2")
	if err != nil {
		log.Fatal(err)
	}
	user := GetUserByJSONBytes(jsonBytes)

	fmt.Println(user.String())
}
