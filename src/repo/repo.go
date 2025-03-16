package repo

import (
	"score/src/database"
	"score/src/models"
)

type FormatRepo interface {
	FetchFormats(db *database.DB) ([]models.Format, error)
}
