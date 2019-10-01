package handlers

import (
	"github.com/fffnite/go-oneroster/ormodel"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var classesCols = []string{
	"sourcedId",
	"status",
	"dateLastModified",
	"title",
	"classCode",
	"classType",
	"location",
	"grades",
	"subjects",
	"course",
	"school",
	"terms",
	"subjectCodes",
	"periods",
	"resources",
}

func GetAllClasses(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("classes")
		res, errP := GetCollection(c, classesCols, w, r)
		out := struct {
			Output       []bson.M `json:"classes,omitempty"`
			ErrorPayload []error  `json:"statusInfoSet,omitempty"`
		}{res, errP}
		render.JSON(w, r, out)
	}
}

func GetClasses(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("classes")
		res, errP := GetDoc(c, classesCols, w, r)
		out := struct {
			Output       bson.M  `json:"class,omitempty"`
			ErrorPayload []error `json:"statusInfoSet,omitempty"`
		}{res, errP}
		render.JSON(w, r, out)
	}
}

func PutClasses(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("classes")
		var data ormodel.Classes
		data.DateLastModified = time.Now()
		PutDoc(c, &data, w, r)
	}
}
