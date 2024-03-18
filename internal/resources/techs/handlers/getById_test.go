package techshandlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetById(t *testing.T) {
	assert := assert.New(t)

	t.Run("when tech was not found", func(_ *testing.T) {
		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()
    
    err := gorm.ErrRecordNotFound
    mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "techs"`)).WillReturnError(err) 
      
		router := mux.NewRouter()
		router.HandleFunc("/api/v1/techs/{id}", GetById(db))

		request, _ := http.NewRequest(http.MethodGet, "/api/v1/techs/1", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusNotFound, expectedStatus, "they should be equal")
	})

	t.Run("when an internal error occurs when searching tech", func(_ *testing.T) {
		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()
    
    err := errors.New("internal server error")
    mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "techs"`)).WillReturnError(err) 
      
		router := mux.NewRouter()
		router.HandleFunc("/api/v1/techs/{id}", GetById(db))

		request, _ := http.NewRequest(http.MethodGet, "/api/v1/techs/1", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusInternalServerError, expectedStatus, "they should be equal")
	})

	t.Run("returns a tech by id", func(_ *testing.T) {
		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()
    
    row := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "svelte")
    mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "techs"`)).WillReturnRows(row)
      
		router := mux.NewRouter()
		router.HandleFunc("/api/v1/techs/{id}", GetById(db))

		request, _ := http.NewRequest(http.MethodGet, "/api/v1/techs/1", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusOK, expectedStatus, "they should be equal")
	})
}
