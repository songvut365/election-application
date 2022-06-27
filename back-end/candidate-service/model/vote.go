package model

type Vote struct {
	ID          string `json:"ID"`
	NationalID  string `json:"NationalID"`
	CandidateID int    `json:"CandidateID"`
}
