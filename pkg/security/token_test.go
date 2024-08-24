package security

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	assert := assert.New(t)

	t.Run("creates a token", func(_ *testing.T) {
		expirationTime := time.Now().Add(time.Second * 1).Unix()
		token, error := Token(uint(2), expirationTime, "my_secret")

		assert.NotEmpty(token, "should not be empty")
		assert.Nil(error, "should be nil")
	})
}

func TestValidateToken(t *testing.T) {
	assert := assert.New(t)

	t.Run("when the token is malformed", func(_ *testing.T) {
		token := "kkk"

		_, _, err := ValidateToken(token, "my_secret")

		assert.Equal("The token is malformed", err.Error(), "they should be equal")
	})

	t.Run("when the token is invalid", func(_ *testing.T) {
		expirationTime := time.Now().Add(time.Second * 1).Unix()
		token, _ := Token(uint(2), expirationTime, "my_secret")

		_, _, err := ValidateToken(token+"a", "my_secret")

		assert.Equal("The token signature is invalid", err.Error(), "they should be equal")
	})

	t.Run("when the token expired", func(_ *testing.T) {
		expirationTime := time.Now().Add(time.Millisecond * 20).Unix()
		token, _ := Token(uint(2), expirationTime, "my_secret")

		time.Sleep(time.Millisecond * 30)
		_, _, err := ValidateToken(token, "my_secret")

		assert.Equal("The token has expired", err.Error(), "they should be equal")
	})

	t.Run("when token is valid", func(_ *testing.T) {
		expirationTime := time.Now().Add(time.Minute * 10).Unix()
		token, _ := Token(2, expirationTime, "my_secret")

		userId, _, _ := ValidateToken(token, "my_secret")

		assert.Equal(uint(2), userId, "they should be equal")
	})
}

func TestGenerateTokenHash(t *testing.T) {
	  assert := assert.New(t)
    
    t.Run("createst token hash", func(_ *testing.T) {
      expirationTime := time.Now().Add(time.Second * 1).Unix()
      token, _ := Token(uint(2), expirationTime, "my_secret")


      tokenHash := GenerateTokenHash(token)

      assert.NotNil(tokenHash, "should be not nil")
    })

}
