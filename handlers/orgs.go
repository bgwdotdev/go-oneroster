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

// Gets a collection of docs
func GetCollection(c *mongo.Collection, pf []string,
	w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter, err := helpers.GetFilters(r.URL.Query(), pf)
	if err != nil {
		log.Error(err)
	}
	options, errP := helpers.GetOptions(r.URL.Query(), pf)
	if errP != nil {
		log.Error(errP)
	}
	cur, err := c.Find(
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

// Gets a specific item based off the sourcedId
func GetDoc(c *mongo.Collection, pf []string,
	w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"sourcedId", chi.URLParam(r, "id")}}
	options, errP := helpers.GetOptions(r.URL.Query(), pf)
	if errP != nil {
		log.Error(errP)
	}
	cur, err := c.Find(
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

// Upserts a specific item based off the sourcedId
func DecodeDoc(c *mongo.Collection, data interface{},
	w http.ResponseWriter, r *http.Request) {
	err := render.DecodeJSON(r.Body, &data)
	if err != nil {
		log.Info(err)
		// TODO: fix response
		render.JSON(w, r, err)
		return
	}
}

func PutDoc(c *mongo.Collection, data interface{},
	w http.ResponseWriter, r *http.Request) {
	filter := bson.D{{"sourcedId", chi.URLParam(r, "id")}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := c.UpdateOne(ctx, filter, bson.D{{"$set", data}}, options.Update().SetUpsert(true))
	if err != nil {
		log.Info(err)
	}
	render.JSON(w, r, res)
}

func PutMongoOrg(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("oneroster").Collection("orgs")
		var p PutOrg
		DecodeDoc(collection, p, w, r)
		p.DateLastModified = time.Now()
		PutDoc(collection, p, w, r)
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
	} `json:"parent,omitempty" bson:"parent,omitempty"`
	Children []struct {
		SourcedId string `json:"sourcedId,omitempty" bson:"sourceId,omitempty"`
		Type      string `json:"type,omitempty" bson:"type,omitempty"`
	} `json:"children,omitempty" bson:"children,omitempty"`
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
