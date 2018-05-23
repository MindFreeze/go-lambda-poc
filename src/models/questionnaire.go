package models

import (
	// "database/sql"
	_ "github.com/jinzhu/gorm"
	"gopkg.in/guregu/null.v3"
)

type Questionnaire struct {
	ID                  uint        `gorm:"primary_key" json:"id"`
	Libraryid           uint        `db:"libraryid" json:"libraryid"`
	Identificationid    string      `db:"identificationid" json:"identificationid"`
	Identificationvalue null.String `json:"identificationvalue"`
	Name                null.String `json:"name"`
	Entrynodeid         null.Int    `json:"entrynodeid"`
	// Questions []Question `gorm:"many2many:questionnaire_questions;" json:"questions"`
}
