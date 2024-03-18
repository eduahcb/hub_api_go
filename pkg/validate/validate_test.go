package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type bodyExample struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func TestValidateRequestBody(t *testing.T) {
	assert := assert.New(t)

	t.Run("when field must be a valid email", func(_ *testing.T) {
		validBody := bodyExample{
			Email:    "test",
			Password: "12345678",
		}

		errors, _ := ValidateRequestBody(validBody)

		expectedMessage := errors[0].Message

		assert.NotNil(errors, "should not be nil")
		assert.Equal("The field 'email' must be a valid email", expectedMessage)
	})

	t.Run("when field is required", func(_ *testing.T) {
		validBody := bodyExample{
			Email: "test@test.com",
		}

		errors, _ := ValidateRequestBody(validBody)

		expectedMessage := errors[0].Message

		assert.NotNil(errors, "should not be nil")
		assert.Equal("The field 'password' is required", expectedMessage)
	})
}
