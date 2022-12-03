package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoney_Getter(t *testing.T) {
	t.Parallel()
	t.Run("Amount", func(t *testing.T) {
		assert.Equal(t, 10, NewDollar(10).Amount())
		assert.Equal(t, 10, NewFranc(10).Amount())
		assert.NotEqual(t, 10, NewFranc(5).Amount())
	})
	t.Run("MonetaryUnit", func(t *testing.T) {
		assert.Equal(t, "USD", NewDollar(10).MonetaryUnit())
		assert.Equal(t, "CHF", NewFranc(10).MonetaryUnit())
		assert.NotEqual(t, "USD", NewFranc(10).MonetaryUnit())
	})
}

func TestMoney_Multiplication(t *testing.T) {
	t.Parallel()
	t.Run("Multiplication: Dollar", func(t *testing.T) {
		d := NewDollar(5)
		assert.Equal(t, NewDollar(10), d.Times(2))
		assert.Equal(t, NewDollar(15), d.Times(3))
	})
	t.Run("Multiplication: Franc", func(t *testing.T) {
		f := NewFranc(5)
		assert.Equal(t, NewFranc(10), f.Times(2))
		assert.Equal(t, NewFranc(15), f.Times(3))
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
