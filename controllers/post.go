package controllers

import (
	"encoding/json"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"time"

	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"typeset.io/models"
	"typeset.io/repository"
)

var (
	err             error
	request         models.Request
	content         models.Paragraph
	contents        []models.Paragraph
	post            models.Post
	posts           []models.Post
	result          bool
	responseSuccess models.Response
	responseFail    models.Response
)

func init() {
	responseSuccess.Status = "Sucess"
	responseFail.Status = "Fail"
}

//PostController - define struct to beego
type PostController struct {
	beego.Controller
}

// @Title Get()
// @Description get post based on Id
// @Success 200 {int}
// @Failure 400 id is invalid
// @Failure 404 id is not present
// @router /api/post [get]
func (_post *PostController) Get() {
	pid := _post.GetString(":id")
	if pid != "" {
		post, er := repository.GetPost(pid)
		if er != nil {
			responseFail.Error = er.Error()
			_post.Data["json"] = responseFail
		} else {
			if post.ID != "" {
				_post.Data["json"] = post
			} else {
				_post.Ctx.ResponseWriter.WriteHeader(404)
				responseFail.Error = "post doesnot exist"
				_post.Data["json"] = responseFail
			}
		}
	} else {
		_post.Ctx.ResponseWriter.WriteHeader(400)
		responseFail.Error = "post id is invalid"
		_post.Data["json"] = responseFail
	}
	_post.ServeJSON()
}

// @Title GetAll()
// @Description get all post
// @Success 200 {int}
// @Failure 400 id is invalid
// @Failure 404 id is not present
// @router /api/post/:id [get]
func (_post *PostController) GetAll() {
	var param int
	queryParam := _post.GetString("page")
	if len(queryParam) > 0 && queryParam != "0" {
		param, err = strconv.Atoi(queryParam)
		if err != nil {
			responseFail.Error = err.Error()
			_post.Data["json"] = responseFail
			_post.ServeJSON()
			return
		}
	} else {
		param = 0
	}
	posts, err = repository.GetAllPost(param)
	if err != nil {
		responseFail.Error = err.Error()
		_post.Data["json"] = responseFail
	}
	_post.Data["json"] = posts
	_post.Ctx.ResponseWriter.WriteHeader(200)
	_post.ServeJSON()
}

// @Title Post()
// @Description create a post
// @Success 200 {int}
// @Failure 400 request body is invalid
// @router /api/post [post]
func (_post *PostController) Post() {
	var contents []models.Paragraph
	err = json.Unmarshal(_post.Ctx.Input.RequestBody, &request)
	if err != nil {
		logs.Error("Error during unmarshal")
		responseFail.Error = err.Error()
		_post.Data["json"] = responseFail
		_post.Ctx.ResponseWriter.WriteHeader(400)
		_post.ServeJSON()
		return
	}

	resultParagraph := strings.Split(request.Content, "/n/n")

	for _, value := range resultParagraph {
		var content models.Paragraph
		content.ID = bson.NewObjectId()
		content.Text = value
		content.CreatedOn = time.Now()
		contents = append(contents, content)
	}

	post.Title = request.Title
	post.Content = contents

	_, err = repository.CreatePost(post)
	if err != nil {
		responseFail.Error = err.Error()
		_post.Data["json"] = responseFail
		_post.Ctx.ResponseWriter.WriteHeader(400)
		_post.ServeJSON()
	}
	_post.Data["json"] = responseSuccess
	_post.Ctx.ResponseWriter.WriteHeader(200)
	_post.ServeJSON()
}

// @Title PostComment()
// @Description create a comment for a Paragraph
// @Success 200 {int}
// @Failure 400 request body is invalid
// @router /api/para/:pid/comment [post]
func (_post *PostController) PostComment() {
	var comment models.Comment
	pid := _post.GetString(":pid")
	err = json.Unmarshal(_post.Ctx.Input.RequestBody, &comment)
	if comment.Text == "" {
		responseFail.Error = "No Blank Space"
		_post.Data["json"] = responseFail
		_post.Ctx.ResponseWriter.WriteHeader(400)
		_post.ServeJSON()
		return
	}
	if err != nil {
		responseFail.Error = err.Error()
		_post.Data["json"] = responseFail
		_post.Ctx.ResponseWriter.WriteHeader(400)
		_post.ServeJSON()
		return
	}
	_, err := repository.PostComment(pid, comment)
	if err != nil {
		responseFail.Error = err.Error()
		_post.Data["json"] = responseFail
		_post.Ctx.ResponseWriter.WriteHeader(400)
		_post.ServeJSON()
		return
	}

	_post.Data["json"] = responseSuccess
	_post.Ctx.ResponseWriter.WriteHeader(200)
	_post.ServeJSON()
}

// @Title GetComment()
// @Description get all post
// @Success 200 {int}
// @router /api/para/:id/comment [get]
func (_post *PostController) GetComment() {
	var paragraph []models.Paragraph
	pid := _post.GetString(":pid")
	paragraph, err = repository.GetComment(pid)
	if err != nil {
		responseFail.Error = err.Error()
		_post.Data["json"] = responseFail
	}
	_post.Data["json"] = paragraph
	_post.Ctx.ResponseWriter.WriteHeader(200)
	_post.ServeJSON()
}
