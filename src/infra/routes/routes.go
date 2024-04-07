package routes

import (
	"github.com/gin-gonic/gin"
)
// 🎯 Criando a var router para ser usada depois em outras parte do código. Ao muda-a sem mudar em outros arquivos, pode surgir alguns erros
var router *gin.Engine

func StartRoute(httpServer *gin.Engine, rStart string) {

	router = httpServer
	// ♟️ O nome da rota precisa ser igual ao nome criado no settings.go, caso isso não aconteça a rota nunca irá ser criada.
	for _, nameroute := range ParamsRoutes {
		if nameroute.Name == rStart {nameroute.Fun()}
	}
}


