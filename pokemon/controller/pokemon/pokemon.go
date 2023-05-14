package pokemon

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pokemon/model/pokemon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

const connectString = "mongodb+srv://someshkumar71524:DbpxF0m8cTfBl9Iw@cluster1.339khvl.mongodb.net/?retryWrites=true&w=majority"
const dbName = "myproject"
const colName = "pokemon"

func init() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(connectString))
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection(colName)
}

func GetAllPokemon(c *gin.Context) {
	var pokemons []pokemon.Pokemon
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var pokemon pokemon.Pokemon
		err := cur.Decode(&pokemon)
		if err != nil {
			log.Fatal(err)
		}
		pokemons = append(pokemons, pokemon)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, pokemons)
}

func GetPokemon(c *gin.Context) {
	var pokemon pokemon.Pokemon
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&pokemon)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, pokemon)
}

func CreatePokemon(c *gin.Context) {
	var pokemon pokemon.Pokemon
	c.BindJSON(&pokemon)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, pokemon)
	if err != nil {
		log.Fatal(err)
	}
	pokemon.ID = result.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusCreated, pokemon)
}

func UpdatePokemon(c *gin.Context) {
	idString := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		log.Fatal(err)
	}

	var updatedPokemon pokemon.Pokemon
	var inputPokemon pokemon.Pokemon
	c.BindJSON(&inputPokemon)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": inputPokemon}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	if err := collection.FindOneAndUpdate(context.Background(), filter, update, options).Decode(&updatedPokemon); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "pokemon": updatedPokemon})
}

func DeletePokemon(c *gin.Context) {
	id := c.Param("id")
	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
