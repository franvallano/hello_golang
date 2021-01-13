package main

import(
	"gopkg.in/mgo.v2"
)	

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	
	ShowError(err)

	return session
}
