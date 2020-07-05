package models

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/milamice62/fakeapi/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client = db.ClientInit()

type Genres struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
}

func GetGenres(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var genres []Genres

	genresCollection := client.Database("mydatabase").Collection("genres")
	cur, err := genresCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var genre Genres
		// & character returns the memory address of the following variable.
		err := cur.Decode(&genre) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}
		// add item our array
		genres = append(genres, genre)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(res).Encode(genres) // encode similar to serialize process.

}

func AddGenre(res http.ResponseWriter, req *http.Request) {

}

func UpdateGenre(res http.ResponseWriter, req *http.Request) {

}

func FindGenre(res http.ResponseWriter, req *http.Request) {

}

func DeleteGenre(res http.ResponseWriter, req *http.Request) {

}
