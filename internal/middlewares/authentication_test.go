package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/eduahcb/hub_api_go/config"
	"github.com/eduahcb/hub_api_go/pkg/security"
	"github.com/stretchr/testify/assert"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func TestAuthentication(t *testing.T) {
	assert := assert.New(t)

	t.Run("when authorization header is not provided", func(_ *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1", nil)
		response := httptest.NewRecorder()

		authHandler := Authentication(handler)
		authHandler(response, request)

		var responseBody ErrorResponse
		json.Unmarshal([]byte(response.Body.String()), &responseBody)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusUnauthorized, expectedStatus, "they should be equal")
		assert.Equal("The authorization header was not provided", responseBody.Message, "they should be equal")
	})

	t.Run("when authorization header does not container a Bearer token", func(_ *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1", nil)
		response := httptest.NewRecorder()

		request.Header.Set("Authorization", "Token dajkdwjlakdjwalkdwa")

		authHandler := Authentication(handler)
		authHandler(response, request)

		var responseBody ErrorResponse
		json.Unmarshal([]byte(response.Body.String()), &responseBody)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusUnauthorized, expectedStatus, "they should be equal")
		assert.Equal("The authorization header does not contain a Bearer token", responseBody.Message, "they should be equal")
	})

	t.Run("when the token is malformed", func(_ *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1", nil)
		response := httptest.NewRecorder()

		request.Header.Set("Authorization", "Bearer dajkdwjlakdjwalkdwa")

		authHandler := Authentication(handler)
		authHandler(response, request)

		var responseBody ErrorResponse
		json.Unmarshal([]byte(response.Body.String()), &responseBody)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusUnauthorized, expectedStatus, "they should be equal")
		assert.Equal("The token is malformed", responseBody.Message, "they should be equal")
	})

	t.Run("when the token is valid", func(_ *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1", nil)
		response := httptest.NewRecorder()

		expirationTime := time.Now().Add(time.Minute * 1).Unix()

		token, _ := security.Token(uint(1), expirationTime, config.Envs.SecretKey)

		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		authHandler := Authentication(handler)
		authHandler(response, request)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusOK, expectedStatus, "they should be equal")
	})
}
