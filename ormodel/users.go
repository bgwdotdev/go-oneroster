package ormodel

import "time"

type NestedUid struct {
	Type       string `json:"type" bson:"type,omitempty"`
	Identifier string `json:"identifier" bson:"identifier,omitempty"`
}

type Users struct {
	SourcedId        string       `json:"sourcedId" bson:"sourcedId,omitempty"`
	Status           string       `json:"status" bson:"status,omitempty"`
	DateLastModified time.Time    `json:"dateLastModified" bson:"dateLastModified,omitempty"`
	Username         string       `json:"username" bson:"username,omitempty"`
	UserIds          []*NestedUid `json:"userIds" bson:"userIds,omitempty"`
	EnabledUser      bool         `json:"enabledUser" bson:"enabledUser,omitempty"`
	GivenName        string       `json:"givenName" bson:"givenName,omitempty"`
	FamilyName       string       `json:"familyName" bson:"familyName,omitempty"`
	MiddleName       string       `json:"middleName" bson:"middleName,omitempty"`
	Role             string       `json:"role" bson:"role,omitempty"`
	Identifier       string       `json:"identifier" bson:"identifier,omitempty"`
	Email            string       `json:"email" bson:"email,omitempty"`
	SMS              string       `json:"sms" bson:"sms,omitempty"`
	Phone            string       `json:"phone" bson:"phone,omitempty"`
	Agents           []*Nested    `json:"agents" bson:"agents,omitempty"`
	Orgs             []*Nested    `json:"orgs" bson:"orgs,omitempty"`
	Grades           []string     `json:"grades" bson:"grades,omitempty"`
	Password         string       `json:"password" bson:"password,omitempty"`
}
