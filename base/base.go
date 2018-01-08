package base

import "fmt"

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
