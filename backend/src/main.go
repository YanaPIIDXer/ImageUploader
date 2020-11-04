package main

import (
	"github.com/gin-gonic/gin"

	"server/api"
	"server/api/http_method"
	"server/api/health_check"
)

func main() {
	r := gin.Default()

	api.RegisterAPI(r, "/", http_method.GET, new(health_check.HealthCheckAPI))
	r.Run(":3000")
}
