package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/fffnite/go-oneroster/helpers"
	"github.com/fffnite/go-oneroster/parameters"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

// Queries database connection for Orgs
func GetAllOrgs(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api := apiRequest{
			Request: Req{w, r},
			DB:      db,
			ORData:  OneRoster{Table: "orgs", Columns: publicCols, OutputName: "orgs"},
			Params:  parameters.Parameters{},
			Fks:     []FK{FK{"parentSourcedId", "orgs", "sourcedId", "sourcedId", "parent"}},
		}
		api.invoke()
	}
}

func GetOrg(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get object based off id from query
		id := chi.URLParam(r, "id")
		//var p parameters.Parameters
		/*
			api := apiRequest{
				Table:   "orgs",
				Columns: publicCols,
				Request: r,
				DB:      db,
				Params:  p,
			}
		*/
		statement := fmt.Sprintf("SELECT sourcedId, name FROM orgs WHERE sourcedId='%v'", id)

		var org Org
		db.QueryRow(statement).Scan(&org.SourcedId, &org.Name)
		//org["children"] = data.QueryNestedProperty("orgs", "parentSourcedId", org["sourcedId"], db)
		// Wrap result
		var output = struct {
			Org Org
		}{org}

		// Output result
		render.JSON(w, r, output)
	}
}

func GetMongoOrgs(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("oneroster").Collection("orgs")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		filter, err := helpers.GetFilters(r.URL.Query(), publicCols)
		if err != nil {
			log.Error(err)
		}
		options, errP := helpers.GetOptions(r.URL.Query(), publicCols)
		if errP != nil {
			log.Error(errP)
		}
		cur, err := collection.Find(
			ctx,
			filter,
			options,
		)
		if err != nil {
			log.Error(err)
		}
		defer cur.Close(ctx)
		var results []bson.M
		for cur.Next(ctx) {
			var result bson.M
			err := cur.Decode(&result)
			if err != nil {
				log.Error(err)
			}
			results = append(results, result)
		}
		render.JSON(w, r, results)
	}
}

func GetMongoOrg(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("oneroster").Collection("orgs")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		id := chi.URLParam(r, "id")
		cur, err := collection.Find(ctx, bson.D{{"sourcedId", id}})
		if err != nil {
			log.Info(err)
		}
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var result bson.M
			err := cur.Decode(&result)
			if err != nil {
				log.Error(err)
			}
			render.JSON(w, r, result)
		}
	}
}

func PutMongoOrg(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("oneroster").Collection("orgs")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		id := chi.URLParam(r, "id")
		var data PutOrg
		err := render.DecodeJSON(r.Body, &data)
		if err != nil {
			log.Info(err)
			// TODO: fix response
			render.JSON(w, r, err)
			return
		}
		data.DateLastModified = time.Now()
		res, err := collection.UpdateOne(ctx, bson.D{{"sourcedId", id}},
			bson.D{{"$set", data}},
			options.Update().SetUpsert(true))
		if err != nil {
			log.Info(err)
		}
		render.JSON(w, r, res)
	}
}

type PutOrg struct {
	SourcedId        string    `json:"sourcedId,omitempty" bson:"sourcedId,omitempty"`
	Status           string    `json:"status,omitempty" bson:"status,omitempty"`
	DateLastModified time.Time `json:"dateLastModified" bson:"dateLastModified"`
	Name             string    `json:"name,omitempty" bson:"name,omitempty"`
	Identifier       string    `json:"identifier,omitempty" bson:"identifier,omitempty"`
	Parent           struct {
		SourcedId string `json:"sourcedId,omitempty" bson:"sourcedId,omitempty"`
		Type      string `json:"type,omitempty" bson:"type,omitempty"`
	}
	Children []struct {
		SourcedId string `json:"sourcedId,omitempty" bson:"sourceId,omitempty"`
		Type      string `json:"type,omitempty" bson:"type,omitempty"`
	}
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user: bob"))
}

var publicCols = []string{"sourcedId",
	"status",
	"dateLastModified",
	"name",
	"type",
	"identifier",
	"parentSourcedId",
}

// JSON out per spec
type Org struct {
	SourcedId        string
	Status           string
	DateLastModified string
	Name             string
	Type             string
	Identifier       string
	Parent           struct {
		Href      string
		SourcedId string
		Type      string
	}
	Children []struct {
		Href      string
		SourcedId string
		Type      string
	}
}
