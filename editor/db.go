package main

import (
	"strings"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

const (
	playgroundColName = "playground"
)

type database struct {
	name          string
	playgroundCol *mgo.Collection
	session       *mgo.Session
}

var db *database

type sketch struct {
	name    string
	content string
}

func findSketch(attr string, val interface{}) (*sketch, error) {
	s := &sketch{}
	attr = strings.ToLower(attr)
	if attr == "id" {
		attr = "_id"
	}
	err := db.playgroundCol.Find(bson.M{attr: val}).One(s)
	if err == mgo.ErrNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return s, nil
}

func findSketchByName(name string) (*sketch, error) {
	return findSketch("name", name)
}
