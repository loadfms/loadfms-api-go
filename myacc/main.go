package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/loadfms/common/utils"
	"github.com/loadfms/myacc/controllers"
)

func main() {
	session := utils.GetMysqlSession()
	healthController := controllers.NewHealthController(session)
	categoryController := controllers.NewCategoryController(session)
	statementController := controllers.NewStatementController(session)

	mux := mux.NewRouter()
	mux.HandleFunc("/api/myacc/health", healthController.Check()).Methods("GET")

	mux.HandleFunc("/api/myacc/category", categoryController.All()).Methods("GET")
	mux.HandleFunc("/api/myacc/category", categoryController.Add()).Methods("POST")

	mux.HandleFunc("/api/myacc/statement", statementController.All()).Methods("GET")
	mux.HandleFunc("/api/myacc/statement", statementController.Add()).Methods("POST")

	http.ListenAndServe(":8080", mux)
	log.Println("Server on port:8080")
}
