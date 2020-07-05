package main

// type Sample struct {
// 	Name string
// }
import (
	example "github.com/milamice62/fakeapi/http"
)

func main() {
	example.HttpServer()
	// client, err := db.ClientInit()
	// if err != nil {
	// 	log.Fatalf("could not connect db: %v", err)
	// }

	// collection := client.Database("mydatabase").Collection("mycollection")
	// _, err = collection.InsertOne(context.Background(), Sample{"chris"})
	// if err != nil {
	// 	log.Fatalf("could not create database: %v", err)
	// }

	// cur, err := collection.Find(context.Background(), bson.D{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer cur.Close(context.Background())

	// var results []*Sample
	// for cur.Next(context.Background()) {
	// 	// To decode into a struct, use cursor.Decode()
	// 	var result Sample
	// 	err := cur.Decode(&result)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	results = append(results, &result)
	// }

	// if err := cur.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// for _, res := range results {
	// 	fmt.Println(res.Name)
	// }
}
