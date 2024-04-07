package handler

import (
  "github.com/gin-gonic/gin"
	"api/src/infra/routes"
	"os"
	"strings"
)

func Initialize() {
	httpServer := gin.Default()

	files, err := os.ReadDir("./src/infra/routes")
	if err == nil {
		for _, isFile := range files {
			var nameFileRouter string
			nameFileRouter = isFile.Name()
			routes.StartRoute(httpServer, strings.Split(nameFileRouter, ".")[0])
		}
	}

	httpServer.Run()
}