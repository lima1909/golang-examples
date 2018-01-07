package main

import (
	"encoding/json"
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
const URL = "https://jsonplaceholder.typicode.com/users/"

// GetJSONBytes get the Json by a User ID
func GetJSONBytes(id string) []byte {
	response, err := http.Get(URL + id)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return jsonBytes
}

// GetUserByJSONBytes get the User structur by the JSON bytes
func GetUserByJSONBytes(jsonBytes []byte) User {
	var user User
	json.Unmarshal(jsonBytes, &user)
	return user
}

func main() {
	jsonBytes := GetJSONBytes("2")
	user := GetUserByJSONBytes(jsonBytes)

	fmt.Println(user.String())
}
