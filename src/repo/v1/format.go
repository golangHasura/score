package v1

import (
	"score/src/database"
	"score/src/models"
	"score/src/repo"
)

type FormatRepo struct{}

func NewFormatRepo() repo.FormatRepo {
	return &FormatRepo{}
}

func (f *FormatRepo) FetchFormats(db *database.DB) ([]models.Format, error) {
	var formats []models.Format
	err := db.DBOp.Find(&formats).Error
	return formats, err
}
