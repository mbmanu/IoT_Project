package main

import (
	"github.com/shujanpannag/iot_project_api/internal/database"
	"github.com/shujanpannag/iot_project_api/internal/model"
	"github.com/shujanpannag/iot_project_api/internal/util"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	// A router without any middleware
	r := gin.New()

	// Logger middleware
	f, log := util.GinLogger()
	defer f.Close()
	r.Use(log)

	// Recovery middleware
	r.Use(gin.Recovery())

	db := database.DbConnect()

	db.AutoMigrate(&model.Camlog{}, &model.Relationship{})

	v0 := r.Group("/v0")
	{

		v0.GET("/nodelog/:date", database.GetNodelog(db))
		v0.POST("/nodelog", database.PostNodelog(db))

		v0.GET("/userrel/:rel", database.GetUser(db))
		v0.POST("/userrel", database.PostUser(db))
		v0.PUT("/userrel/:name", database.PutUser(db))
		v0.DELETE("/userrel/:name", database.DelUser(db))

		v0.GET("/getlog", database.GetLogReport)
	}

	// Equivalent to 0.0.0.0:8080
	r.Run()
}
