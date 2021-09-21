package routes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_registerPaymentRoutes (t *testing.T) {
	NewTest()
	mapPayment := registerPaymentRoutes()
	assert.Equal(t, registerPaymentRoutes(), mapPayment)
	assert.NotEqual(t, registerPaymentRoutes(), map[string][]interface{}{"alias": []interface{}{"Bond", "Bourne", 7}})
}