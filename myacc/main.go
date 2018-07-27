package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/loadfms/common/utils"
	"github.com/loadfms/myacc/controllers"
)

func main() {
	session := utils.GetSession()
	healthController := controllers.NewHealthController(session)
	statementController := controllers.NewStatementController(session)

	mux := mux.NewRouter()
	mux.HandleFunc("/api/myacc/health", healthController.Check()).Methods("GET")

	mux.HandleFunc("/api/myacc/statement", statementController.All()).Methods("GET")
	mux.HandleFunc("/api/myacc/statement", statementController.Add()).Methods("POST")

	http.ListenAndServe(":8080", mux)
	log.Println("Server on port:8080")
}
