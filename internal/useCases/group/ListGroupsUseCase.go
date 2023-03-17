package group

import (
	"github.com/gabrielgl9/removies/internal/domain"
)

type ListGroupOutputDto struct {
	Id string
	Name string
}

type ListGroupsUseCase struct {
	listGroupRepository domain.IGroupRepository
}


func (listGroupUseCase *ListGroupsUseCase) NewListGroupsUseCase(listGroupRepository domain.IGroupRepository) {
	listGroupRepository.NewGroupRepository()
	listGroupUseCase.listGroupRepository = listGroupRepository
}

func (listGroupUseCase *ListGroupsUseCase) Handle() ([]*ListGroupOutputDto, error) {
	groups, err := listGroupUseCase.listGroupRepository.GetAll()

	if err != nil {
		return nil, err
	}

	var listGroupsOutput []*ListGroupOutputDto

	for _, group := range groups {
		listGroupsOutput = append(listGroupsOutput, &ListGroupOutputDto{
			Id: group.Id.String(),
			Name: group.Name,
		})
	}
	
	return listGroupsOutput, nil
}