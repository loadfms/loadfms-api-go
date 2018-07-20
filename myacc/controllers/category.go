package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/loadfms/common/utils"
	"github.com/loadfms/myacc/models"
)

type CategoryController struct {
	session *sql.DB
}

func NewCategoryController(session *sql.DB) *CategoryController {
	return &CategoryController{session}
}

func (categoryController CategoryController) All() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := categoryController.session.Query("select category_id, category_name from tb_category")

		if err != nil {
			utils.LogError(err)
		}

		defer rows.Close()
		var categories []models.Category
		for rows.Next() {
			var item models.Category
			err := rows.Scan(&item.ID, &item.Name)
			categories = append(categories, item)
			if err != nil {
				utils.LogError(err)
				log.Fatal(err)
			}
		}
		err = rows.Err()

		if err != nil {
			utils.LogError(err)
		}

		respBody, err := json.MarshalIndent(categories, "", "  ")
		if err != nil {
			utils.LogError(err)
		}

		utils.ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

func (categoryController CategoryController) Add() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := "insert into tb_category (category_name) values (?)"
		decoder := json.NewDecoder(r.Body)
		var t models.Category
		err := decoder.Decode(&t)
		if err != nil {
			utils.LogError(err)
		}

		_, err = categoryController.session.Exec(query, t.Name)

		if err != nil {
			utils.LogError(err)
		}
		utils.MessageWithJSON(w, "ok", http.StatusOK)

	}
}
