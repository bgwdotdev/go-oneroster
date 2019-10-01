package ormodel

import "time"

type Nested struct {
	SourcedId string `json:"sourcedId" bson:"sourcedId,omitempty"`
	Type      string `json:"type" bson:"type,omitempty"`
}

type Classes struct {
	SourcedId        string    `json:"sourcedId" bson:"sourcedId,omitempty"`
	Status           string    `json:"status" bson:"status,omitempty"`
	DateLastModified time.Time `json:"dateLastModified,omitempty" bson:"dateLastModified,omitempty"`
	Title            string    `json:"title" bson:"title,omitempty"`
	ClassCode        string    `json:"classCode" bson:"classCode,omitempty"`
	ClassType        string    `json:"classType" bson:"classType,omitempty"`
	Location         string    `json:"location" bson:"location,omitempty"`
	Grades           []string  `json:"grades" bson:"grades,omitempty"`
	Subjects         []string  `json:"subjects" bson:"subjects,omitempty"`
	Course           *Nested   `json:"course" bson:"course,omitempty"`
	School           *Nested   `json:"school" bson:"school,omitempty"`
	Terms            []*Nested `json:"terms" bson:"terms,omitempty"`
	SubjectCodes     []string  `json:"subjectCodes" bson:"subjectCodes,omitempty"`
	Periods          []string  `json:"periods" bson:"periods,omitempty"`
	Resources        []*Nested `json:"resources" bson:"resources,omitempty"`
}
