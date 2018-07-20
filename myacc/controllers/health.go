package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/loadfms/common/utils"
)

type HealthController struct {
	session *sql.DB
}

func NewHealthController(session *sql.DB) *HealthController {
	return &HealthController{session}
}

func (healthController HealthController) Check() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var result string
		rows, err := healthController.session.Query("select 'ok'")

		if err != nil {
			utils.LogError(err)
		}

		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&result)
			if err != nil {
				utils.LogError(err)
				log.Fatal(err)
			}
		}
		err = rows.Err()

		if err != nil {
			utils.LogError(err)
		}
		utils.MessageWithJSON(w, result, http.StatusOK)
	}
}
