package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/lima1909/golang-examples/base"
)

var jsonBytes []byte

type MockWriter struct {
}

func (mw MockWriter) Write(jsonB []byte) (int, error) {
	jsonBytes = jsonB
	return 0, nil
}
func (mw MockWriter) WriteHeader(int) {}
func (mw MockWriter) Header() http.Header {
	return make(http.Header)
}

func TestGetUserByID(t *testing.T) {
	u0 := getUserByID(0)
	if 0 != u0.UserID {
		t.Errorf("Expected 0, but was %d", u0.UserID)
	}

	// Test UserID = 1
	u1 := getUserByID(1)
	if 1 != u1.UserID || "Jasmin" != u1.Name {
		t.Errorf("Expected UserID = 1 and Name = Jasmin, but UserID was %d and Name was %s", u1.UserID, u1.Name)
	}
}

func TestGetUserByID_InvalidID(t *testing.T) {
	u100 := getUserByID(100)
	if 100 != u100.UserID || "USER NOT FOUND" != u100.Name {
		t.Errorf("Expected UserID = 100 and Name = USER NOT FOUND, but UserID was %d and Name was %s", u100.UserID, u100.Name)
	}
}

func TestUserHandler(t *testing.T) {
	var mockRequest = http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/user/2"}}
	var mockWriter = MockWriter{}
	userHandler(mockWriter, &mockRequest)

	var user2 *base.User
	json.Unmarshal(jsonBytes, &user2)
	t.Log("USER: ", user2)
	if user2.UserID != 2 {
		t.Errorf("Expected UderID = 2, but it was %d", user2.UserID)
	}
	if user2.Name != "Golang" {
		t.Errorf("Expected Name = GoLang, but it was %s", user2.Name)
	}

}

// func TestMain(m *testing.M) {
// 	os.Exit(m.Run())
// }
