package model

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model
	Name       string `json:"Name"`
	DOB        string `json:"DOB"`
	BioLink    string `json:"BioLink" `
	ImageLink  string `json:"ImageLink"`
	Policy     string `json:"Policy"`
	VotedCount *uint  `json:"VotedCount" gorm:"default:0"`
}
