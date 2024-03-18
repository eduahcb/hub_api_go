package techshandlers

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eduahcb/hub_api_go/internal/database"
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

func TestGetAll(t *testing.T) {
	assert := assert.New(t)

	t.Run("when an internal error occurs when searching techs", func(_ *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/techs", nil)
		response := httptest.NewRecorder()

		ctx := request.Context()
		ctx = context.WithValue(ctx, "userId", uint(1))

    db, dbMock, mock := databaseMockSuite()
    defer dbMock.Close()
    
    err := errors.New("internal error")
    mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "techs"`)).WillReturnError(err)

		handler := GetAll(db)
		handler(response, request.WithContext(ctx))

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusInternalServerError, expectedStatus, "they should be equal")
	})

	t.Run("returns all techs by user", func(_ *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/techs", nil)
		response := httptest.NewRecorder()

		ctx := request.Context()
		ctx = context.WithValue(ctx, "userId", uint(1))

    db, dbMock, mock := databaseMockSuite()
    defer dbMock.Close()

    row := sqlmock.NewRows([]string{"id", "name", "user_id"}).AddRow(1, "react", 1)
    mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "techs"`)).WillReturnRows(row)

		handler := GetAll(db)
		handler(response, request.WithContext(ctx))

		expectedStatus := response.Result().StatusCode
		assert.Equal(http.StatusOK, expectedStatus, "they should be equal")
	})
}
