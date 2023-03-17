package controllers

import (
	"net/http"

	"github.com/gabrielgl9/removies/internal/domain"
	"github.com/gabrielgl9/removies/internal/useCases/group"
	"github.com/gin-gonic/gin"
)

type GroupController struct {
	groupRepository domain.IGroupRepository
	listGroupsUseCase group.ListGroupsUseCase
}

func (groupController *GroupController) NewGroupController(groupRepository domain.IGroupRepository) {
	groupController.groupRepository = groupRepository
	
	var listGroupsUseCase group.ListGroupsUseCase
	listGroupsUseCase.NewListGroupsUseCase(groupController.groupRepository)
	groupController.listGroupsUseCase = listGroupsUseCase
}

func (groupController *GroupController) GetAll(c *gin.Context) {
	groups, err := groupController.listGroupsUseCase.Handle()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, groups)
}