package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoney_NewDollar(t *testing.T) {
	t.Parallel()
	t.Run("NewDollar", func(t *testing.T) {
		assert.Equal(t, &Money{amount: 5, monetaryUnit: "USD"}, NewDollar(5))
	})
	t.Run("NewFranc", func(t *testing.T) {
		assert.Equal(t, &Money{amount: 5, monetaryUnit: "CHF"}, NewFranc(5))
	})
}

func TestMoney_Getter(t *testing.T) {
	t.Parallel()
	t.Run("Amount", func(t *testing.T) {
		assert.Equal(t, 10, NewDollar(10).Amount())
		assert.NotEqual(t, 10, NewDollar(5).Amount())
	})
	t.Run("MonetaryUnit", func(t *testing.T) {
		assert.Equal(t, "USD", NewDollar(10).MonetaryUnit())
		assert.NotEqual(t, "CHF", NewDollar(10).MonetaryUnit())
	})
}

func TestMoney_Multiplication(t *testing.T) {
	t.Parallel()
	t.Run("Multiplication", func(t *testing.T) {
		d := NewDollar(5)
		assert.Equal(t, NewDollar(10), d.Times(2))
		assert.Equal(t, NewDollar(15), d.Times(3))
	})
}

func TestMoney_Equals(t *testing.T) {
	t.Parallel()
	t.Run("Equals", func(t *testing.T) {
		d := NewDollar(5)
		assert.True(t, d.Equals(NewDollar(5)))
		assert.False(t, d.Equals(NewDollar(10)))
		assert.False(t, d.Equals(NewFranc(5)))
	})
}
