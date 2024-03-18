package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHashPassowrd(t *testing.T) {
	assert := assert.New(t)

	t.Run("creates a hash password", func(_ *testing.T) {
		password := "123"

		hashPassword, error := CreateHashPassword(password)

		assert.NotEqual(password, string(hashPassword), "they not should be equal")
		assert.Nil(error, "should be nil")
	})
}

func TestComparePasswords(t *testing.T) {
	assert := assert.New(t)

	t.Run("when they are equivalent", func(_ *testing.T) {
		password := "123"
		hashPassword, _ := CreateHashPassword(password)
    
    error := ComparePasswords(string(hashPassword), password) 

		assert.Nil(error, "should be nil")
	})


	t.Run("when they are not equivalent", func(_ *testing.T) {
		password := "123"
		hashPassword, _ := CreateHashPassword("1212131")
    
    error := ComparePasswords(string(hashPassword), password) 

		assert.NotNil(error, "should not be nil")
	})
}
