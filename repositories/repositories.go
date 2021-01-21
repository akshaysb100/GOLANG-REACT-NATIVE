package repositories

import (
	"GOLANG-REACT-NATIVE/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/godror/godror"
)

func initDB() *sql.DB {
	db, err := sql.Open("godror", `user="akshay" password="password123" connectString="localhost:1521/ORCL"
                               libDir="/Users/akshaybavalekar/Downloads/instantclient_19_8"`)
	if err != nil {
		panic(err)
	}
	return db
}

type UserRepository struct {
	DB *sql.DB
}

func ProvideUserRepostiory(DB *sql.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (p *UserRepository) FindAll() {

	rows, err := p.DB.Query("SELECT *FROM Users")
	if err != nil {
		fmt.Println(".....Error processing query")
		fmt.Println(err)
		return
	}

	var NAME, MOBILENUMBER string

	for rows.Next() {
		rows.Scan(&NAME, &MOBILENUMBER)
		fmt.Println(NAME, MOBILENUMBER)
	}
}

func (p *UserRepository) CreateUser(user entity.User) entity.User {
	db := initDB()

	ctx := context.Background()
	var err error

	var ID string = ""
	var USERNAME string = ""
	var MOBILENUMBER string = ""

	ID = user.ID
	USERNAME = user.USERNAME
	MOBILENUMBER = user.MOBILENUMBER

	err = db.PingContext(ctx)
	if err != nil {
		return entity.User{
			ID:           ID,
			USERNAME:     USERNAME,
			MOBILENUMBER: MOBILENUMBER,
		}
	}

	tsql := "INSERT INTO User (ID,USERNAME , MOBILENUMBER) VALUES (@ID, @USERNAME, @MOBILENUMBER);"
	//tsql := "INSERT INTO Users  (ID, USERNAME, MOBILENUMBER)  VALUES  ('1', 'Akshay','7350055254');"

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return entity.User{
			ID:           ID,
			USERNAME:     USERNAME,
			MOBILENUMBER: MOBILENUMBER,
		}
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("Name", ID),
		sql.Named("USERNAME", USERNAME),
		sql.Named("MOBILENUMBER", MOBILENUMBER))

	var users entity.User
	err = row.Scan(&users)
	if err != nil {
		return entity.User{
			ID:           ID,
			USERNAME:     USERNAME,
			MOBILENUMBER: MOBILENUMBER,
		}
	}

	return users
}

func (p *UserRepository) CreateIssue(issues entity.Isuues) entity.Isuues {
	ctx := context.Background()
	var err error

	if p.DB == nil {
		err = errors.New("CreateEmployee: db is null")
		return entity.Isuues{}
	}
	var ID string = ""
	var ISSUE string = ""

	ID = issues.ID
	ISSUE = issues.ISSUE

	//Check if database is alive.
	err = p.DB.PingContext(ctx)
	if err != nil {
		return entity.Isuues{}
	}

	tsql := "INSERT INTO issues (ID, ISSUE) VALUES (@ID, @ISSUE);"

	stmt, err := p.DB.Prepare(tsql)
	if err != nil {
		return entity.Isuues{}
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("Name", ID),
		sql.Named("Name", ISSUE))
	var issue entity.Isuues
	err = row.Scan(&issues)
	if err != nil {
		return entity.Isuues{}
	}

	return issue
}

func (p *UserRepository) ShowUsers() int {
	db := initDB()
	ctx := context.Background()

	// Check if database is alive.
	if db == nil {
		return -1
	}

	err := db.PingContext(ctx)
	if err != nil {
		return -1
	}

	tsql := fmt.Sprintf("SELECT *FROM Users")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
	for rows.Next() {
		var ID, USERNAME, MOBILENUMBER string

		// Get values from row.
		err := rows.Scan(&ID, &USERNAME, &MOBILENUMBER)
		if err != nil {
			return -1
		}

		//fmt.Printf("ID: %d, Name: %s, Location: %s\n", ID, USERNAME, MOBILENUMBER)
		count++
	}

	return count
}

func (p *UserRepository) ShowIssues() int {
	db := initDB()
	ctx := context.Background()

	// Check if database is alive.
	if db == nil {
		return -1
	}

	err := db.PingContext(ctx)
	if err != nil {
		return -1
	}

	tsql := fmt.Sprintf("SELECT *FROM issues")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
	for rows.Next() {
		var ID, ISSUE string

		// Get values from row.
		err := rows.Scan(&ID, &ISSUE)
		if err != nil {
			return -1
		}

		//fmt.Printf("ID: %d, Name: %s, Location: %s\n", ID, USERNAME, MOBILENUMBER)
		count++
	}

	return count
}
