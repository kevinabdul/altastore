package routes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_registerProductRoutes (t *testing.T) {
	NewTest()
	maProduct := registerProductRoutes()
	assert.Equal(t, registerProductRoutes(), maProduct)
	assert.NotEqual(t, registerProductRoutes(), 347)
}