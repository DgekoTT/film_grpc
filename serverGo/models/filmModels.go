package models

type Film struct {
	ID             uint64 `gorm:"primary_key;autoIncrement;unique"`
	FilmName       string `gorm:"size:64"`
	ProductionYear uint32
	Genres         []*Genre `gorm:"many2many:FilmGenre;constraint:OnDelete:CASCADE"`
}
