package api

import (
	"gisogd/SettingsService/internal/dal"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	routerGroup *gin.RouterGroup
	database dal.Orm
}

func (c *Controller) InitController (rg *gin.RouterGroup, path string, dbConnStr string) {
	c.routerGroup = rg.Group("/" + path)
	c.database = dal.Orm{}
	c.database.Init(dbConnStr)
}