package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonetaryUnit_NewMoneyUnit(t *testing.T) {
	t.Parallel()
	t.Run("Define monetary unit: USD", func(t *testing.T) {
		got, err := NewMoneyUnit("USD")
		assert.Equal(t, MonetaryUnitUSD, got)
		assert.NoError(t, err)
	})
	t.Run("Invalid monetary unit", func(t *testing.T) {
		got, err := NewMoneyUnit("Unknown")
		assert.Equal(t, MonetaryUnitUnknown, got)
		assert.Error(t, err)
	})
}

func TestMonetaryUnit_String(t *testing.T) {
	t.Parallel()
	t.Run("Get string from correct monetary unit: USD", func(t *testing.T) {
		got, err := MonetaryUnitUSD.String()
		assert.Equal(t, "USD", got)
		assert.NoError(t, err)
	})
	t.Run("Raise error if monetary unit is invalid", func(t *testing.T) {
		got, err := MonetaryUnitUnknown.String()
		assert.Equal(t, "Unknown", got)
		assert.Error(t, err)
	})
}
