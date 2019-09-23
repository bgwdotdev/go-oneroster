package handlers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type courses struct {
	SourcedId        string    `json:"sourcedId" bson:"sourcedId,omitempty"`
	Status           string    `json:"status" bson:"status,omitempty"`
	DateLastModified time.Time `json:"dateLastModified" bson:"dateLastModified,omitempty"`
	Title            string    `json:"title" bson:"title,omitempty"`
	SchoolYear       *Nested   `json:"schoolYear" bson:"schoolYear,omitempty"`
	CourseCode       string    `json:"coursecode" bson:"courseCode,omitempty"`
	Grades           []string  `json:"grades" bson:"grades,omitempty"`
	Subjects         []string  `json:"subjects" bson:"subjects,omitempty"`
	Org              *Nested   `json:"org" bson:"org,omitempty"`
	SubjectCodes     []string  `json:"subjectCodes" bson:"subjectCodes,omitempty"`
	Resources        []*Nested `json:"resources" bson:"resources,omitempty"`
}

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
		GetCollection(c, coursesCols, w, r)
	}
}

func GetCourses(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("courses")
		GetDoc(c, coursesCols, w, r)
	}
}

func PutCourses(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("courses")
		var data courses
		data.DateLastModified = time.Now()
		PutDoc(c, &data, w, r)
	}
}
