package models

import (
  "time"
  // "encoding/json"
)

type ShowArticle struct {
	ID 			int64      `json:"id" bson:"id" binding:"required"`	
	Author      string      `json:"author" bson:"author" binding:"required"`
	Tittle      string      `json:"tittle" bson:"tittle" binding:"required"`
	Body        string      `json:"body" bson:"body" binding:"required"`
	Created_at       time.Time        `json:"created_at" bson:"created_at"`
	Updated_at       time.Time        `json:"updated_at" bson:"updated_at"`
}

type Article struct {
  	ID        uint        `gorm:"id,primary_key"`
	Author    string      `gorm:"size:255"`
	Tittle    string      `gorm:"size:1000"`
	Body      string      `gorm:"size:2000"`
	Created_at time.Time
	Updated_at time.Time
}

func (b *Article) TableName() string {
	return "article"
}