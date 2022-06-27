package model

type ToggleInput struct {
	Enable bool `json:"enable"`
}

type Count struct {
	ID         int  `json:"ID"`
	VotedCount uint `json:"VotedCount"`
}

type ResultCandidate struct {
	ID         int     `json:"ID"`
	Name       string  `json:"Name"`
	DOB        string  `json:"DOB"`
	BioLink    string  `json:"BioLink" `
	ImageLink  string  `json:"ImageLink"`
	Policy     string  `json:"Policy"`
	VotedCount *uint   `json:"VotedCount" gorm:"default:0"`
	Percentage *string `json:"Percentage"`
}
