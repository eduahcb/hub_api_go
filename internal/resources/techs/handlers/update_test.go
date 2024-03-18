package techshandlers

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
	"github.com/eduahcb/hub_api_go/internal/resources/techs"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func createUpdateRequest(body bytes.Buffer, db *database.Database) (*httptest.ResponseRecorder, *http.Request) {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/techs/{id}", Update(db))

	request, _ := http.NewRequest(http.MethodPut, "/api/v1/techs/1", &body)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response, request
}

func createUpdateBody(name string, leveId uint) bytes.Buffer {
	var body bytes.Buffer

	tech := techs.TechRequest{
		Name:    name,
		LevelId: leveId,
	}

	json.NewEncoder(&body).Encode(tech)

	return body
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)

	t.Run("when the body is invalid", func(_ *testing.T) {
		var body bytes.Buffer
		body.WriteString("{ invalid_body }")

		db, dbMock, _ := databaseMockSuite()
		defer dbMock.Close()

		response, _ := createUpdateRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusBadRequest, expectedStatus, "they should be equal")
	})

	t.Run("when the name is empty", func(_ *testing.T) {
		body := createUpdateBody("", 1)

		db, dbMock, _ := databaseMockSuite()
		defer dbMock.Close()

		response, _ := createUpdateRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusBadRequest, expectedStatus, "they should be equal")
	})

	t.Run("when level was not found", func(_ *testing.T) {
		body := createUpdateBody("svelte", 1)

		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()

		err := gorm.ErrRecordNotFound
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "levels"`)).WillReturnError(err)

		response, _ := createUpdateRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusNotFound, expectedStatus, "they should be equal")
	})

	t.Run("when an internal error occurs when searching level", func(_ *testing.T) {
		body := createUpdateBody("svelte", 1)

		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()

		err := errors.New("internal server error")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "techs"`)).WillReturnError(err)

		response, _ := createUpdateRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusInternalServerError, expectedStatus, "they should be equal")
	})

	t.Run("when an internal server error occurs when searching tech", func(_ *testing.T) {
		body := createUpdateBody("svelte", 1)

		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()

		levelRow := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "iniciante")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "levels"`)).WillReturnRows(levelRow)

		err := errors.New("internal server error")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "techs"`)).WillReturnError(err)

		response, _ := createUpdateRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusInternalServerError, expectedStatus, "they should be equal")
	})

	t.Run("when the tech was not found", func(_ *testing.T) {
		body := createUpdateBody("svelte", 1)

		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()

		levelRow := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "iniciante")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "levels"`)).WillReturnRows(levelRow)

		err := gorm.ErrRecordNotFound
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "techs"`)).WillReturnError(err)

		response, _ := createUpdateRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusNotFound, expectedStatus, "they should be equal")
	})

	t.Run("when an internal server error occurs when updating tech", func(_ *testing.T) {
		body := createUpdateBody("svelte", 1)

		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()

		levelRow := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "iniciante")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "levels"`)).WillReturnRows(levelRow)

		techRow := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "tech")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "techs"`)).WillReturnRows(techRow)

		mock.ExpectBegin()

		err := errors.New("internal error")
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "techs" SET`)).WillReturnError(err)
		mock.ExpectRollback()

		response, _ := createUpdateRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusInternalServerError, expectedStatus, "they should be equal")
	})

	t.Run("returns a tech", func(_ *testing.T) {
		body := createUpdateBody("svelte", 1)

		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()

		row := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "iniciante")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "levels"`)).WillReturnRows(row)

		techRow := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "tech")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "techs"`)).WillReturnRows(techRow)

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "techs" SET`)).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		response, _ := createUpdateRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusOK, expectedStatus, "they should be equal")
	})
}
