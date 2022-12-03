package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFranc_Multiplication(t *testing.T) {
	t.Run("multiplication", func(t *testing.T) {
		f := Franc{amount: 5}
		assert.Equal(t, &Franc{amount: 10}, f.times(2), fmt.Sprintf("expected $5*2=$10, but got $5*2=%d", f.amount))
		assert.Equal(t, &Franc{amount: 15}, f.times(3), fmt.Sprintf("expected $5*3=$15, but got $5*3=%d", f.amount))
	})
}

func TestFranc_Equals(t *testing.T) {
	t.Run("equals", func(t *testing.T) {
		d := Franc{amount: 5}
		assert.True(t, d.equals(Franc{amount: 5}))
		assert.False(t, d.equals(Franc{amount: 7}))
	})
}
