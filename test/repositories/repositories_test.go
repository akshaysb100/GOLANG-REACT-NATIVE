package repositories

import (
	"GOLANG-REACT-NATIVE/entity"
	"GOLANG-REACT-NATIVE/repositories"
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

var user = entity.User{
	ID:           "1",
	USERNAME:     "akshay",
	MOBILENUMBER: "88123456789",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	repo := repositories.UserRepository{}
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sql := "INSERT INTO Users (ID, USERNAME, MOBILENUMBER) VALUES ( '" + user.ID + "', '" + user.USERNAME + "', '" + user.MOBILENUMBER + "' )"
	_, err = db.Exec(sql)

	if err != nil {
		t.Errorf("error '%s' was not expected, while inserting a row", err)
	}

	mock.ExpectExec("INSERT INTO mytable\\(a, b, c\\)").
		WithArgs("A", "B", "C").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// mock.ExpectQuery("SELECT (.+) FROM testmock").
	// 	WillReturnRows(rows)

	// mock.ExpectQuery("^SELECT \\* FROM user WHERE id=\\? LIMIT 1$").WithArgs("1").WillReturnRows(rows)
	repo.CreateUser(user)
	//mock.ExpectExec(sql)

	// assert.NoError(t, err)
}

func TestShowUsers(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()
	repo := repositories.UserRepository{}
	repo.InitDB()
	rows := sqlmock.NewRows([]string{"ID", "USERNAME", "MOBILENUMBER"}).
		AddRow("1", "akshay", "7350055253").
		AddRow("2", "aniket", "2336325234")

	mock.ExpectQuery("SELECT (.+) FROM testmock").
		WillReturnRows(rows)

	repo.ShowUsers() // <- Add the call to actual function here. Before mock.ExpectationsWereMet
}

func TestShowIssues(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"ID", "Issues"}).
		AddRow("1", "akshay").
		AddRow("2", "aniket")

	mock.ExpectQuery("SELECT (.+) FROM testmock").
		WillReturnRows(rows)
	repo := repositories.UserRepository{}
	repo.ShowIssues() // <- Add the call to actual function here. Before mock.ExpectationsWereMet

	//assert.NotNil(t, val)
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// }
}
