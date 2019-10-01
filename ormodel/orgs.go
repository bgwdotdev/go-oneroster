package ormodel

import "time"

type Orgs struct {
	SourcedId        string    `json:"sourcedId" bson:"sourcedId,omitempty"`
	Status           string    `json:"status" bson:"status,omitempty"`
	DateLastModified time.Time `json:"dateLastModified" bson:"dateLastModified,omitempty"`
	Name             string    `json:"name" bson:"name,omitempty"`
	Type             string    `json:"type" bson:"type,omitempty"`
	Identifier       string    `json:"identifier" bson:"identifier,omitempty"`
	Parent           *Nested   `json:"parent" bson:"parent,omitempty"`
	Children         []*Nested `json:"children" bson:"children,omitempty"`
}
