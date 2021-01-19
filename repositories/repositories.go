package repositories

import (
	"GOLANG-REACT-NATIVE/entity"
	"database/sql"
	"fmt"

	_ "github.com/go-goracle/go-oracledb/internal"
)

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

	p.DB.Save(&user)
	return user
}
