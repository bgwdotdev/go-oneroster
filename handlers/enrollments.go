package handlers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type enrollments struct {
	SourcedId        string    `json:"sourcedId" bson:"sourcedId,omitempty"`
	Status           string    `json:"status" bson:"status,omitempty"`
	DateLastModified time.Time `json:"dateLastModified" bson:"dateLastModified,omitempty"`
	User             *Nested   `json:"user" bson:"user,omitempty"`
	Class            *Nested   `json:"class" bson:"class,omitempty"`
	School           *Nested   `json:"school" bson:"school,omitempty"`
	Role             string    `json:"role" bson:"role,omitempty"`
	Primary          bool      `json:"primary" bson:"primary,omitempty"`
	BeginDate        string    `json:"beginDate" bson:"beginDate,omitempty"`
	EndDate          string    `json:"endDate" bson:"endDate,omitempty"`
}

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
		GetCollection(c, enrollmentsCols, w, r)
	}
}

func GetEnrollments(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("enrollments")
		GetDoc(c, enrollmentsCols, w, r)
	}
}

func PutEnrollments(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("enrollments")
		var data enrollments
		data.DateLastModified = time.Now()
		PutDoc(c, &data, w, r)
	}
}
