package postgresql

import (
	"database/sql"

	"github.com/gabrielgl9/removies/internal/domain"
	"github.com/gabrielgl9/removies/internal/infra/database/postgresql"
	"github.com/google/uuid"
)

type GroupRepository struct {
	db *sql.DB
}

func (groupRepository *GroupRepository) NewGroupRepository() {
	var configPostgre postgresql.ConfigPostgreDb
	configPostgre.NewConfigPostgreDb()

	db, err := configPostgre.SetupDB()
	if err != nil {
		panic(err)
	}

	groupRepository.db = db
}

func (groupRepository *GroupRepository) GetAll() ([]*domain.Group, error) {
	rows, err := groupRepository.db.Query("SELECT * FROM groups")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []*domain.Group

	for rows.Next() {
		var group domain.Group
		var stringUuid *string

		err = rows.Scan(&stringUuid, &group.Name)

		if err != nil {
			return nil, err
		}

		group.Id = uuid.MustParse(*stringUuid)
		
		groups = append(groups, &group)
	}

	return groups, nil
}