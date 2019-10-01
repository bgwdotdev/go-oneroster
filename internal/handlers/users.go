package handlers

import (
	"github.com/fffnite/go-oneroster/ormodel"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var userCols = []string{
	"sourcedId",
	"status",
	"dateLastModified",
	"username",
	"userIds",
	"enabledUser",
	"givenName",
	"familyName",
	"middleName",
	"role",
	"identifier",
	"email",
	"sms",
	"phone",
	"agents",
	"orgs",
	"grades",
	"password",
}

func GetAllUsers(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("users")
		res, errP := GetCollection(c, userCols, w, r)
		out := struct {
			Output       []bson.M `json:"users,omitempty"`
			ErrorPayload []error  `json:"statusInfoSet,omitempty"`
		}{res, errP}
		render.JSON(w, r, out)
	}
}

func GetUser(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("users")
		res, errP := GetDoc(c, userCols, w, r)
		out := struct {
			Output       bson.M  `json:"user,omitempty"`
			ErrorPayload []error `json:"statusInfoSet,omitempty"`
		}{res, errP}
		render.JSON(w, r, out)
	}
}

func PutUser(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("users")
		var data ormodel.Users
		data.DateLastModified = time.Now()
		PutDoc(c, &data, w, r)
	}
}

func PutUserId(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("users")
		var data ormodel.NestedUid
		PutNestedDoc(c, &data, "userIds", "type", w, r)
	}
}
