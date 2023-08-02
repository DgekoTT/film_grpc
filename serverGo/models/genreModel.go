package models

type Genre struct {
	ID        uint32 `gorm:"primary_key;autoIncrement;unique"`
	GenreName string `gorm:"size:16"`
}
