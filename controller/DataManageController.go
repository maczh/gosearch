package controller

import (
	"github.com/maczh/gintool/mgresult"
	"ququ.im/gosearch/service"
)

// AddDocument	godoc
// @Summary		插入文档到全文索引库API
// @Description	插入文档到全文索引库<br>说明：1.插入的doc文档只允许一层结构，不要做嵌套<br>2.允许字符串数组字段参与全文检索<br>3.所有包含有中文的字段将自动生成一个加Spy后缀的拼音首字母字段，用于支持拼音首字母检索，不输出在搜索结果中<br>4.数据插入后将自动生成一个名为id的uuid值字段，原有id字段将会被覆盖
// @Tags	数据维护
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	database formData string true "数据库名称"
// @Param	table formData string false "表名称"
// @Param	doc formData string true "文档数据,JSON类型"
// @Param	searchFields formData string false "模糊搜索字段列表,JSON数组格式"
// @Success 200 {string} string	"ok"
// @Router	/gosearch/add [post]
func AddDocument(params map[string]string) mgresult.Result {
	return service.AddDocument(params["database"], params["table"], params["doc"], params["searchFields"])
}

// AddDocuments	godoc
// @Summary		批量插入文档到全文索引库API
// @Description	批量插入文档到全文索引库
// @Tags	数据维护
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	database formData string true "数据库名称"
// @Param	table formData string false "表名称"
// @Param	docs formData string true "文档数据数组,JSON数组类型"
// @Param	searchFields formData string false "模糊搜索字段列表,JSON数组格式"
// @Success 200 {string} string	"ok"
// @Router	/gosearch/add/batch [post]
func AddDocuments(params map[string]string) mgresult.Result {
	return service.AddDocuments(params["database"], params["table"], params["docs"], params["searchFields"])
}

// DeleteDocument	godoc
// @Summary		从全文索引库中删除文档API
// @Description 从全文索引库中删除文档
// @Tags	数据维护
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	database formData string true "数据库名称"
// @Param	table formData string true "表名称"
// @Param	id formData string true "文档id,在插入时返回的id值"
// @Success 200 {string} string	"ok"
// @Router	/gosearch/del [post]
func DeleteDocument(params map[string]string) mgresult.Result {
	return service.DeleteDocument(params["database"], params["table"], params["id"])
}

// DeleteDocuments	godoc
// @Summary		从全文索引库中批量删除文档API
// @Description 从全文索引库中批量删除文档
// @Tags	数据维护
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	database formData string true "数据库名称"
// @Param	table formData string false "表名称"
// @Param	ids formData string true "文档id列表,JSON数组格式"
// @Success 200 {string} string	"ok"
// @Router	/gosearch/del/batch [post]
func DeleteDocuments(params map[string]string) mgresult.Result {
	return service.DeleteDocuments(params["database"], params["table"], params["ids"])
}

// DeleteDocumentByQuery	godoc
// @Summary		按条件从全文索引库中删除文档API
// @Description 按条件从全文索引库中删除文档
// @Tags	数据维护
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	database formData string true "数据库名称"
// @Param	table formData string false "表名称"
// @Param	query formData string false "查找相等条件,JSON格式，精确查找"
// @Param	in formData string false "一个字段符合多个值之一，支持字符串、int、float类型的字段，相当于SQL的in,类似 AND field in ('value1','value2'),JSON格式，{字段名:[值1,值2...]},可多字段"
// @Success 200 {string} string	"ok"
// @Router	/gosearch/del/query [post]
func DeleteDocumentByQuery(params map[string]string) mgresult.Result {
	return service.DeleteDocumentByQuery(params["database"], params["table"], params["query"], params["in"])
}

// UpdateDocument	godoc
// @Summary		从全文索引库中更新文档API
// @Description 从全文索引库中更新文档
// @Tags	数据维护
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	database formData string true "数据库名称"
// @Param	table formData string false "表名称"
// @Param	id formData string true "文档id,在插入时返回的id值"
// @Param	update formData string true "要更新的字段内容,JSON格式"
// @Success 200 {string} string	"ok"
// @Router	/gosearch/update [post]
func UpdateDocument(params map[string]string) mgresult.Result {
	return service.UpdateDocument(params["database"], params["table"], params["id"], params["update"])
}

// UpdateDocumentByQuery	godoc
// @Summary		按条件从全文索引库中更新文档API
// @Description 按条件从全文索引库中更新文档
// @Tags	数据维护
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	database formData string true "数据库名称"
// @Param	table formData string false "表名称"
// @Param	query formData string true "查找条件,JSON格式，精确查找"
// @Param	update formData string true "要更新的字段内容,JSON格式"
// @Success 200 {string} string	"ok"
// @Router	/gosearch/update/query [post]
func UpdateDocumentByQuery(params map[string]string) mgresult.Result {
	return service.UpdateDocumentByQuery(params["database"], params["table"], params["query"], params["update"])
}

// DeleteTable	godoc
// @Summary		从全文索引库中删除表API
// @Description 从全文索引库中删除表
// @Tags	数据维护
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	database formData string true "数据库名称"
// @Param	table formData string false "表名称"
// @Success 200 {string} string	"ok"
// @Router	/gosearch/del/table [post]
func DeleteTable(params map[string]string) mgresult.Result {
	return service.DeleteTable(params["database"], params["table"])
}

// DeleteDatabase	godoc
// @Summary		删除数据库API
// @Description 删除数据库
// @Tags	数据维护
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	database formData string true "数据库名称"
// @Success 200 {string} string	"ok"
// @Router	/gosearch/del/database [post]
func DeleteDatabase(params map[string]string) mgresult.Result {
	return service.DeleteDatabase(params["database"])
}
