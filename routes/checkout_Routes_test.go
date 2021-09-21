package routes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_registerCheckoutRoutes (t *testing.T) {
	NewTest()
	mapCheckout := registerCheckoutRoutes()
	assert.Equal(t, registerCheckoutRoutes(), mapCheckout)
	assert.NotEqual(t, registerCheckoutRoutes(), map[string][]interface{}{})
}