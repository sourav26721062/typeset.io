package test

import (
	"path/filepath"
	"runtime"
	"testing"

	_ "typeset.io/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"typeset.io/models"
	"typeset.io/repository"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

//Test for DB for All Post
func TestGetAllPostRepo(t *testing.T) {
	var result []models.Post
	result, _ = repository.GetAllPost()

	Convey("Subject: TestGetAllJobRepo\n", t, func() {
		Convey("The Result Should Not Be Empty", func() {
			So(len(result), ShouldBeGreaterThan, 0)
		})
	})
}
