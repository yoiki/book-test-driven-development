package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBank_NewBank(t *testing.T) {
	t.Parallel()
	t.Run("New bank", func(t *testing.T) {
		got := NewBank()
		assert.Equal(t, Bank{}, got)
	})
}

func TestBank_Reduce(t *testing.T) {
	t.Parallel()
	t.Run("Reduce sum", func(t *testing.T) {
		m := Money{amount: 5, monetaryUnit: MonetaryUnitUSD}
		sum := m.Plus(m, MonetaryUnitUSD)

		b := NewBank()
		got := b.Reduce(sum)

		assert.Equal(t, Money{amount: 10, monetaryUnit: MonetaryUnitUSD}, got)
	})
}
