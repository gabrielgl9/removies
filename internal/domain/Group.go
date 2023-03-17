package domain

import "github.com/google/uuid"

type IGroupRepository interface {
	NewGroupRepository()
	GetAll() ([]*Group, error)
}

type Group struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
}

func (group *Group) NewGroup(name string) {
	group.Id = uuid.New()
	group.Name = name
}

