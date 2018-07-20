package models

import (
	"time"
)

type Statement struct {
	ID          int       `json:"id"`
	Category    Category  `json:"category"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Value       float64   `json:"value"`
	Type        int       `json:"type"`
	RedordDate  time.Time `json:"record_date"`
}
