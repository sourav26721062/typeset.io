package repository

import (
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"gopkg.in/mgo.v2/bson"
	"typeset.io/db"
	"typeset.io/models"
)

var (
	mongoDB string
	err     error
	link    string
)

func init() {
	mongoDB = beego.AppConfig.String("mongodb")
	link = beego.AppConfig.String("postlink")
}

//CreatePost - Creating a post in DB
func CreatePost(post models.Post) (bool, error) {
	post.ID = bson.NewObjectId()
	post.CreatedOn = time.Now()
	post.Link = link + post.ID.Hex()
	session := db.GetConnection()
	defer session.Close()
	c := session.DB(mongoDB).C("posts")
	err = c.Insert(post)
	if err != nil {
		log.Println("Error creating Post: ", err.Error())
		return false, err
	}
	return true, nil
}

//GetPost - Get a post from DB based on id
func GetPost(postIdentity string) (models.Post, error) {
	session := db.GetConnection()
	defer session.Close()
	c := session.DB(mongoDB).C("posts")
	var post models.Post
	pid := bson.ObjectIdHex(postIdentity)
	err = c.FindId(pid).One(&post)
	if err != nil {
		log.Println("Error getting post: ", err.Error())
	}
	return post, nil

}

//GetAllPost - Get all post from DB
func GetAllPost(n int) ([]models.Post, error) {
	session := db.GetConnection()
	defer session.Close()
	c := session.DB(mongoDB).C("posts")
	var posts []models.Post

	if n == 0 {
		err = c.Find(bson.M{}).Sort("-createdOn").Limit(500).All(&posts)
	} else {
		result := c.Find(bson.M{}).Sort("-createdOn").Limit(5)
		result = result.Skip((n - 1) * 10)
		err = result.All(&posts)
	}

	if err != nil {
		log.Println("Error getting all Post: ", err.Error())
		return nil, err
	}
	return posts, nil
}

//PostComment - Create a comment for a specific paragraph in DB
func PostComment(paragraphIdentity string, comment models.Comment) (bool, error) {
	session := db.GetConnection()
	defer session.Close()
	c := session.DB(mongoDB).C("posts")
	pid := bson.ObjectIdHex(paragraphIdentity)
	query := bson.M{"content._id": pid}
	comment.ID = bson.NewObjectId()
	comment.CreatedOn = time.Now()
	update := bson.M{"$push": bson.M{"content.$.comments": comment}}

	err = c.Update(query, update)
	if err != nil {
		log.Println("Error getting all Post: ", err.Error())
		return false, err
	}
	return true, nil
}

//PostComment - Create a comment for a specific paragraph in DB
func GetComment(paragraphIdentity string) ([]models.Paragraph, error) {
	var post models.Post
	session := db.GetConnection()
	defer session.Close()
	c := session.DB(mongoDB).C("posts")
	pid := bson.ObjectIdHex(paragraphIdentity)
	query := bson.M{"content._id": pid}
	err = c.Find(query).One(&post)
	logs.Info(post)
	if err != nil {
		log.Println("Error getting all Post: ", err.Error())
		return post.Content, err
	}
	return post.Content, nil
}
