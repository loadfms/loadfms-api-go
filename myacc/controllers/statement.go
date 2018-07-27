package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/loadfms/common/utils"
	"github.com/loadfms/myacc/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type StatementController struct {
	session *mgo.Session
}

func NewStatementController(session *mgo.Session) *StatementController {
	return &StatementController{session}
}

func (statementController StatementController) All() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c := statementController.session.DB("myacc").C("statement")

		var items []models.Statement
		err := c.Find(bson.M{}).All(&items)
		if err != nil {
			utils.LogError(err)
			log.Println("Failed get all statements: ", err)
			return
		}

		respBody, err := json.MarshalIndent(items, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		utils.ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

func (statementController StatementController) Add() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var item models.Statement
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&item)
		if err != nil {
			utils.LogError(err)
			return
		}

		c := statementController.session.DB("myacc").C("statement")

		err = c.Insert(item)
		if err != nil {
			utils.LogError(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}
