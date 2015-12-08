package main

import (
	"strings"
	"time"

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
	Name         string `bson:"_id" json:"id"`
	Content      string
	CreationDate *time.Time
}

func (s *sketch) create() error {
	t := time.Now().UTC()
	s.CreationDate = &t

	err := db.playgroundCol.Insert(s)
	if err != nil {
		return err
	}
	return nil
}

func (s *sketch) update() error {
	err := db.playgroundCol.Update(bson.M{"_id": s.Name}, s)
	if err != nil {
		return err
	}
	return nil
}

func (s *sketch) delete() error {
	return db.playgroundCol.Remove(bson.M{"_id": s.Name})
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
	return findSketch("id", name)
}
