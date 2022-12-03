package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoney_Multiplication(t *testing.T) {
	t.Run("Multiplication: Dollar", func(t *testing.T) {
		d := NewDollar(5)
		assert.Equal(
			t,
			NewDollar(10),
			d.Times(2),
			fmt.Sprintf("expected $5*2=$10, but got $5*2=%d", d.amount),
		)
		assert.Equal(
			t,
			NewDollar(15),
			d.Times(3),
			fmt.Sprintf("expected $5*3=$15, but got $5*3=%d", d.amount),
		)
	})
	t.Run("Multiplication: Franc", func(t *testing.T) {
		f := NewFranc(5)
		assert.Equal(
			t,
			NewFranc(10),
			f.Times(2),
			fmt.Sprintf("expected $5*2=$10, but got $5*2=%d", f.amount),
		)
		assert.Equal(
			t,
			NewFranc(15),
			f.Times(3),
			fmt.Sprintf("expected $5*3=$15, but got $5*3=%d", f.amount),
		)
	})
}

func TestMoney_Equals(t *testing.T) {
	t.Run("Equals", func(t *testing.T) {
		d := NewDollar(5)
		assert.True(t, d.Equals(NewDollar(5)))
		assert.False(t, d.Equals(NewDollar(10)))
		assert.False(t, d.Equals(NewFranc(5)))
	})
}
