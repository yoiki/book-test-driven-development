package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBank_NewBank(t *testing.T) {
	t.Parallel()
	t.Run("New bank", func(t *testing.T) {
		got, err := NewBank()
		assert.Equal(t, Bank{}, got)
		assert.NoError(t, err)
	})
}
