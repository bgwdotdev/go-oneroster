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
	SourcedId        string    `json:"sourcedId" bson:"sourcedId"`
	Status           string    `json:"status" bson:"status"`
	DateLastModified time.Time `json:"dateLastModified" bson:"dateLastModified"`
	Name             string    `json:"name" bson:"name"`
	Identifier       string    `json:"identifier" bson:"identifier"`
	Parent           struct {
		SourcedId string `json:"sourcedId" bson:"sourcedId"`
		Type      string `json:"type" bson:"type"`
	} `json:"parent" bson:"parent"`
	Children []struct {
		SourcedId string `json:"sourcedId" bson:"sourceId"`
		Type      string `json:"type" bson:"type"`
	} `json:"children" bson:"children"`
}

var publicCols = []string{"sourcedId",
	"status",
	"dateLastModified",
	"name",
	"type",
	"identifier",
	"parentSourcedId",
}
