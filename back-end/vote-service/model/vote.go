package model

import "time"

type Vote struct {
	ID          *string    `json:"ID,omitempty" bson:"_id,omitempty"`
	NationalID  string     `json:"NationalID"`
	CandidateID *int       `json:"CandidateID"`
	CreatedAt   time.Time  `json:"CreatedAt"`
	UpdatedAt   time.Time  `json:"UpdatedAt"`
	DeletedAt   *time.Time `json:"DeletedAt"`
}
