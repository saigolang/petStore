package backend

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateId(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		result := validateId("1")
		assert.Nil(t, result)
	})

	t.Run("bad input", func(t *testing.T) {
		result := validateId("")
		assert.NotNil(t, result)
	})
}
