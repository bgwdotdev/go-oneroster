// Package providing oauth2 authentication
// for go-oneroster-api service

package auth

import (
	"context"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fffnite/go-oneroster/internal/database"
	"github.com/go-chi/jwtauth"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Login(u, p string) (string, error) {
	client := database.ConnectDb()
	c := client.Database("credentials").Collection("clients")
	err := validateSecret(u, p, c)
	if err != nil {
		log.Infof("Bad login: %v", u)
		log.Error(err)
		//TODO: 401
		return "", err
	}
	t, err := createToken(u)
	if err != nil {
		// review
		log.Error(err)
		return "", err
	}
	return t, nil
}

type creds struct {
	ClientSecret string
}

func getSecret(id string, c *mongo.Collection) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur := c.FindOne(
		ctx,
		bson.D{{"clientId", id}},
	)
	var s creds
	cur.Decode(&s)
	return s.ClientSecret, nil
}

func validateSecret(u, p string, c *mongo.Collection) error {
	hash, err := getSecret(u, c)
	if err != nil {
		return err
	}
	h := []byte(hash)
	b := []byte(p)
	err = bcrypt.CompareHashAndPassword(h, b)
	if err != nil {
		return err
	}
	return nil
}

func createToken(u string) (string, error) {
	keyA := viper.GetString("auth_key_alg")
	key := viper.GetString("auth_key")
	tokenAuth := jwtauth.New(keyA, []byte(key), nil)
	t := jwt.MapClaims{"aud": u}
	jwtauth.SetIssuedNow(t)
	jwtauth.SetExpiryIn(t, (time.Duration(30 * time.Minute)))
	_, tokenString, err := tokenAuth.Encode(t)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
