package model

import "gorm.io/gorm"

type MotherInfo struct {
	gorm.Model
	StudentId        uint   `json:"studentId"`
	WorkOutside      bool   `json:"workOutside"`
	ReceiveBenefit   bool   `json:"receiveBenefit"`
	HasHelper        bool   `json:"hasHelper"`
	HelperName       string `json:"helperName"`
	ProjectName      string `json:"projectName"`
	AuthorizedPeople string `json:"authorizedPeople"`
}
