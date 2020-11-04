package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// APIインタフェース
type ApiInterface interface {
	// APIの呼び出し
	Call(request gin.Context) interface {}
}

// APIクラス
type Api struct {
	request ApiInterface
}

// アクセス時の処理
func (this Api) OnRequest(request gin.Context) []byte {
	response := this.request
	responseJson, _ := json.Marshal(response)
	return responseJson
}

// APIを生成
func MakeAPI(request ApiInterface) *Api {
	api := new(Api)
	api.request = request
	return api
}
