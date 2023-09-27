package model

import (
	"gorm.io/gorm"
)

type AditionalInfo struct {
	gorm.Model
	Income            int    `json:"income"`
	ResponsableName   string `json:"responsableName"`
	NisNumber         string `json:"nisNumber"`
	ImageRight        bool   `json:"imageRight"`
	GovernmentBenefit bool   `json:"governmentBenefit"`
	HasBrother        bool   `json:"hasBrother"`
	StudentId         uint   `json:"studentId"`
}
