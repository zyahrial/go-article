package models

import (
  "time"
  "github.com/kamva/mgm/v3"
  "go.mongodb.org/mongo-driver/mongo/options"
   "go.mongodb.org/mongo-driver/bson"
		
  "fmt"
  "context"
  "go.mongodb.org/mongo-driver/mongo"
)

func init() {
	credential := options.Credential{
		Username: "article",
		Password: "article",
	}
	
	err := mgm.SetDefaultConfig(nil,"admin", options.Client().ApplyURI("mongodb://" + "localhost:27017" + "/" + "admin" + "?authSource=admin").SetAuth(credential))

	if err != nil {
		fmt.Printf("failed connection")
	}

	AddIndex(bson.D{{"tittle", "text"},{"body", "text"}}, bson.D{{"tittle", 5}, {"body", 5}}) // to descending set it to -1
}

func AddIndex(indexKeys interface{}, indexWeight interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Coll := mgm.Coll(&Article{})
	
    _, err := Coll.Indexes().CreateOne(ctx, mongo.IndexModel{Keys: indexKeys, Options: options.Index().SetWeights(indexWeight)})
    if err != nil {
        return err
    }
    return nil
}

type ShowArticle struct {
	Id_ 		string      `json:"_id" bson:"_id" binding:"required"`	
	Id 			int64      `json:"id" bson:"id" binding:"required"`	
	Author      string      `json:"author" bson:"author" binding:"required"`
	Tittle      string      `json:"tittle" bson:"tittle" binding:"required"`
	Body        string      `json:"body" bson:"body" binding:"required"`
	CreatedAt       time.Time        `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at" bson:"updated_at"`
}

type Article struct {
	mgm.DefaultModel `bson:",inline"`
	Id 		  int64      `json:"id" bson:"id" binding:"required"`	
	Author    string      `json:"author" bson:"author" binding:"required"`
	Tittle    string      `json:"tittle" bson:"tittle" binding:"required"`
	Body      string      `json:"body" bson:"body" binding:"required"`
	Created_at time.Time
	Updated_at time.Time
}

func NewArticle(id int64, author string, tittle string, body string) *Article{
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	return &Article{
	   Id:  id,
	   Author: author,
	   Tittle: tittle,
	   Body: body,
	   Created_at: now,
	}
}