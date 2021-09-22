package routes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_registerAdminRoutes (t *testing.T) {
	NewTest()
	mapAdmin := registerAdminRoutes()
	assert.Equal(t, registerAdminRoutes(), mapAdmin)
	assert.NotEqual(t, registerAdminRoutes(), "ok boom")
}