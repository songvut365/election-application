package model

type ToggleInput struct {
	Enable bool `json:"enable"`
}

type Count struct {
	ID         int  `json:"ID"`
	VotedCount uint `json:"VotedCount"`
}
