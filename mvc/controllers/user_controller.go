package controllers

import (
	"microservice_course/mvc/services"
	"microservice_course/mvc/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetUser function that get an User
func GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}

	user, apiErr := services.GetUser(userID)
	if apiErr != nil {
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}
	c.JSON(http.StatusOK, user)

}
