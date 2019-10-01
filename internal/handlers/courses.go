package handlers

import (
	"github.com/fffnite/go-oneroster/ormodel"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var coursesCols = []string{
	"sourcedId",
	"status",
	"dateLastModified",
	"title",
	"schoolYear",
	"courseCode",
	"grades",
	"subjects",
	"org",
	"subjectCodes",
	"resources",
}

func GetAllCourses(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("courses")
		res, errP := GetCollection(c, coursesCols, w, r)
		out := struct {
			Output       []bson.M `json:"courses,omitempty"`
			ErrorPayload []error  `json:"statusInfoSet,omitempty"`
		}{res, errP}
		render.JSON(w, r, out)
	}
}

func GetCourses(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("courses")
		res, errP := GetDoc(c, coursesCols, w, r)
		out := struct {
			Output       bson.M  `json:"course,omitempty"`
			ErrorPayload []error `json:"statusInfoSet,omitempty"`
		}{res, errP}
		render.JSON(w, r, out)
	}
}

func PutCourses(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("courses")
		var data ormodel.Courses
		data.DateLastModified = time.Now()
		PutDoc(c, &data, w, r)
	}
}
