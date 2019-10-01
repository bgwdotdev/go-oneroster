package handlers

import (
	"github.com/fffnite/go-oneroster/ormodel"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var enrollmentsCols = []string{
	"sourcedId",
	"status",
	"dateLastModified",
	"user",
	"class",
	"school",
	"role",
	"primary",
	"beginDate",
	"endDate",
}

func GetAllEnrollments(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("enrollments")
		res, errP := GetCollection(c, enrollmentsCols, w, r)
		out := struct {
			Output       []bson.M `json:"enrollments,omitempty"`
			ErrorPayload []error  `json:"statusInfoSet,omitempty"`
		}{res, errP}
		render.JSON(w, r, out)
	}
}

func GetEnrollments(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("enrollments")
		res, errP := GetDoc(c, enrollmentsCols, w, r)
		out := struct {
			Output       bson.M  `json:"enrollment,omitempty"`
			ErrorPayload []error `json:"statusInfoSet,omitempty"`
		}{res, errP}
		render.JSON(w, r, out)
	}
}

func PutEnrollments(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("enrollments")
		var data ormodel.Enrollments
		data.DateLastModified = time.Now()
		PutDoc(c, &data, w, r)
	}
}
