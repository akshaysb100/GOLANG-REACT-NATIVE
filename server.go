package main

import (
	"GOLANG-REACT-NATIVE/controller"
	"GOLANG-REACT-NATIVE/repositories"
	"GOLANG-REACT-NATIVE/service"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
	"github.com/google/wire"
	// _ "golang.org/x/xerrors"
	// _ "github.com/go-goracle/go-oracledb"
	//_ "gopkg.in/goracle.v2"
)

// func main() {

// 	db, err := sql.Open("godror", `user="akshay" password="password123" connectString="localhost:1521/ORCL"
//                                libDir="/Users/akshaybavalekar/Downloads/instantclient_19_8"`)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	dbQuery := "CREATE TABLE User(ID VARCHAR2(20), USERNAME VARCHAR2(20),MOBILENUMBER VARCHAR2(20))"

// 	rows, err := db.Query(dbQuery)
// 	if err != nil {
// 		fmt.Println(".....Error processing query")
// 		fmt.Println(err)
// 		return
// 	}
// 	defer rows.Close()
// 	fmt.Println("... Connected ")

// }

// import (
// 	"os"

// 	"github.com/gin-gonic/gin"
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"

// 	"rest-gin-gorm/product"
// )

func initDB() *sql.DB {
	db, err := sql.Open("godror", `user="akshay" password="password123" connectString="localhost:1521/ORCL"
                               libDir="/Users/akshaybavalekar/Downloads/instantclient_19_8"`)
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&product.Product{})

	return db
}

func initProductAPI(db *sql.DB) controller.ProductAPI {
	wire.Build(repositories.ProvideProductRepostiory, service.ProvideProductService, controller.ProvideProductAPI)

	return controller.ProductAPI{}
}

func main() {
	db := initDB()
	defer db.Close()

	productAPI := initProductAPI(db)

	r := gin.Default()

	r.POST("/Users", productAPI.Create)
	// r.GET("/products/:id", productAPI.FindByID)
	// r.POST("/products", productAPI.Create)
	// r.PUT("/products/:id", productAPI.Update)
	// r.DELETE("/products/:id", productAPI.Delete)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
