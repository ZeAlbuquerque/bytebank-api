package models

import "time"

type Transferencia struct {
	ID      int       `json:"id"`
	Valor   float64   `json:"valor"`
	Destino string    `json:"destino"`
	Data    time.Time `json:"data"`
}
