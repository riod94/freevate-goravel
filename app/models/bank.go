package models

import (
	"github.com/goravel/framework/database/orm"
)

type Bank struct {
	orm.Model
	Name         string
	Code         string
	Status       string
	Category     string
	Link         string
	Contact      string
	LogoFilepath string
	orm.SoftDeletes
}
