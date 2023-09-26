package database

import (
	"context"
	"log"
	"os"
	"time"

	. "github.com/farimarwat/go-books/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var client *mongo.Client
var bookCollection *mongo.Collection
var COLLECTION = "Books"
func GetClient() *mongo.Client{
	uri := os.Getenv("DATABASE_URL")
	//getting context
	if client != nil {
		return client
	}
	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	//getting client
	client, err := mongo.Connect(ctx,options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func GetCollection(client *mongo.Client, collectioName string)*mongo.Collection{
	if bookCollection != nil {
		return bookCollection
	}
	bookCollection := client.Database("BookShop").Collection(collectioName)
	return bookCollection
}

func Disconnect(){
	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	if client == nil{
		return
	}
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

//Query database
func List_Books()[]Book{
	client := GetClient()
	bookCollection := GetCollection(client,COLLECTION)
	//mongo queries
	ctx, cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	var bookList []Book
	cursor, err := bookCollection.Find(ctx,bson.D{})
	defer cursor.Close(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	//Iterating through the book elements
	for cursor.Next(ctx){
		var book Book
		err := cursor.Decode(&book)
		if err != nil {
			log.Fatalln(err)
		}
		bookList = append(bookList, book)
	}
	
	return bookList
}
func Find_Book(name string)*Book{
	client := GetClient()
	bookCollection := GetCollection(client,COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	var book *Book
	filter := bson.D{{"name",name}}
	err := bookCollection.FindOne(ctx,filter).Decode(&book)
	if err != nil {
		return nil
	}
	return book
}
func Create_Book(book Book)string{
	client := GetClient()
	bookCollection := GetCollection(client,COLLECTION)
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	bookToPost := Book{
		Id: primitive.NewObjectID(),
		Name: book.Name,
		Price: book.Price,
	}
	result, err := bookCollection.InsertOne(ctx,bookToPost)
	if err != nil {
		return ""
	}
	return result.InsertedID.(primitive.ObjectID).Hex()
}
