package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/loadfms/common/utils"
	"github.com/loadfms/myacc/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const DB = "myacc"
const COLLECTION = "statement"

type StatementController struct {
	session *mgo.Session
}

func NewStatementController(session *mgo.Session) *StatementController {
	return &StatementController{session}
}

func (statementController StatementController) All() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c := statementController.session.DB(DB).C(COLLECTION)

		var items []models.Statement
		err := c.Find(bson.M{}).All(&items)
		if err != nil {
			utils.LogError(err)
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

		c := statementController.session.DB(DB).C(COLLECTION)
		result := statementController.checkDuplicate(item, c)
		w.Header().Set("Content-Type", "application/json")

		if !result {
			err = c.Insert(item)
			if err != nil {
				utils.LogError(err)
				return
			}

			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func (statementController StatementController) checkDuplicate(item models.Statement, c *mgo.Collection) bool {

	var items []models.Statement
	err := c.Find(bson.M{"description": item.Description, "date": item.Date, "value": item.Value}).All(&items)
	if err != nil {
		utils.LogError(err)
	}
	if len(items) > 0 {
		return true
	}
	return false
}
