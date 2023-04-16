package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoney_NewMoney(t *testing.T) {
	t.Parallel()
	t.Run("New money: Dollar", func(t *testing.T) {
		got := NewMoney(5, MonetaryUnitUSD)
		assert.Equal(t, got, Money{amount: 5, monetaryUnit: MonetaryUnitUSD})
	})
	t.Run("New money: Franc", func(t *testing.T) {
		got := NewMoney(5, MonetaryUnitCHF)
		assert.Equal(t, got, Money{amount: 5, monetaryUnit: MonetaryUnitCHF})
	})
}

func TestMoney_Getter(t *testing.T) {
	t.Parallel()
	t.Run("Amount", func(t *testing.T) {
		m := Money{amount: 10, monetaryUnit: MonetaryUnitUSD}
		assert.Equal(t, 10, m.Amount())
		assert.NotEqual(t, 5, m.Amount())
	})
	t.Run("MonetaryUnit", func(t *testing.T) {
		m := Money{amount: 10, monetaryUnit: MonetaryUnitUSD}
		assert.Equal(t, MonetaryUnitUSD, m.MonetaryUnit())
		assert.NotEqual(t, MonetaryUnitCHF, m.MonetaryUnit())
	})
}

func TestMoney_Reduce(t *testing.T) {
	t.Parallel()
	t.Run("Reduce money", func(t *testing.T) {
		m := Money{4, MonetaryUnitUSD}

		b := NewBank()
		got := b.Reduce(m)

		assert.Equal(t, m, got)
	})
}

func TestMoney_Multiplication(t *testing.T) {
	t.Parallel()
	t.Run("Multiplication", func(t *testing.T) {
		m := Money{amount: 5, monetaryUnit: MonetaryUnitUSD}
		assert.Equal(t, Money{amount: 10, monetaryUnit: MonetaryUnitUSD}, m.Times(2))
		assert.Equal(t, Money{amount: 15, monetaryUnit: MonetaryUnitUSD}, m.Times(3))
	})
}

func TestMoney_Equals(t *testing.T) {
	t.Parallel()
	t.Run("Equals", func(t *testing.T) {
		m := Money{5, MonetaryUnitUSD}
		assert.True(t, m.Equals(m))
		assert.False(t, m.Equals(Money{15, MonetaryUnitUSD}))
		assert.False(t, m.Equals(Money{5, MonetaryUnitCHF}))
	})
}

func TestSum_Reduce(t *testing.T) {
	t.Parallel()
	t.Run("Reduce sum", func(t *testing.T) {
		m := Money{amount: 5, monetaryUnit: MonetaryUnitUSD}
		sum := m.Plus(m, MonetaryUnitUSD)

		b := NewBank()
		got := b.Reduce(sum)

		assert.Equal(t, Money{amount: 10, monetaryUnit: MonetaryUnitUSD}, got)
	})
}
