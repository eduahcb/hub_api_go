package usershandlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/resources/users"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func createSignupBody(email, password, confirmPassowrd string) bytes.Buffer {
	var body bytes.Buffer

	user := users.SignupRequest{
		Name:            "Cleitin",
		Email:           email,
		Password:        password,
		ConfirmPassword: confirmPassowrd,
		ModuleId:        1,
	}

	json.NewEncoder(&body).Encode(user)

	return body
}

func createSignupRequest(body bytes.Buffer, db *database.Database) (*httptest.ResponseRecorder, *http.Request) {
	request, _ := http.NewRequest(http.MethodGet, "/api/v1/signup", &body)
	response := httptest.NewRecorder()

	handler := Signup(db)
	handler(response, request)

	return response, request
}

func TestSignup(t *testing.T) {
	assert := assert.New(t)

	t.Run("when the body is invalid", func(_ *testing.T) {
		var body bytes.Buffer
		body.WriteString("{ invalid_body }")

		db, dbMock, _ := databaseMockSuite()
		defer dbMock.Close()

		response, _ := createSignupRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusBadRequest, expectedStatus, "they should be equal")
	})

	t.Run("when the password and confirm_password don't match", func(_ *testing.T) {
		body := createSignupBody("test@test.com", "12345678", "dawlkdjwadwa")

		db, dbMock, _ := databaseMockSuite()
		defer dbMock.Close()

		response, _ := createSignupRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusBadRequest, expectedStatus, "they should be equal")
	})

	t.Run("when the module was not found ", func(_ *testing.T) {
		body := createSignupBody("test@test.com", "12345678", "12345678")

		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()

		error := gorm.ErrRecordNotFound
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "modules"`)).WillReturnError(error)

		response, _ := createSignupRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusNotFound, expectedStatus, "they should be equal")
	})

	t.Run("when an internal server error occurs when searching for module", func(_ *testing.T) {
		body := createSignupBody("test@test.com", "12345678", "12345678")

		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()

		error := errors.New("internal server error")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "modules"`)).WillReturnError(error)

		response, _ := createSignupRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusInternalServerError, expectedStatus, "they should be equal")
	})

	t.Run("when the email already exists", func(_ *testing.T) {
		body := createSignupBody("test@test.com", "12345678", "12345678")

		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()

		row := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "module 1")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "modules"`)).WillReturnRows(row)

		userRow := sqlmock.NewRows([]string{"id", "email"}).AddRow(1, "test@test.com")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).WillReturnRows(userRow)

		response, _ := createSignupRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusBadRequest, expectedStatus, "they should be equal")
	})

	t.Run("when an internal server error occurs when entering a new user", func(_ *testing.T) {
		body := createSignupBody("test@test.com", "12345678", "12345678")

		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()

		row := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "module 1")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "modules"`)).WillReturnRows(row)

		err := gorm.ErrRecordNotFound
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).WillReturnError(err)

		mock.ExpectBegin()

		err = errors.New("insert error")
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).WillReturnError(err)
		mock.ExpectRollback()

		response, _ := createSignupRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusInternalServerError, expectedStatus, "they should be equal")
	})

	t.Run("creates a new user", func(_ *testing.T) {
		body := createSignupBody("test@test.com", "12345678", "12345678")

		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()

		row := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "module 1")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "modules"`)).WillReturnRows(row)

		err := gorm.ErrRecordNotFound
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).WillReturnError(err)

		mock.ExpectBegin()

		insertRow := sqlmock.NewRows([]string{"id"}).AddRow(1)
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).WillReturnRows(insertRow)
		mock.ExpectCommit()

		response, _ := createSignupRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusNoContent, expectedStatus, "they should be equal")
	})
}
