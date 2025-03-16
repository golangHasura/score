package models

type Format struct {
	EntityId  int    `gorm:"entity_id"`
	Name      string `gorm:"unique"`
	NoOfBalls int    `gorm:"number_of_balls"`
}

func (Format) TableName() string {
	return "formats"
}
