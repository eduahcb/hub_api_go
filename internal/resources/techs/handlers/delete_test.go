package techshandlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	assert := assert.New(t)

	t.Run("when an internal server error occurs when deleting tech", func(_ *testing.T) {
		db, dbMock, mock := databaseMockSuite()
		defer dbMock.Close()

		mock.ExpectBegin()

		err := errors.New("internal error")
		mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "techs"`)).WillReturnError(err)
		mock.ExpectRollback()

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/techs/{id}", Delete(db))

		request, _ := http.NewRequest(http.MethodDelete, "/api/v1/techs/1", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusInternalServerError, expectedStatus, "they should be equal")
	})
}
