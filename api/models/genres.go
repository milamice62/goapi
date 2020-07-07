package models

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
		fmt.Printf("error list genre: %v", err)
		return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var genre Genres
		// & character returns the memory address of the following variable.
		err := cur.Decode(&genre) // decode similar to deserialize process.
		if err != nil {
			fmt.Printf("error list genre: %v", err)
			return
		}
		// add item our array
		genres = append(genres, genre)
	}

	if err := cur.Err(); err != nil {
		fmt.Printf("error list genre: %v", err)
		return
	}

	json.NewEncoder(res).Encode(genres) // encode similar to serialize process.

}

func AddGenre(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	//Unmarshall request body to object
	var genre Genres
	err := json.NewDecoder(req.Body).Decode(&genre)
	if err != nil {
		fmt.Printf("error read genre body: %v", err)
		return
	}
	//Add genre to Collection
	genresCollection := client.Database("mydatabase").Collection("genres")
	result, err := genresCollection.InsertOne(context.TODO(), genre)

	if err != nil {
		fmt.Printf("error add genre: %v", err)
		return
	}

	json.NewEncoder(res).Encode(result)
}

func UpdateGenre(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	//Validate id
	var params = mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		fmt.Printf("Not valid id: %v", err)
		return
	}
	//Init genre collection
	var genre Genres
	genresCollection := client.Database("mydatabase").Collection("genres")
	//Unmarshall request body to object
	err = json.NewDecoder(req.Body).Decode(&genre)
	if err != nil {
		fmt.Printf("error read genre body: %v", err)
		return
	}
	//Find and update genre
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"name", genre.Name},
		}},
	}
	err = genresCollection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&genre)
	if err != nil {
		fmt.Printf("error update genre: %v", err)
		return
	}
	json.NewEncoder(res).Encode(genre)
}

func FindGenre(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	//Generate objectID via id extracted in parameter
	var params = mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		fmt.Printf("Not valid id: %v", err)
		return
	}
	//Find the specific genre
	var genre Genres
	filter := bson.M{
		"_id": id,
	}
	genresCollection := client.Database("mydatabase").Collection("genres")
	err = genresCollection.FindOne(context.TODO(), filter).Decode(&genre)
	if err != nil {
		json.NewEncoder(res).Encode(fmt.Sprintf("genre id %s not found", params["id"]))
		return
	}

	json.NewEncoder(res).Encode(genre)
}

func DeleteGenre(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	//Generate objectID via id extracted in parameter
	var params = mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		fmt.Printf("Not valid id: %v", err)
		return
	}
	//Find specific genre then delete
	var genre Genres
	filter := bson.M{
		"_id": id,
	}
	genresCollection := client.Database("mydatabase").Collection("genres")
	err = genresCollection.FindOneAndDelete(context.TODO(), filter).Decode(&genre)
	if err != nil {
		json.NewEncoder(res).Encode(fmt.Sprintf("genre id %s not found", params["id"]))
		return
	}

	json.NewEncoder(res).Encode(genre)
}
