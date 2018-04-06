package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/loadfms/models"
	"github.com/loadfms/utils"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DeviceController struct {
	session *mgo.Session
}

func NewDeviceController(session *mgo.Session) *DeviceController {
	return &DeviceController{session}
}

func (deviceController DeviceController) All() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c := deviceController.session.DB("load-api").C("devices")

		var devices []models.Device
		err := c.Find(bson.M{}).All(&devices)
		if err != nil {
			utils.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed get all devices: ", err)
			return
		}

		respBody, err := json.MarshalIndent(devices, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		utils.ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

func (deviceController DeviceController) DeviceByPort() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c := deviceController.session.DB("load-api").C("devices")
		params := mux.Vars(r)
		port, err := strconv.Atoi(params["port"])

		if err != nil {
			utils.ErrorWithJSON(w, "Invalid port", http.StatusInternalServerError)
			return
		}

		var device models.Device
		err = c.Find(bson.M{"port": port}).One(&device)
		if err != nil {
			utils.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed get device: ", err, " - ", params["port"])
			return
		}

		respBody, err := json.MarshalIndent(device, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		utils.ResponseWithJSON(w, respBody, http.StatusOK)
	}

}

func (deviceController DeviceController) New() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var device models.Device
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&device)
		if err != nil {
			utils.ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			return
		}

		c := deviceController.session.DB("load-api").C("devices")

		err = c.Insert(device)
		if err != nil {
			utils.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed insert device: ", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}

func (deviceController DeviceController) Update() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var device models.Device
		params := mux.Vars(r)
		name := params["name"]

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&device)
		if err != nil {
			utils.ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			return
		}

		c := deviceController.session.DB("load-api").C("devices")

		err = c.Update(bson.M{"name": name}, &device)
		if err != nil {
			utils.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed insert device: ", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}

func (deviceController DeviceController) UpdateState() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var device models.Device
		params := mux.Vars(r)
		name := params["name"]
		stateParam := params["state"]
		state := 0

		if stateParam == "on" {
			state = 1
		} else {
			state = 0
		}

		c := deviceController.session.DB("load-api").C("devices")

		err := c.Find(bson.M{"name": name}).One(&device)
		if err != nil {
			utils.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Not found: ", err)
			return
		}

		device.State = state
		err = c.Update(bson.M{"name": name}, &device)
		if err != nil {
			utils.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed insert device: ", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}
