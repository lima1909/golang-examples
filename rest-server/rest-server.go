package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/lima1909/golang-examples/base"
)

// two kind of init a struct-slice
var users = []base.User{
	{UserID: 0, Name: "Mario", EMail: "LiMa@foo.org"},
	{UserID: 1, Name: "Jasmin", EMail: "JaRo@foo.org"},
}

// second part of init the struct-slice
func init() {
	users = append(users, base.User{UserID: 2, Name: "Golang", EMail: "Go@foo.org"})
	users = append(users, base.User{UserID: 3, Name: "Docker", EMail: "Docker@foo.org"})
	users = append(users, base.User{UserID: 4, Name: "Linus", EMail: "Linux@foo.org"})
}

// find the User by the arg id from th slice-struct users
func getUserByID(id int) base.User {
	for _, u := range users {
		if id == u.UserID {
			return u
		}
	}
	return base.User{UserID: id, Name: "USER NOT FOUND"}
}

// userHandler create a User struc
func userHandler(w http.ResponseWriter, req *http.Request) {
	// not the best idea, better with: http://www.gorillatoolkit.org/pkg/mux ?
	id := strings.TrimPrefix(req.URL.Path, "/user/")
	log.Println("ID: ", id)

	log.Println("Request: ", req.Method, req.URL.Path)
	strID, _ := strconv.Atoi(id)
	user := getUserByID(strID)
	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}
	log.Println("JSON: ", string(userJSON))
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)

}

func main() {
	log.Println("Starting Http-Server ...")
	http.HandleFunc("/user/", userHandler)
	http.ListenAndServe(":8082", nil)
}
