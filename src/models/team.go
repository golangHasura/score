package models

type Team struct {
	EntityId int    `gorm:"primaryKey;type:int"`
	Name     string `gorm:"unique;type:varchar(255)"`
	Logo     string `gorm:"index;type:varchar(255)"`
	FormatId int    `gorm:"index;type:int"`
	Format   Format `gorm:"foreignKey:FormatId; references:EntityId"`
}

func (Team) TableName() string {
	return "teams"
}
