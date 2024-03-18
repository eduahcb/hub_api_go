package usershandlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/pkg/security"

	"github.com/eduahcb/hub_api_go/internal/resources/users"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func databaseMockSuite() (*database.Database, *sql.DB, sqlmock.Sqlmock) {
	dbMock, mock, _ := sqlmock.New()

	dialector := postgres.New(postgres.Config{
		Conn: dbMock,
	})

	client, _ := gorm.Open(dialector, &gorm.Config{})

	db := &database.Database{
		Client: client,
	}

	return db, dbMock, mock
}

func createSigninBody(email, password string) bytes.Buffer {
	var b bytes.Buffer

	body := users.SigninRequest{
		Email:    email,
		Password: password,
	}

	json.NewEncoder(&b).Encode(body)
	return b
}

func createSigninRequest(body bytes.Buffer, db *database.Database) (*httptest.ResponseRecorder, *http.Request) {
	request, _ := http.NewRequest(http.MethodGet, "/api/v1/signin", &body)
	response := httptest.NewRecorder()

	handler := Signin(db)
	handler(response, request)

	return response, request
}

func TestSignin(t *testing.T) {
	assert := assert.New(t)

	t.Run("when the body request is invalid", func(_ *testing.T) {
		var body bytes.Buffer
		body.WriteString(`{ invalid_json }`)

		db, dbMock, _ := databaseMockSuite()
		defer dbMock.Close()

		response, _ := createSigninRequest(body, db)

		expecteStatus := response.Result().StatusCode

		assert.Equal(http.StatusBadRequest, expecteStatus, "they should be equal")
	})

	t.Run("when the email field is empty", func(_ *testing.T) {
		db, dMock, _ := databaseMockSuite()
		defer dMock.Close()

		body := createSigninBody("", "12345678")

		response, _ := createSigninRequest(body, db)

		expecteStatus := response.Result().StatusCode
		assert.Equal(http.StatusBadRequest, expecteStatus, "they should be equal")
	})

	t.Run("when the user not found", func(_ *testing.T) {
		db, dMock, mock := databaseMockSuite()
		defer dMock.Close()

		error := gorm.ErrRecordNotFound
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).WillReturnError(error)

		body := createSigninBody("test@test.com", "12345678")

		response, _ := createSigninRequest(body, db)

		expecteStatus := response.Result().StatusCode
		assert.Equal(http.StatusNotFound, expecteStatus, "they should be equal")
	})

	t.Run("when an internal server error occurs", func(_ *testing.T) {
		db, dMock, mock := databaseMockSuite()
		defer dMock.Close()

		error := errors.New("internal server error")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).WillReturnError(error)

		body := createSigninBody("test@test.com", "12345678")

		response, _ := createSigninRequest(body, db)

		expecteStatus := response.Result().StatusCode
		assert.Equal(http.StatusInternalServerError, expecteStatus, "they should be equal")
	})

	t.Run("when the credentials don't match", func(_ *testing.T) {
		db, dMock, mock := databaseMockSuite()
		defer dMock.Close()

		hashPassword, _ := security.CreateHashPassword("123456789")

		row := sqlmock.NewRows([]string{"id", "password"}).AddRow(1, string(hashPassword))
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).WillReturnRows(row)

		body := createSigninBody("test@test.com", "12345678")
		response, _ := createSigninRequest(body, db)

		expecteStatus := response.Result().StatusCode
		assert.Equal(http.StatusUnauthorized, expecteStatus, "they should be equal")
	})

	t.Run("when the credentials match", func(_ *testing.T) {
		db, dMock, mock := databaseMockSuite()
		defer dMock.Close()

		hashPassword, _ := security.CreateHashPassword("12345678")

		row := sqlmock.NewRows([]string{"id", "password"}).AddRow(1, string(hashPassword))
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).WillReturnRows(row)

		body := createSigninBody("test@test.com", "12345678")
		response, _ := createSigninRequest(body, db)

		expecteStatus := response.Result().StatusCode
		assert.Equal(http.StatusNoContent, expecteStatus, "they should be equal")
	})
}
