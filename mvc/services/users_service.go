package services

import (
	"microservice_course/mvc/domain"
	"microservice_course/mvc/utils"
)

//GetUser function that get an user
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
