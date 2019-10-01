package ormodel

import "time"

type Enrollments struct {
	SourcedId        string    `json:"sourcedId" bson:"sourcedId,omitempty"`
	Status           string    `json:"status" bson:"status,omitempty"`
	DateLastModified time.Time `json:"dateLastModified" bson:"dateLastModified,omitempty"`
	User             *Nested   `json:"user" bson:"user,omitempty"`
	Class            *Nested   `json:"class" bson:"class,omitempty"`
	School           *Nested   `json:"school" bson:"school,omitempty"`
	Role             string    `json:"role" bson:"role,omitempty"`
	Primary          bool      `json:"primary" bson:"primary,omitempty"`
	BeginDate        string    `json:"beginDate" bson:"beginDate,omitempty"`
	EndDate          string    `json:"endDate" bson:"endDate,omitempty"`
}
