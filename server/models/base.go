package models

import (
	"github.com/virusvn/go-starter-kit/server/db"
	//	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
)

// Base model
type Base struct {
}

// Get mongdb collection session
func newCollection(v interface{}) *db.Collection {
	return db.NewCollectionSession(strings.ToLower(reflect.TypeOf(v).Name() + "s"))
}

// Create is insert a record
func (b Base) Create(v interface{}) (interface{}, error) {
	var err error
	c := newCollection(v)
	defer c.Close()
	err = c.Session.Insert(&v)
	if err != nil {
		return v, err
	}
	return v, err
}

// Save not implemented yet
func (b Base) Save(v interface{}) (interface{}, error) {
	var err error
	c := newCollection(v)
	defer c.Close()
	err = c.Session.Insert(&v)
	if err != nil {
		return v, err
	}
	return v, err
}

// Delete not implemented yet
func (b Base) Delete(v interface{}) (interface{}, error) {
	var err error
	c := newCollection(v)
	defer c.Close()
	err = c.Session.Insert(&v)
	if err != nil {
		return v, err
	}
	return v, err
}

// Find a record
func (b Base) Find(v interface{}, q map[string]interface{}) (interface{}, error) {
	c := newCollection(v)
	defer c.Close()
	err := c.Session.Find(q).One(&v)
	if err != nil {
		return v, err
	}
	return v, err
}

// FindAll not implemented yet
func (b Base) FindAll(v interface{}) (interface{}, error) {
	var err error
	c := newCollection(v)
	defer c.Close()
	err = c.Session.Insert(&v)
	if err != nil {
		return v, err
	}
	return v, err
}
