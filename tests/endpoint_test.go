package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	_ "typeset.io/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"typeset.io/models"
)

var (
	dummyPostID = "586cdb369dcec07eb5d6ba93"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestGetPost for successful return a post
func TestGetPost(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/post/"+dummyPostID, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGetPost", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Get Post Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

// TestGetAllPost for successful return all post
func TestGetAllPost(t *testing.T) {
	var result []models.Post
	r, _ := http.NewRequest("GET", "/api/post", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	json.Unmarshal(w.Body.Bytes(), &result)
	Convey("Subject: Test GetAllPost Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
