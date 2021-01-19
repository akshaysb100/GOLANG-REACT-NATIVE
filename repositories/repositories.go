package repositories

import (
	"GOLANG-REACT-NATIVE/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/godror/godror"
)

var db *sql.DB

type ProductRepository struct {
	DB *sql.DB
}

func ProvideProductRepostiory(DB *sql.DB) ProductRepository {
	return ProductRepository{DB: DB}
}

func (p *ProductRepository) FindAll() {

	// var users []entity.User
	// p.DB.Find(&users)

	// return users

	rows, err := p.DB.Query("SELECT *FROM user")
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

func (p *ProductRepository) Save(user entity.User) entity.User {

	// res, err := p.DB.Exec(`INSERT INTO user(ID, NAME, MOBILENUMBER) VALUES (1, 'Akshay', '7350003100')`, time.Now())
	// if err != nil {
	// 	panic(err)
	// }

	// numDeleted, err := res.RowsAffected()
	// if err != nil {
	// 	panic(err)
	// }
	// print(numDeleted)

	//p.DB.Save(&user)
	return user
}

func (p *ProductRepository) CreateUser(user entity.User) entity.User {
	ctx := context.Background()
	var err error

	if p.DB == nil {
		err = errors.New("CreateEmployee: db is null")
		return entity.User{}
	}
	var ID string = ""
	var USERNAME string = ""
	var MOBILENUMBER string = ""

	ID = user.ID
	USERNAME = user.USERNAME
	MOBILENUMBER = user.MOBILENUMBER

	fmt.Println(ID, USERNAME, MOBILENUMBER)
	// Check if database is alive.
	err = p.DB.PingContext(ctx)
	if err != nil {
		return entity.User{}
	}

	tsql := "INSERT INTO User (ID,USERNAME , MOBILENUMBER) VALUES (@ID, @USERNAME,@MOBILENUMBER); select convert(bigint, SCOPE_IDENTITY());"

	stmt, err := p.DB.Prepare(tsql)
	if err != nil {
		return entity.User{}
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
		return entity.User{}
	}

	return users
}

func (p *ProductRepository) CreateIssue(issues entity.Isuues) entity.Isuues {
	ctx := context.Background()
	var err error

	if p.DB == nil {
		err = errors.New("CreateEmployee: db is null")
		return entity.Isuues{}
	}
	var id string = ""
	var name string = ""

	id = issues.ID
	name = issues.Issues

	fmt.Println(id, name)
	// Check if database is alive.
	err = p.DB.PingContext(ctx)
	if err != nil {
		return entity.Isuues{}
	}

	tsql := "INSERT INTO user (name, mobileNumber) VALUES (@id, @name,@mobileNumber); select convert(bigint, SCOPE_IDENTITY());"

	stmt, err := p.DB.Prepare(tsql)
	if err != nil {
		return entity.Isuues{}
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("Name", name))
	var issue entity.Isuues
	err = row.Scan(&issues)
	if err != nil {
		return entity.Isuues{}
	}

	return issue
}

func ReadEmployees() (int, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("SELECT Id, Name, Location FROM TestSchema.Employees;")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
	for rows.Next() {
		var name, location string
		var id int

		// Get values from row.
		err := rows.Scan(&id, &name, &location)
		if err != nil {
			return -1, err
		}

		fmt.Printf("ID: %d, Name: %s, Location: %s\n", id, name, location)
		count++
	}

	return count, nil
}
