package routes

import (
	"api/src/domain/client/controller"
)


func GetUserRoute() {
	router.GET("/getuser/:id", controller.ControllerGetUserById)
}