package model

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name          string    `json:"name"`
	Gender        string    `json:"gender"`
	BirthDate     time.Time `json:"birth_date"`
	Phone         string    `json:"phone"`
	ClassYear     string    `json:"class_year"`
	TimeImdz      string    `json:"time_imdz"`
	AditionalInfo AditionalInfo
	MotherInfo    MotherInfo
}
