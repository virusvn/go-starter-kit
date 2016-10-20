package main

import (
	"gopkg.in/labstack/echo.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// API is a defined as struct bundle
// for api. Feel free to organize
// your app as you wish.
type API struct{}

// Vertex For test only, no meaning
type Vertex struct {
	X int
	Y int
}

// Bind attaches api routes
func (api *API) Bind(group *echo.Group) {

	group.Get("/v1/conf", api.ConfHandler)
	group.Get("/v2/conf", api.VertexHandler)
	group.Get("/v1/articles", api.ArticleHandler)

}

// ConfHandler handle the app config, for example
func (api *API) ConfHandler(c *echo.Context) error {
	app := c.Get("app").(*App)
	// <-time.After(time.Millisecond * 500)
	c.JSON(200, app.Conf.Root)
	return nil
}

// VertexHandler An example return a struct
func (api *API) VertexHandler(c *echo.Context) error {
	c.JSON(200, Vertex{1, 2})
	return nil
}

// ArticleHandler An example return an article
func (api *API) ArticleHandler(c *echo.Context) error {

	// database
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	collection := session.DB("blog").C("urls")
	// Check title with current site is crawled or not
	var article Article
	err = collection.Find(bson.M{"title": "Hệ thống điệp viên Huginn: Thế giới trong tầm tay - Fullstack Station", "site": "www.businesscard.vn"}).One(&article)
	if err != nil {
		panic(err)
	}
	c.JSON(200, article)
	return nil
}

// Article struct
type Article struct {
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
