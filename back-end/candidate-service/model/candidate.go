package model

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primary_key;autoIncrement"`
	Name       string `json:"name"`
	DOB        string `json:"dob"`
	BioLink    string `json:"bioLink" `
	ImageLink  string `json:"imageLink"`
	Policy     string `json:"policy"`
	VotedCount *uint  `json:"votedCount" gorm:"default:0"`
}
