package controller

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/utils"
	"ququ.im/gosearch/service"
	"strconv"
)

// IncrSuggest	godoc
// @Summary		增加搜索关键字一次或新增关键字
// @Description	增加搜索关键字一次或新增关键字
// @Tags	热搜补全
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	database formData string true "数据库名称"
// @Param	table formData string true "表名称"
// @Param	keyword formData string true "搜索关键字"
// @Param	customize formData string false "个性化定制标识编码，如用户id、文件id等，若不传则为同一张表共用的热搜词"
// @Param	incr formData int false "关键字增量值"
// @Success 200 {string} string	"ok"
// @Router	/gosearch/suggest/incr [post]
func IncrSuggest(params map[string]string) mgresult.Result {
	incr := 1
	if utils.Exists(params, "incr") {
		incr, _ = strconv.Atoi(params["incr"])
	}
	return service.IncrSuggest(params["database"], params["table"], params["customize"], params["keyword"], incr)
}

// ListSearchSuggest	godoc
// @Summary		获取下拉补全列表或热搜词
// @Description	获取下拉补全列表或热搜词
// @Tags	热搜补全
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	database formData string true "数据库名称"
// @Param	table formData string true "表名称"
// @Param	customize formData string false "个性化定制标识编码，如用户id、文件id等，若不传则为同一张表共用的热搜词"
// @Param	prefix formData string false "关键字前缀，若为空则返回热搜列表"
// @Param	size formData int false "返回记录数，默认10条"
// @Success 200 {string} string	"ok"
// @Router	/gosearch/suggest [post]
func ListSearchSuggest(params map[string]string) mgresult.Result {
	size := 10
	if utils.Exists(params, "size") {
		size, _ = strconv.Atoi(params["size"])
	}
	return service.ListSearchSuggest(params["database"], params["table"], params["customize"], params["prefix"], size)
}

// DeleteSuggest	godoc
// @Summary		删除关键字
// @Description	删除关键字
// @Tags	热搜补全
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	database formData string true "数据库名称"
// @Param	table formData string true "表名称"
// @Param	keyword formData string true "搜索关键字"
// @Param	customize formData string false "个性化定制标识编码，如用户id、文件id等，若不传则为同一张表共用的热搜词"
// @Success 200 {string} string	"ok"
// @Router	/gosearch/suggest/del [post]
func DeleteSuggest(params map[string]string) mgresult.Result {
	return service.DeleteSuggest(params["database"], params["table"], params["customize"], params["keyword"])
}
