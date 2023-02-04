/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2023/2/4 21:36
 * @Project_Name : scaffold
 * @File : router.go
 * @Software :GoLand
 */

package router

import (
	"demo/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "this is a simple scaffold")
	})

	return r
}
