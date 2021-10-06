package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	FirstName   string             `bson:"firstName,omitempty"`
	PhoneNumber string             `bson:"phoneNumber,omitempty"`
	GoogleId    string             `bson:"googleId,omitempty"`
	FacebookId  string             `bson:"facebookId,omitempty"`
	InstagramId string             `bson:"instagramId,omitempty"`
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer client.Disconnect(ctx)

	database := client.Database("account")
	userCollection := database.Collection("users")

	user := User{
		FirstName: "The Polyglot Developer",
	}

	insertResult, _ := userCollection.InsertOne(ctx, user)
	fmt.Println(insertResult.InsertedID)

	//
	var found bson.M
	userCollection.FindOne(
		context.TODO(),
		bson.D{{"_id", insertResult.InsertedID}},
	).Decode(&found)
	fmt.Println(found)
	//
	//
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/api/user/signin", signin)
	mux.HandleFunc("/api/user", getMe)

	log.Println("Starting server on :4000")
	error1 := http.ListenAndServe(":4000", mux)
	log.Fatal(error1)
}
