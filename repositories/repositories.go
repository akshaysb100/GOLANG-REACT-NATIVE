package repositories

import (
	"GOLANG-REACT-NATIVE/entity"
	"context"
	"database/sql"
	"fmt"
	"time"

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

	sql := "INSERT INTO Users (ID, USERNAME, MOBILENUMBER) VALUES ( '" + user.ID + "', '" + user.USERNAME + "', '" + user.MOBILENUMBER + "' )"
	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println("Error updating the database", "err", err)
		return entity.User{
			ID:           "",
			USERNAME:     "",
			MOBILENUMBER: "",
		}
	}

	return user
}

func (p *UserRepository) CreateIssue(issues entity.Isuues) entity.Isuues {
	db := initDB()

	sql := "INSERT INTO issues (ID, ISSUE) VALUES ( '" + issues.ID + "', '" + issues.ISSUE + "')"
	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println("Error updating the database", "err", err)
		return entity.Isuues{
			ID:    "",
			ISSUE: "",
		}
	}

	return issues
}

func (p *UserRepository) ShowUsers() []*entity.User {
	db := initDB()

	users := make([]*entity.User, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT ID, USERNAME, MOBILENUMBER FROM Users")
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		user := new(entity.User)
		err = rows.Scan(
			&user.ID,
			&user.USERNAME,
			&user.MOBILENUMBER,
		)

		if err != nil {
			return nil
		}
		users = append(users, user)
	}

	return users
}

func (p *UserRepository) ShowIssues() []*entity.Isuues {
	db := initDB()

	users := make([]*entity.Isuues, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT ID, ISSUE FROM issues")
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		user := new(entity.Isuues)
		err = rows.Scan(
			&user.ID,
			&user.ISSUE,
		)

		if err != nil {
			return nil
		}
		users = append(users, user)
	}

	return users
}
