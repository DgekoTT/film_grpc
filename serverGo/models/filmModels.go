package models

import "gorm.io/gorm"

type Film struct {
	gorm.Model
	FilmName       string `gorm:"size:64"`
	ProductionYear uint16
	Genres         []*Genre `gorm:"many2many:FilmGenre;constraint:OnDelete:CASCADE"`
}
