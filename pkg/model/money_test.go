package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoney_Multiplication(t *testing.T) {
	t.Run("Multiplication: Dollar", func(t *testing.T) {
		m := Money{amount: 5, monetaryUnit: MonetaryUnit(0)}
		assert.Equal(
			t,
			&Money{amount: 10, monetaryUnit: MonetaryUnit(0)},
			m.Times(2),
			fmt.Sprintf("expected $5*2=$10, but got $5*2=%d", m.amount),
		)
		assert.Equal(
			t,
			&Money{amount: 15, monetaryUnit: MonetaryUnit(0)},
			m.Times(3),
			fmt.Sprintf("expected $5*3=$15, but got $5*3=%d", m.amount),
		)
	})
	t.Run("Multiplication: Franc", func(t *testing.T) {
		m := Money{amount: 5, monetaryUnit: MonetaryUnit(1)}
		assert.Equal(
			t,
			&Money{amount: 10, monetaryUnit: MonetaryUnit(1)},
			m.Times(2),
			fmt.Sprintf("expected $5*2=$10, but got $5*2=%d", m.amount),
		)
		assert.Equal(
			t,
			&Money{amount: 15, monetaryUnit: MonetaryUnit(1)},
			m.Times(3),
			fmt.Sprintf("expected $5*3=$15, but got $5*3=%d", m.amount),
		)
	})
}

func TestMoney_Equals(t *testing.T) {
	t.Run("Equals", func(t *testing.T) {
		m := Money{amount: 5, monetaryUnit: MonetaryUnit(0)}
		assert.True(t, m.Equals(&Money{amount: 5, monetaryUnit: MonetaryUnit(0)}))
		assert.False(t, m.Equals(&Money{amount: 7, monetaryUnit: MonetaryUnit(0)}))
		assert.False(t, m.Equals(&Money{amount: 5, monetaryUnit: MonetaryUnit(1)}))

		m = Money{amount: 5, monetaryUnit: MonetaryUnit(1)}
		assert.True(t, m.Equals(&Money{amount: 5, monetaryUnit: MonetaryUnit(1)}))
		assert.False(t, m.Equals(&Money{amount: 7, monetaryUnit: MonetaryUnit(1)}))
		assert.False(t, m.Equals(&Money{amount: 5, monetaryUnit: MonetaryUnit(0)}))
	})
}
