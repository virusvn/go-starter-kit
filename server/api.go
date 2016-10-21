package main

import (
	"github.com/virusvn/go-starter-kit/server/models"
	"gopkg.in/labstack/echo.v1"
	"gopkg.in/mgo.v2/bson"
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
	// Check title with current site is crawled or not
	// var base models.Base
	var article models.Article
	result, err := article.Find(article, bson.M{"site": "www.businesscard.vn"})
	if err != nil {
		panic(err)
	}
	c.JSON(200, result)
	return nil
}
