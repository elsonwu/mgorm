package model

import (
	"labix.org/v2/mgo"
)

var _database *Database

func GetDatabase() *Database {
	if _database == nil {
		_database = &Database{connectString: "127.0.0.1", dbName: "testcn10", collections: make(map[string]*mgo.Collection)}
	}
	return _database
}

type Database struct {
	db            *mgo.Database
	dbName        string
	connectString string
	collections   map[string]*mgo.Collection
}

func (self *Database) DB() *mgo.Database {
	if self.db == nil {
		sess, err := mgo.Dial(self.connectString)
		if err != nil {
			panic("connect database error")
		}
		sess.SetMode(mgo.Monotonic, true)
		self.db = sess.DB(self.dbName)
	}

	return self.db
}

func (self *Database) GetCollection(name string) *mgo.Collection {
	if collection, ok := self.collections[name]; ok {
		return collection
	}
	self.collections[name] = self.DB().C(name)
	return self.collections[name]
}
