package repositories

import (
	"GOLANG-REACT-NATIVE/entity"
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var u = &entity.User{
	ID:           "1",
	USERNAME:     "Momo",
	MOBILENUMBER: "08123456789",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCreateUser(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT INTO Users (ID, USERNAME , MOBILENUMBER) VALUES (@ID, @USERNAME, @MOBILENUMBER); select convert(bigint, SCOPE_IDENTITY());"

	prep := mock.Prepare(query)
	prep.ExpectExec().WithArgs(u.ID, u.USERNAME, u.MOBILENUMBER).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Create(u)
	assert.NoError(t, err)
}
