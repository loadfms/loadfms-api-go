package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/loadfms/devices/controllers"
	"github.com/loadfms/devices/utils"
)

func main() {
	session := utils.GetSession()
	deviceController := controllers.NewDeviceController(session)

	mux := mux.NewRouter()
	mux.HandleFunc("/devices", deviceController.All()).Methods("GET")
	mux.HandleFunc("/devices", deviceController.New()).Methods("POST")
	mux.HandleFunc("/devices/{name}", deviceController.Update()).Methods("PUT")
	mux.HandleFunc("/devices/{name}/{state}", deviceController.UpdateState()).Methods("PUT")
	mux.HandleFunc("/devices/{port}", deviceController.DeviceByPort()).Methods("GET")

	http.ListenAndServe("localhost:8080", mux)
	log.Println("Server on port:8080")
}
