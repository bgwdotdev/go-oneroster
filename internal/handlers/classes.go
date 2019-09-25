package handlers

import (
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type classes struct {
	SourcedId        string    `json:"sourcedId" bson:"sourcedId,omitempty"`
	Status           string    `json:"status" bson:"status,omitempty"`
	DateLastModified time.Time `json:"dateLastModified" bson:"dateLastModified,omitempty"`
	Title            string    `json:"title" bson:"title,omitempty"`
	ClassCode        string    `json:"classCode" bson:"classCode,omitempty"`
	ClassType        string    `json:"classType" bson:"classType,omitempty"`
	Location         string    `json:"location" bson:"location,omitempty"`
	Grades           []string  `json:"grades" bson:"grades,omitempty"`
	Subjects         []string  `json:"subjects" bson:"subjects,omitempty"`
	Course           string    `json:"course" bson:"course,omitempty"`
	School           string    `json:"school" bson:"school,omitempty"`
	Terms            []*Nested `json:"terms" bson:"terms,omitempty"`
	SubjectCodes     []string  `json:"subjectCodes" bson:"subjectCodes,omitempty"`
	Periods          []string  `json:"periods" bson:"periods,omitempty"`
	Resources        []*Nested `json:"resources" bson:"resources,omitempty"`
}

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
		var data classes
		data.DateLastModified = time.Now()
		PutDoc(c, &data, w, r)
	}
}
