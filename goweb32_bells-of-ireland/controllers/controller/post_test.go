package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goccy/go-json"
	"github.com/magiconair/properties/assert"

	"github.com/gin-gonic/gin"
)

/**
参考官网: https://gin-gonic.com/zh-cn/docs/testing/
*/

func TestCreateArticleHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	url := "/api/v1/post"
	router.POST(url, CreateArticleHandler)
	body := `{
	"community_id": 1234566,
	"title":"test",
	"content": "水浒传"
	}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//第一种
	//assert.Equal(t, 200, w.Code)
	//assert.Contains(t, w.Body.String(), "Token 未登录")

	//第二种
	respon := &ResponseData{}
	json.Unmarshal(w.Body.Bytes(), respon)
	assert.Equal(t, respon.Code, CodeNotLogin)
}
