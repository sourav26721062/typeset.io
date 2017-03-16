package db

import (
	"fmt"

	"github.com/astaxie/beego"
	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session

func init() {
	Connect()
}

//GetConnection - Keeping session singleton
func GetConnection() *mgo.Session {
	return session.Clone()
}

//Connect - Creating Mongo Session
func Connect() *mgo.Session {
	url := beego.AppConfig.String("mongourl")

	sess, err := mgo.Dial(url)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Unable to connect to mongo instance!")
	}

	session = sess
	return session

}
