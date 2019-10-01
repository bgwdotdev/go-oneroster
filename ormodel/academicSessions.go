package ormodel

import "time"

type AcademicSessions struct {
	SourcedId        string    `json:"sourcedId" bson:"sourcedId,omitempty"`
	Status           string    `json:"status" bson:"status,omitempty"`
	DateLastModified time.Time `json:"dateLastModified" bson:"dateLastModified,omitempty"`
	Title            string    `json:"title" bson:"title,omitempty"`
	StartDate        string    `json:"startDate" bson:"startDate,omitempty"`
	EndDate          string    `json:"endDate" bson:"endDate,omitempty"`
	Type             string    `json:"type" bson:"type,omitempty"`
	Parent           *Nested   `json:"parent" bson:"parent,omitempty"`
	Children         []*Nested `json:"children" bson:"children,omitempty"`
	SchoolYear       string    `json:"schoolYear" bson:"schoolYear,omitempty"`
}
