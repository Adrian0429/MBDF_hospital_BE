package entities

type Day_name struct {
	Day_Of_Week int    `gorm:"type:int;not null" json:"day_of_week"`
	Name        string `gorm:"type:varchar(20); not null" json:"name"`
}
