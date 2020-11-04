package health_check

import (
	"github.com/gin-gonic/gin"
)

// レスポンス
type responseType struct {
	// ヘルスチェックの結果
	ResultMessage string `json:"result_message"`
}

type HealthCheckAPI struct {}

// APIの呼び出し
func(this HealthCheckAPI) Call(request *gin.Context) interface {} {
	response := new(responseType)
	response.ResultMessage = "OK"
	return response
}
