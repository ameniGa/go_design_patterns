package factory_test

import (
	"github.com/ameniGa/go_design_patterns/creational/factory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPaymentMethod(t *testing.T) {
	t.Run("valid type : cash", func(t *testing.T) {
		payment, err := factory.GetPaymentMethod(1)
		assert.NoError(t, err)
		msg := payment.Pay(100)
		assert.Contains(t, msg, "cash")
	})
	t.Run("valid type : debit card", func(t *testing.T) {
		payment, err := factory.GetPaymentMethod(2)
		assert.NoError(t, err)
		msg := payment.Pay(100)
		assert.Contains(t, msg, "debit card")
	})
	t.Run("invalid type", func(t *testing.T) {
		_, err := factory.GetPaymentMethod(3)
		assert.Error(t, err)
	})
}
