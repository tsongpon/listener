package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const DBNAME = "redPlanet"
const COLLECTIONNAME = "userActivities"
var dbHost = GetDBHost()

func Save(change UserChange) {
	log.Println("Saving user change data")
	session, err := mgo.Dial(dbHost)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(DBNAME).C(COLLECTIONNAME)
	err = c.Insert(&change)
	if err != nil {
		log.Fatal(err)
	}
}

func CountActivities(query Query) int {
	session, err := mgo.Dial(dbHost)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(DBNAME).C(COLLECTIONNAME)

	total, err := c.Find(composeCondition(query)).Count()

	if err != nil {
		panic(err)
	}

	return total
}

func QueryActivities(query Query) []UserChange {
	var result []UserChange
	session, err := mgo.Dial(dbHost)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(DBNAME).C(COLLECTIONNAME)

	err = c.Find(composeCondition(query)).Limit(query.Size).All(&result)

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