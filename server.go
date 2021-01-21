package main

import (
	"database/sql"

	"GOLANG-REACT-NATIVE/controller"
	"GOLANG-REACT-NATIVE/repositories"
	"GOLANG-REACT-NATIVE/service"

	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
	"github.com/google/wire"
)

var db *sql.DB

func initDB() *sql.DB {
	db, err := sql.Open("godror", `user="akshay" password="password123" connectString="localhost:1521/ORCL"
                               libDir="/Users/akshaybavalekar/Downloads/instantclient_19_8"`)
	if err != nil {
		panic(err)
	}
	return db
}

func initProductAPI(db *sql.DB) controller.UserAPI {
	wire.Build(repositories.ProvideUserRepostiory, service.ProvideUserService, controller.ProvideUserAPI)
	return controller.UserAPI{}
}

func main() {
	db := initDB()
	defer db.Close()

	userAPI := initProductAPI(db)

	r := gin.Default()

	r.GET("/showUsers", userAPI.ShowUsers)
	r.GET("/showIssues", userAPI.ShowIssues)
	r.POST("/Users", userAPI.CreateUser)
	r.POST("/issues", userAPI.CreateIssues)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
