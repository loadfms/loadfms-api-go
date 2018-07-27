package controllers

import (
	"net/http"

	"github.com/loadfms/common/utils"
	"gopkg.in/mgo.v2"
)

type HealthController struct {
	session *mgo.Session
}

func NewHealthController(session *mgo.Session) *HealthController {
	return &HealthController{session}
}

func (healthController HealthController) Check() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		utils.MessageWithJSON(w, "ok", http.StatusOK)
	}
}
