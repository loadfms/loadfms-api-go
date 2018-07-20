package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/loadfms/common/utils"
	"github.com/loadfms/myacc/models"
)

type StatementController struct {
	session *sql.DB
}

func NewStatementController(session *sql.DB) *StatementController {
	return &StatementController{session}
}

func (statementController StatementController) All() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := statementController.session.Query("select statement_id, statement_description, statement_date, statement_value, statement_type, c.category_id, c.category_name, record_date from tb_statement s inner join tb_category c on s.category_id = c.category_id")

		if err != nil {
			utils.LogError(err)
		}

		defer rows.Close()
		var statements []models.Statement
		for rows.Next() {
			var item models.Statement
			err := rows.Scan(
				&item.ID,
				&item.Description,
				&item.Date,
				&item.Value,
				&item.Type,
				&item.Category.ID,
				&item.Category.Name,
				&item.RedordDate,
			)
			statements = append(statements, item)
			if err != nil {
				utils.LogError(err)
				log.Fatal(err)
			}
		}
		err = rows.Err()

		if err != nil {
			utils.LogError(err)
		}

		respBody, err := json.MarshalIndent(statements, "", "  ")
		if err != nil {
			utils.LogError(err)
		}

		utils.ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

func (statementController StatementController) Add() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := "insert into tb_statement (statement_description, statement_date, statement_value, statement_type, category_id) values (?, ?, ?, ?, ?)"
		decoder := json.NewDecoder(r.Body)
		var t models.Statement
		err := decoder.Decode(&t)
		if err != nil {
			utils.LogError(err)
		}

		//TODO: Encontrar melhor categoria
		t.Category.ID = 1

		_, err = statementController.session.Exec(query, t.Description, t.Date, t.Value, t.Type, t.Category.ID)

		if err != nil {
			utils.LogError(err)
		}
		utils.MessageWithJSON(w, "ok", http.StatusOK)

	}
}
