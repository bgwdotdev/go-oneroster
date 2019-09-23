package handlers

import (
	"database/sql"
	"fmt"
	"github.com/fffnite/go-oneroster/parameters"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/mongo"
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
		GetCollection(collection, publicCols, w, r)
	}
}

func GetMongoOrg(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("oneroster").Collection("orgs")
		GetDoc(collection, publicCols, w, r)
	}
}

func PutMongoOrg(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("oneroster").Collection("orgs")
		var data PutOrg
		data.DateLastModified = time.Now()
		PutDoc(collection, &data, w, r)
	}
}

type PutOrg struct {
	SourcedId        string    `json:"sourcedId" bson:"sourcedId"`
	Status           string    `json:"status" bson:"status"`
	DateLastModified time.Time `json:"dateLastModified" bson:"dateLastModified"`
	Name             string    `json:"name" bson:"name"`
	Identifier       string    `json:"identifier" bson:"identifier"`
	Parent           struct {
		SourcedId string `json:"sourcedId" bson:"sourcedId"`
		Type      string `json:"type" bson:"type"`
	} `json:"parent" bson:"parent"`
	Children []struct {
		SourcedId string `json:"sourcedId" bson:"sourceId"`
		Type      string `json:"type" bson:"type"`
	} `json:"children" bson:"children"`
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
