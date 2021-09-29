package main

import (
	"github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	"github.com/maczh/gintool"
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/mgtrace"
	"github.com/maczh/utils"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"ququ.im/gosearch/controller"
	_ "ququ.im/gosearch/docs"
)

/**
统一路由映射入口
*/
func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	engine := gin.Default()

	//添加跟踪日志
	engine.Use(mgtrace.TraceId())

	//设置接口日志
	engine.Use(gintool.SetRequestLogger())
	//添加跨域处理
	engine.Use(gintool.Cors())

	//添加swagger支持
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//engine.GET("/docs/*any", swagger_skin.HandleReDoc)

	//处理全局异常
	engine.Use(nice.Recovery(recoveryHandler))

	//设置404返回的内容
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, *mgresult.Error(-1, "请求的方法不存在"))
	})

	var result mgresult.Result
	//添加所需的路由映射
	engine.Any("/gosearch/add", func(c *gin.Context) {
		result = controller.AddDocument(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/gosearch/add/batch", func(c *gin.Context) {
		result = controller.AddDocuments(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/gosearch/del", func(c *gin.Context) {
		result = controller.DeleteDocument(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/gosearch/del/batch", func(c *gin.Context) {
		result = controller.DeleteDocuments(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/gosearch/del/query", func(c *gin.Context) {
		result = controller.DeleteDocumentByQuery(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/gosearch/query", func(c *gin.Context) {
		result = controller.SearchDocument(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/gosearch/update", func(c *gin.Context) {
		result = controller.UpdateDocument(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/gosearch/update/query", func(c *gin.Context) {
		result = controller.UpdateDocumentByQuery(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/gosearch/del/table", func(c *gin.Context) {
		result = controller.DeleteTable(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/gosearch/del/database", func(c *gin.Context) {
		result = controller.DeleteDatabase(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	//热搜与下拉补全服务模块
	engine.Any("/gosearch/suggest/incr", func(c *gin.Context) {
		result = controller.IncrSuggest(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/gosearch/suggest", func(c *gin.Context) {
		result = controller.ListSearchSuggest(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/gosearch/suggest/del", func(c *gin.Context) {
		result = controller.DeleteSuggest(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})

	return engine
}

func recoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(http.StatusOK, *mgresult.Error(-1, "系统异常，请联系客服"))
}
