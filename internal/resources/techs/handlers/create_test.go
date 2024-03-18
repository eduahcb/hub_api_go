package techshandlers

import (
	"bytes"
	"context"
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

func createCreateRequest(body bytes.Buffer, db *database.Database) (*httptest.ResponseRecorder, *http.Request) {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/techs/{id}", Create(db))

	request, _ := http.NewRequest(http.MethodPost, "/api/v1/techs/1", &body)
	response := httptest.NewRecorder()

	ctx := request.Context()
	ctx = context.WithValue(ctx, "userId", uint(1))
	
  router.ServeHTTP(response, request.WithContext(ctx))

	return response, request
}

func createCreateBody(name string, leveId uint) bytes.Buffer {
	var body bytes.Buffer

	tech := techs.TechRequest{
		Name:    name,
		LevelId: leveId,
	}

	json.NewEncoder(&body).Encode(tech)

	return body
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	t.Run("when the body is invalid", func(_ *testing.T) {
		var body bytes.Buffer
		body.WriteString("{ invalid_body }")

		db, dbMock, _ := databaseMockSuite()
		defer dbMock.Close()

		response, _ := createCreateRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusBadRequest, expectedStatus, "they should be equal")
	})

	t.Run("when the name is empty", func(_ *testing.T) {
		body := createCreateBody("", 1)

		db, dbMock, _ := databaseMockSuite()
		defer dbMock.Close()

		response, _ := createCreateRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusBadRequest, expectedStatus, "they should be equal")
	})

	t.Run("when level was not found", func(_ *testing.T) {
		body := createCreateBody("svelte", 1)

		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()

		err := gorm.ErrRecordNotFound
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "levels"`)).WillReturnError(err)

		response, _ := createCreateRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusNotFound, expectedStatus, "they should be equal")
	})

	t.Run("when an internal error occurs when searching level", func(_ *testing.T) {
		body := createCreateBody("svelte", 1)

		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()

		err := errors.New("internal server error")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "techs"`)).WillReturnError(err)

		response, _ := createCreateRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusInternalServerError, expectedStatus, "they should be equal")
	})

	t.Run("when an internal server error occurs when creating tech", func(_ *testing.T) {
		body := createCreateBody("svelte", 1)

		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()

		levelRow := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "iniciante")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "levels"`)).WillReturnRows(levelRow)

		mock.ExpectBegin()
		err := errors.New("insert error")

		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "techs"`)).WillReturnError(err)
		mock.ExpectRollback()

		response, _ := createCreateRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusInternalServerError, expectedStatus, "they should be equal")
	})

	t.Run("creates a new tech", func(_ *testing.T) {
		body := createCreateBody("svelte", 1)

		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()

		row := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "iniciante")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "levels"`)).WillReturnRows(row)

		mock.ExpectBegin()

		techRow := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "svelte")
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "techs"`)).WillReturnRows(techRow)
		mock.ExpectCommit()

		response, _ := createCreateRequest(body, db)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusCreated, expectedStatus, "they should be equal")
	})
}
