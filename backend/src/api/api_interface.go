package api

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"server/api/http_method"
)

// APIインタフェース
type ApiInterface interface {
	// APIの呼び出し
	Call(context *gin.Context) interface {}
}

// APIクラス
type Api struct {
	request ApiInterface
}

// アクセス時の処理
func (this Api) OnRequest(context *gin.Context) {
	response := this.request.Call(context)
	context.JSON(http.StatusOK, response)
}

// APIを登録
func RegisterAPI(router *gin.Engine, endpoint string, method http_method.HttpMethod, request ApiInterface) {
	api := new(Api)
	api.request = request
	switch method {
		case http_method.POST:
			router.POST(endpoint, api.OnRequest)
		case http_method.GET:
			router.GET(endpoint, api.OnRequest)
	}
}
