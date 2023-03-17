package routes

import (
	"github.com/gabrielgl9/removies/internal/infra/repositories/postgresql"
	"github.com/gabrielgl9/removies/internal/infra/web/controllers"
	"github.com/gin-gonic/gin"
)

var groupRepository postgresql.GroupRepository
var groupController controllers.GroupController

func SetupRoutes(router gin.IRouter) {
	groupController.NewGroupController(&groupRepository)

	router.GET("/groups", groupController.GetAll)

}