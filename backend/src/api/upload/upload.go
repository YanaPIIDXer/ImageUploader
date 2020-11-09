package upload


import (
	"github.com/gin-gonic/gin"
)

// レスポンス
type responseType struct {
	// アップロード結果
	IsSuccess bool `json:"success"`
}

type UploadAPI struct {}

// APIの呼び出し
func(this UploadAPI) Call(request *gin.Context) interface {} {
	response := new(responseType)
	response.IsSuccess = true
	return response
}
