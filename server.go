package main

import (
	"GOLANG-REACT-NATIVE/controller"
	"GOLANG-REACT-NATIVE/repositories"
	"GOLANG-REACT-NATIVE/service"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
	"github.com/google/wire"
)

func initDB() *sql.DB {
	db, err := sql.Open("godror", `user="akshay" password="password123" connectString="localhost:1521/ORCL"
                               libDir="/Users/akshaybavalekar/Downloads/instantclient_19_8"`)
	if err != nil {
		panic(err)
	}
	return db
}

func initProductAPI(db *sql.DB) controller.ProductAPI {
	wire.Build(repositories.ProvideUserRepostiory, service.ProvideUserService, controller.ProvideUserAPI)

	return controller.ProductAPI{}
}

func main() {
	db := initDB()
	defer db.Close()

	userAPI := initProductAPI(db)

	r := gin.Default()

	r.POST("/Users", userAPI.Create)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
