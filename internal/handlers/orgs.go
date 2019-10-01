package handlers

import (
	"github.com/fffnite/go-oneroster/ormodel"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

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

func GetAllOrgs(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("oneroster").Collection("orgs")
		res, errP := GetCollection(collection, orgCols, w, r)
		out := struct {
			Output       []bson.M `json:"orgs,omitempty"`
			ErrorPayload []error  `json:"statusInfoSet,omitempty"`
		}{res, errP}
		render.JSON(w, r, out)
	}
}

func GetOrg(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("oneroster").Collection("orgs")
		res, errP := GetDoc(collection, orgCols, w, r)
		out := struct {
			Output       bson.M  `json:"org,omitempty"`
			ErrorPayload []error `json:"statusInfoSet,omitempty"`
		}{res, errP}
		render.JSON(w, r, out)
	}
}

func PutOrg(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("oneroster").Collection("orgs")
		var data ormodel.Orgs
		data.DateLastModified = time.Now()
		PutDoc(collection, &data, w, r)
	}
}
