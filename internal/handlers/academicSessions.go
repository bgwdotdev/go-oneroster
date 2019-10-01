package handlers

import (
	"github.com/fffnite/go-oneroster/ormodel"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var asCols = []string{
	"sourcedId",
	"status",
	"dateLastModified",
	"title",
	"startDate",
	"endDate",
	"type",
	"parent",
	"children",
	"schoolYear",
}

func GetAllAcademicSessions(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("academicSessions")
		res, errP := GetCollection(c, asCols, w, r)
		out := struct {
			Output       []bson.M `json:"academicSessions,omitempty"`
			ErrorPayload []error  `json:"statusInfoSet,omitempty"`
		}{res, errP}
		render.JSON(w, r, out)

	}
}

func GetAcademicSession(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("academicSessions")
		res, errP := GetDoc(c, asCols, w, r)
		out := struct {
			Output       bson.M  `json:"academicSession,omitempty"`
			ErrorPayload []error `json:"statusInfoSet,omitempty"`
		}{res, errP}
		render.JSON(w, r, out)
	}
}

func PutAcademicSession(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("academicSessions")
		var data ormodel.AcademicSessions
		data.DateLastModified = time.Now()
		PutDoc(c, &data, w, r)
	}
}
