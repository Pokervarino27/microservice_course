package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotUserFound(t *testing.T) {

	user, err := GetUser(0)

	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "User 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 1, user.ID)
	assert.EqualValues(t, "Diego", user.FirstName)
	assert.EqualValues(t, "Comihual", user.LastName)
	assert.EqualValues(t, "diego@gmail.com", user.Email)
}
