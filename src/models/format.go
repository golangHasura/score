package models

type Format struct {
	EntityId int    `gorm:"primaryKey;type:int"`
	Name     string `gorm:"unique"`
}

func (Format) TableName() string {
	return "formats"
}
