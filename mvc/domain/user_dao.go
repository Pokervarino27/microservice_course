package domain

import (
	"fmt"
	"microservice_course/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {ID: 1, FirstName: "Diego", LastName: "Comihual", Email: "diego@gmail.com"},
	}
)

//GetUser function that get users from database
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
