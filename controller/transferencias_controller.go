package controller

import (
	"bytebank-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var transferencias = []models.Transferencia{
	{ID: 1, Valor: 200, Destino: "12345", Data: time.Now()},
	{ID: 2, Valor: 400, Destino: "12345", Data: time.Now()},
}

func GetTransferencias(c *gin.Context) {
	c.JSON(http.StatusOK, transferencias)
}

func InsertTransferencia(c *gin.Context) {
	var transferencia models.Transferencia
	if err := c.ShouldBindJSON(&transferencia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	id := len(transferencias) - 1
	data := time.Now()
	transferencia.Data = data
	transferencia.ID = id
	transferencias = append(transferencias, transferencia)
	c.JSON(http.StatusOK, transferencia)
}
