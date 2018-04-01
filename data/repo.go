package data

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const collectionName = "userActivities"

var db *mgo.Database
var Dao = UserDao{}

type UserDao struct {
	Server   string
	Database string
}

func InitDB(dbHost string, dbName string) {
	log.Println("Initialing database connection")
	Dao.Server = dbHost
	Dao.Database = dbName
	Dao.connect()
}

func CloseDB() {
	log.Println("Closing database connection")
	db.Session.Close()
}

func (dao *UserDao) connect() {
	session, err := mgo.Dial(dao.Server)
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		log.Fatal(err)
	}
	db = session.Clone().DB(dao.Database)
}

func (dao *UserDao) Save(change UserChange) {
	log.Println("Saving user change data")
	c := db.C(collectionName)
	err := c.Insert(&change)
	if err != nil {
		log.Fatal(err)
	}
}

func (dao *UserDao) CountActivities(query Query) int {
	c := db.C(collectionName)
	total, err := c.Find(composeCondition(query)).Count()
	if err != nil {
		panic(err)
	}
	return total
}

func (dao *UserDao) QueryActivities(query Query) []UserChange {
	var result []UserChange
	c := db.C(collectionName)
	err := c.Find(composeCondition(query)).Skip(query.Start).Limit(query.Size).Sort("-time").All(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func composeCondition(query Query) bson.M {
	var filter []bson.M
	if len(query.UserId) > 0 {
		filter = append(filter, bson.M{"userId": query.UserId})
	}
	if len(query.Value) > 0 {
		filter = append(filter, bson.M{"field": query.Value})
	}
	condition := bson.M{}
	if len(filter) > 0 {
		condition = bson.M{"$or": filter}
	}
	return condition
}
