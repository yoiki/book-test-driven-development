package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoney_NewMoney(t *testing.T) {
	t.Parallel()
	t.Run("New money: Dollar", func(t *testing.T) {
		got, err := NewMoney(5, MonetaryUnitUSD)
		assert.Equal(t, got, Money{5, MonetaryUnitUSD})
		assert.NoError(t, err)
	})
	t.Run("New money: Franc", func(t *testing.T) {
		got, err := NewMoney(5, MonetaryUnitCHF)
		assert.Equal(t, got, Money{5, MonetaryUnitCHF})
		assert.NoError(t, err)
	})
	t.Run("Invalid money", func(t *testing.T) {
		got, err := NewMoney(5, MonetaryUnitUnknown)
		assert.Equal(t, got, Money{})
		assert.Error(t, err)
	})
}

func TestMoney_Getter(t *testing.T) {
	t.Parallel()
	t.Run("Amount", func(t *testing.T) {
		m := Money{10, MonetaryUnitUSD}
		assert.Equal(t, 10, m.Amount())
		assert.NotEqual(t, 5, m.Amount())
	})
	t.Run("MonetaryUnit", func(t *testing.T) {
		m := Money{10, MonetaryUnitUSD}
		assert.Equal(t, MonetaryUnitUSD, m.MonetaryUnit())
		assert.NotEqual(t, MonetaryUnitCHF, m.MonetaryUnit())
	})
}

func TestMoney_Multiplication(t *testing.T) {
	t.Parallel()
	t.Run("Multiplication", func(t *testing.T) {
		m := Money{5, MonetaryUnitUSD}
		assert.Equal(t, Money{10, MonetaryUnitUSD}, m.Times(2))
		assert.Equal(t, Money{15, MonetaryUnitUSD}, m.Times(3))
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

func TestMoney_Plus(t *testing.T) {
	t.Parallel()
	t.Run("Simple Addition", func(t *testing.T) {
		b, err := NewBank()
		if err != nil {
			t.Fatal(err)
		}
		m := Money{5, MonetaryUnitUSD}
		sum := m.Plus(m)
		got := b.Reduce(sum)
		assert.Equal(t, Money{10, MonetaryUnitUSD}, got)
		//
		//sum = m.Plus(Money{15, MonetaryUnitUSD})
		//got = b.Reduce(sum)
		//assert.Equal(t, Money{20, MonetaryUnitUSD}, got)
	})
}
