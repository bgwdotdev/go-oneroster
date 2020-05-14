package handlers

import (
	"context"
	"github.com/fffnite/go-oneroster/ormodel"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
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

func PutBulkUser(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := client.Database("oneroster").Collection("users")
		var data []ormodel.Userwrap
		// PutBulk(c, &data, w, r)

		err := render.DecodeJSON(r.Body, &data)
		if err != nil {
			log.Error(err)
			render.JSON(w, r, err)
			return
		}

		var queries []mongo.WriteModel
		for _, v := range data {
			q := mongo.NewUpdateOneModel()
			q.SetUpsert(true)
			q.SetFilter(bson.D{{"sourcedId", v.User.SourcedId}})
			q.SetUpdate(bson.D{{"$set", v.User}})
			queries = append(queries, q)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		res, err := c.BulkWrite(ctx, queries)
		if err != nil {
			log.Error(err)
		}
		render.JSON(w, r, res)
	}
}
