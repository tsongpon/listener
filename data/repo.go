package data

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"github.com/tsongpon/listener/config"
)

const collectionName = "userActivities"
var db *mgo.Database
var Context = DBContext{}

type DBContext struct {
	Server   string
	Database string
}

func InitDB() {
	log.Println("Initialing database connection")
	Context.Server = config.GetDBHost()
	Context.Database = config.DatabaseName
	Context.Connect()
}

func CloseDB() {
	db.Session.Close()
}

func (ctx *DBContext) Connect() {
	session, err := mgo.Dial(ctx.Server)
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		log.Fatal(err)
	}
	db = session.Clone().DB(ctx.Database)
}

func (ctx *DBContext) Close() {
	db.Session.Close()
}

func (ctx *DBContext) Save(change UserChange) {
	log.Println("Saving user change data")
	c := db.C(collectionName)
	err := c.Insert(&change)
	if err != nil {
		log.Fatal(err)
	}
}

func (ctx *DBContext) CountActivities(query Query) int {
	c := db.C(collectionName)

	total, err := c.Find(composeCondition(query)).Count()

	if err != nil {
		panic(err)
	}

	return total
}

func (ctx *DBContext) QueryActivities(query Query) []UserChange {
	var result []UserChange
	c := db.C(collectionName)

	err := c.Find(composeCondition(query)).Skip(query.Start).Limit(query.Size).Sort("time").All(&result)

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
