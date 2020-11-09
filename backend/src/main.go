package main

import (
	"github.com/gin-gonic/gin"

	"server/api"
	"server/api/http_method"
	"server/api/health_check"
	"server/api/upload"
)

func main() {
	r := gin.Default()

	api.RegisterAPI(r, "/", http_method.GET, new(health_check.HealthCheckAPI))
	api.RegisterAPI(r, "/upload", http_method.POST, new(upload.UploadAPI))
	r.Run(":3000")
}
