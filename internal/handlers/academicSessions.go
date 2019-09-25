package handlers

import (
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type academicSessions struct {
	SourcedId        string    `json:"sourcedId" bson:"sourcedId,omitempty"`
	Status           string    `json:"status" bson:"status,omitempty"`
	DateLastModified time.Time `json:"dateLastModified" bson:"dateLastModified,omitempty"`
	Title            string    `json:"title" bson:"title,omitempty"`
	StartDate        string    `json:"startDate" bson:"startDate,omitempty"`
	EndDate          string    `json:"endDate" bson:"endDate,omitempty"`
	Type             string    `json:"type" bson:"type,omitempty"`
	Parent           *Nested   `json:"parent" bson:"parent,omitempty"`
	Children         []*Nested `json:"children" bson:"children,omitempty"`
	SchoolYear       string    `json:"schoolYear" bson:"schoolYear,omitempty"`
}

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
		var data academicSessions
		data.DateLastModified = time.Now()
		PutDoc(c, &data, w, r)
	}
}
