package v1

import (
	"score/src/database"
	"score/src/models"
	"score/src/repo"
)

type TeamRepo struct{}

func NewTeamRepo() repo.TeamRepo {
	return &TeamRepo{}
}

func (f *TeamRepo) FetchTeams(db *database.DB) ([]models.Team, error) {
	var formats []models.Team
	err := db.DBOp.Preload("Format").Find(&formats).Error
	return formats, err
}
