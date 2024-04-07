package controller

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"api/src/domain/adapters"
)

func ControllerGetUserById(rServer *gin.Context) {
	idStr := rServer.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		rServer.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	result, err := adapters.GetUser(id)
	if err != nil {
		rServer.JSON(http.StatusNotFound, gin.H{
			"value": false,
			"message": "Usuário não encontrado.",
		})
	}
	rServer.JSON(http.StatusOK, gin.H{
		"value": true,
		"data": result,
	})
}