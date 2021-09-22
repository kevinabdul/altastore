package routes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_registerUserRoutes (t *testing.T) {
	NewTest()
	mapUser := registerUserRoutes()
	assert.Equal(t, registerUserRoutes(), mapUser)
	assert.NotEqual(t, registerUserRoutes(), 12)
}