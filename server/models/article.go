package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Article struct
type Article struct {
	Base
	Content string        `bson:"content,omitempty"`
	URL     string        `bson:"url,omitempty"`
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Slug    string        `bson:"slug,omitempty"`
	Title   string        `bson:"title,omitempty"`
	Images  []struct {
		URL  string `bson:"url,omitempty"`
		Size string `bson:"size,omitempty"`
	} `bson:"images,omitempty"`
	Technology  []string `bson:"technology,omitempty"`
	Seri        string   `bson:"seri,omitempty"`
	Language    []string `bson:"language,omitempty"`
	Website     string   `bson:"website,omitempty"`
	Domain      string   `bson:"domain,omitempty"`
	Topic       []string `bson:"topic,omitempty"`
	Locale      string   `bson:"locale,omitempty"`
	Framework   []string `bson:"framework,omitempty"`
	Description string   `bson:"description,omitempty"`
	Level       []int64  `bson:"level,omitempty"`
	Site        string   `bson:"site,omitempty"`
	Comments    []struct {
		Updated  time.Time `bson:"updated,omitempty"`
		Comment  string    `bson:"comment,omitempty"`
		Username string    `bson:"username,omitempty"`
		Created  time.Time `bson:"created,omitempty"`
	} `bson:"comments,omitempty"`
	Point   int64     `bson:"point,omitempty"`
	Updated time.Time `bson:"updated,omitempty"`
	Created time.Time `bson:"created,omitempty"`
}
