package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/loadfms/controllers"
	"github.com/loadfms/utils"
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
}
