package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DefaultResult struct {
	Message string `json:"message"`
}

func MessageWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	dResult := DefaultResult{}
	dResult.Message = message

	respBody, err := json.MarshalIndent(dResult, "", "  ")
	if err != nil {
		LogError(err)
	}

	w.Write(respBody)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func LogError(err error) {
	fmt.Println("------")
	fmt.Println(err)
	fmt.Println("------")
}
