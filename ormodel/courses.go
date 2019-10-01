package ormodel

import "time"

type Courses struct {
	SourcedId        string    `json:"sourcedId" bson:"sourcedId,omitempty"`
	Status           string    `json:"status" bson:"status,omitempty"`
	DateLastModified time.Time `json:"dateLastModified" bson:"dateLastModified,omitempty"`
	Title            string    `json:"title" bson:"title,omitempty"`
	SchoolYear       *Nested   `json:"schoolYear" bson:"schoolYear,omitempty"`
	CourseCode       string    `json:"coursecode" bson:"courseCode,omitempty"`
	Grades           []string  `json:"grades" bson:"grades,omitempty"`
	Subjects         []string  `json:"subjects" bson:"subjects,omitempty"`
	Org              *Nested   `json:"org" bson:"org,omitempty"`
	SubjectCodes     []string  `json:"subjectCodes" bson:"subjectCodes,omitempty"`
	Resources        []*Nested `json:"resources" bson:"resources,omitempty"`
}
