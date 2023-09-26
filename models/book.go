package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id   	primitive.ObjectID 		`json:"id"`
	Name 	string 					`json:"name"`
	Price 	int						`json:"price"`
}