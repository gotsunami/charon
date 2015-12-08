package main

import (
	"errors"
	"net/url"

	"labix.org/v2/mgo"
)

func playgroundIndexes() []mgo.Index {
	idxs := []mgo.Index{
		{
			Key:        []string{"_id"},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     false,
		},
	}
	return idxs
}

// dial connects to MongoDB, registers new session and setup
// collections.
func dial(uri string, conf *config) error {
	if conf == nil {
		return errors.New("nil app config")
	}

	u, err := url.Parse(uri)
	if err != nil {
		return err
	}
	if u.Path == "" {
		return errors.New("no database name specified!")
	}
	if u.Scheme == "" {
		return errors.New("missing mongodb:// scheme in url")
	}
	dbName := u.Path[1:]

	s, err := mgo.Dial(uri)
	if err != nil {
		return err
	}
	db = &database{name: dbName, session: s}

	// Setup collections
	db.playgroundCol = s.DB(dbName).C(playgroundColName)

	// Setup indexes
	for _, idx := range playgroundIndexes() {
		if err := db.playgroundCol.EnsureIndex(idx); err != nil {
			return err
		}
	}

	return nil
}
