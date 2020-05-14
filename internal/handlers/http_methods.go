package handlers

import (
	"context"
	"github.com/fffnite/go-oneroster/internal/helpers"
	// "github.com/fffnite/go-oneroster/ormodel"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
	"time"
)

type Nested struct {
	SourcedId string `json:"sourcedId" bson:"sourcedId,omitempty"`
	Type      string `json:"type" bson:"type,omitempty"`
}

// Gets a collection of docs
func GetCollection(
	c *mongo.Collection, pf []string,
	w http.ResponseWriter, r *http.Request,
) ([]bson.M, []error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter, err := helpers.GetFilters(r.URL.Query(), pf)
	if err != nil {
		log.Error(err)
	}
	options, errP := helpers.GetOptions(r.URL.Query(), pf)
	cur, err := c.Find(
		ctx,
		filter,
		options,
	)
	if err != nil {
		log.Error(err)
	}
	defer cur.Close(ctx)
	totalCount, err := c.CountDocuments(
		ctx,
		filter,
	)
	if err != nil {
		log.Error(err)
	}
	var results []bson.M
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Error(err)
		}
		results = append(results, result)
	}
	w.Header().Set("X-Total-Count", strconv.FormatInt(totalCount, 10))
	w.Header().Set("Link", helpers.GetLinkHeaders(totalCount, r))
	return results, errP
}

// Gets a specific item based off the sourcedId
func GetDoc(
	c *mongo.Collection, pf []string,
	w http.ResponseWriter, r *http.Request,
) (bson.M, []error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"sourcedId", chi.URLParam(r, "id")}}
	options, errP := helpers.GetOption(r.URL.Query(), pf)
	cur := c.FindOne(
		ctx,
		filter,
		options,
	)
	var result bson.M
	err := cur.Decode(&result)
	if err != nil {
		log.Error(err)
	}
	return result, errP
}

/*
// Unsure how to range over type in interface -- generics? reflection? other?
// Upserts a large array of items based off each doc sourcedId
func PutBulk(c *mongo.Collection, data ormodel.Data,
	w http.ResponseWriter, r *http.Request) {

	err := render.DecodeJSON(r.Body, &data)
	if err != nil {
		log.Info(err)
		// TODO: fix reponse
		render.JSON(w, r, err)
		return
	}

	var queries []mongo.WriteModel
	for _, v := range data {
		q := mongo.NewUpdateOneModel()
		q.SetUpsert(true)
		q.SetFilter(bson.D{{"sourcedId", v.Id()}})
		q.SetUpdate(bson.D{{"$set", v}})
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
*/

// Upserts a specific item based off the sourcedId
func PutDoc(c *mongo.Collection, data interface{},
	w http.ResponseWriter, r *http.Request) {
	err := render.DecodeJSON(r.Body, &data)
	if err != nil {
		log.Info(err)
		// TODO: fix response
		render.JSON(w, r, err)
		return
	}
	filter := bson.D{{"sourcedId", chi.URLParam(r, "id")}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := c.UpdateOne(
		ctx,
		filter,
		bson.D{{"$set", data}},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Info(err)
	}
	render.JSON(w, r, res)
}

// Performs an upsert operation to a nested array
func PutNestedDoc(
	c *mongo.Collection, data interface{},
	obj string, field string,
	w http.ResponseWriter, r *http.Request,
) {
	err := render.DecodeJSON(r.Body, &data)
	if err != nil {
		log.Info(err)
		// TODO: fix response
		render.JSON(w, r, err)
		return
	}
	filter := bson.D{
		{"sourcedId", chi.URLParam(r, "id")},
		{obj + "." + field, chi.URLParam(r, "subId")},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	count, _ := c.CountDocuments(
		ctx,
		filter,
	)
	// update
	if count > 0 {
		res, err := c.UpdateOne(
			ctx,
			filter,
			bson.D{{"$set", bson.D{{obj + ".$", &data}}}},
		)
		if err != nil {
			// TODO: return 500?
			log.Info(err)
		}
		// TODO: return success update
		render.JSON(w, r, res)
		return
	}
	// insert
	res, err := c.UpdateOne(
		ctx,
		bson.D{{"sourcedId", chi.URLParam(r, "id")}},
		bson.D{{"$addToSet", bson.D{{obj, &data}}}},
	)
	if err != nil {
		log.Info(err)
		// TODO: return 500?
	}
	// TODO: return success insert
	render.JSON(w, r, res)
}
