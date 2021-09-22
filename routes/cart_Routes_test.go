package routes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_registerCartRoutes (t *testing.T) {
	NewTest()
	mapCart := registerCartRoutes()
	assert.Equal(t, registerCartRoutes(), mapCart)
	assert.NotEqual(t, registerCartRoutes(), mapCart["GET"])
}