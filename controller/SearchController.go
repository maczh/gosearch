package controller

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/utils"
	"ququ.im/gosearch/service"
	"strconv"
)

// SearchDocument	godoc
// @Summary		按组合条件从全文索引库中搜索文档API
// @Description 按组合条件从全文索引库中搜索文档<br>说明:1.输入中文内容的搜索建议放在mustWildcard中，将同时自动生成一个拼音首字母放在mustWildcard中<br>2.由于没有分词，所以中文将自动拆字，只要包含中文串中的一个字都会命中，因此同时自动加上简拼搜索会大幅度提升命中率<br>3.多个组合条件精确查找的也都放在must当中，比如shopId=xxxx and userId=xxxx and data like '%xx%'<br>4.如果需要更精确的结果，建议在搜索结果中再次进行一遍过滤
// @Tags	全文检索
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	database formData string true "数据库名称"
// @Param	table formData string true "表名称"
// @Param	should formData string false "OR相等查找条件,JSON格式，{字段名:值},即field1=value1 OR field2=value2,可多字段"
// @Param	shouldWildcard formData string false "OR模糊搜索，相当于"*关键字*",类似 OR field like '%keyword%',JSON格式，{字段名:关键字},可多字段"
// @Param	must formData string false "AND相等查找条件,JSON格式，{字段名:值},即field1=value1 AND field2=value2,可多字段，若某个字段的值为JSON数组，则相当于in查询"
// @Param	mustWildcard formData string false "AND模糊搜索，相当于*关键字*,类似 AND field like '%keyword%',JSON格式，{字段名:关键字},可多字段"
// @Param	not formData string false "不相等,JSON格式，{字段名:值},类似 AND field <> value,可多字段"
// @Param	notWildcard formData string false "不包含，相当于*关键字*,类似 AND NOT (field like '%keyword%'),JSON格式，{字段名:关键字},可多字段"
// @Param	in formData string false "一个字段符合多个值之一，支持字符串、int、float类型的字段，相当于SQL的in,类似 AND field in ('value1','value2'),JSON格式，{字段名:[值1,值2...]},可多字段"
// @Param	contains formData string false "一个字符串数组字段符合多个值之一,仅支持字符串数组类型字段，原始数据集合与查询数据集合有交集则匹配,JSON格式，{字段名:[值1,值2...]},可多字段"
// @Param	range formData string false "按>/>=/</<=搜索，类似MongoDB，gt/gte/lt/lte,JSON格式，如:{字段名:{"gte":值1,"lt":值2}},可多字段"
// @Param	sort formData string false "搜索结果排序方式，字段名JSON数组，字段名前加+号代表正排序，加-号代表倒排序"
// @Param	offset formData int false "分页偏移量，从第几条开始的偏移量，默认为0"
// @Param	size formData int false "分页的单页记录数，默认不分页"
// @Param	almost formData int false "是否准精确搜索 1-是 0-否， 默认为1"
// @Success 200 {string} string	"ok"
// @Router	/gosearch/query [post]
func SearchDocument(params map[string]string) mgresult.Result {
	offset := 0
	size := 0
	almost := 1
	if utils.Exists(params, "size") {
		size, _ = strconv.Atoi(params["size"])
	}
	if utils.Exists(params, "offset") {
		offset, _ = strconv.Atoi(params["offset"])
	}
	if utils.Exists(params, "almost") {
		almost, _ = strconv.Atoi(params["almost"])
	}
	return service.SearchDocument(params["database"], params["table"], params["should"], params["must"], params["not"], params["shouldWildcard"], params["mustWildcard"], params["notWildcard"], params["range"], params["in"], params["contains"], params["sort"], offset, size, almost)
}
