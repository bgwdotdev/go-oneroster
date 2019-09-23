package handlers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

func GetAllOrgs(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("oneroster").Collection("orgs")
		GetCollection(collection, publicCols, w, r)
	}
}

func GetOrg(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("oneroster").Collection("orgs")
		GetDoc(collection, publicCols, w, r)
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
	Identifier       string    `json:"identifier" bson:"identifier,omitempty"`
	Parent           struct {
		SourcedId string `json:"sourcedId" bson:"sourcedId,omitempty"`
		Type      string `json:"type" bson:"type,omitempty"`
	} `json:"parent" bson:"parent,omitempty"`
	Children []struct {
		SourcedId string `json:"sourcedId" bson:"sourceId,omitempty"`
		Type      string `json:"type" bson:"type,omitempty"`
	} `json:"children" bson:"children,omitempty"`
}

var publicCols = []string{"sourcedId",
	"status",
	"dateLastModified",
	"name",
	"type",
	"identifier",
	"parentSourcedId",
}
