package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDollar_Multiplication(t *testing.T) {
	t.Run("multiplication", func(t *testing.T) {
		d := Dollar{amount: 5}
		assert.Equal(t, &Dollar{amount: 10}, d.times(2), fmt.Sprintf("expected $5*2=$10, but got $5*2=%d", d.amount))
		assert.Equal(t, &Dollar{amount: 15}, d.times(3), fmt.Sprintf("expected $5*3=$15, but got $5*3=%d", d.amount))
	})
}

func TestDollar_Equals(t *testing.T) {
	t.Run("equals", func(t *testing.T) {
		d := Dollar{amount: 5}
		assert.True(t, d.equals(Dollar{amount: 5}))
		assert.False(t, d.equals(Dollar{amount: 7}))
	})
}
