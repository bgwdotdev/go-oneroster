package handlers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

func GetAllOrgs(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("oneroster").Collection("orgs")
		GetCollection(collection, orgCols, w, r)
	}
}

func GetOrg(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("oneroster").Collection("orgs")
		GetDoc(collection, orgCols, w, r)
	}
}

func PutOrg(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("oneroster").Collection("orgs")
		var data Org
		data.DateLastModified = time.Now()
		PutDoc(collection, &data, w, r)
	}
}

type Org struct {
	SourcedId        string    `json:"sourcedId" bson:"sourcedId,omitempty"`
	Status           string    `json:"status" bson:"status,omitempty"`
	DateLastModified time.Time `json:"dateLastModified" bson:"dateLastModified,omitempty"`
	Name             string    `json:"name" bson:"name,omitempty"`
	Type             string    `json:"type" bson:"type,omitempty"`
	Identifier       string    `json:"identifier" bson:"identifier,omitempty"`
	Parent           *Nested   `json:"parent" bson:"parent,omitempty"`
	Children         []*Nested `json:"children" bson:"children,omitempty"`
}

var orgCols = []string{
	"sourcedId",
	"status",
	"dateLastModified",
	"name",
	"type",
	"identifier",
	"parent",
	"children",
}
