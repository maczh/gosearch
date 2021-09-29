// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/gosearch/add": {
            "post": {
                "description": "插入文档到全文索引库\u003cbr\u003e说明：1.插入的doc文档只允许一层结构，不要做嵌套\u003cbr\u003e2.允许字符串数组字段参与全文检索\u003cbr\u003e3.所有包含有中文的字段将自动生成一个加Spy后缀的拼音首字母字段，用于支持拼音首字母检索，不输出在搜索结果中\u003cbr\u003e4.数据插入后将自动生成一个名为id的uuid值字段，原有id字段将会被覆盖",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "数据维护"
                ],
                "summary": "插入文档到全文索引库API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据库名称",
                        "name": "database",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "表名称",
                        "name": "table",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "文档数据,JSON类型",
                        "name": "doc",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "模糊搜索字段列表,JSON数组格式",
                        "name": "searchFields",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gosearch/add/batch": {
            "post": {
                "description": "批量插入文档到全文索引库",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "数据维护"
                ],
                "summary": "批量插入文档到全文索引库API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据库名称",
                        "name": "database",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "表名称",
                        "name": "table",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "文档数据数组,JSON数组类型",
                        "name": "docs",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "模糊搜索字段列表,JSON数组格式",
                        "name": "searchFields",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gosearch/del": {
            "post": {
                "description": "从全文索引库中删除文档",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "数据维护"
                ],
                "summary": "从全文索引库中删除文档API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据库名称",
                        "name": "database",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "表名称",
                        "name": "table",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文档id,在插入时返回的id值",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gosearch/del/batch": {
            "post": {
                "description": "从全文索引库中批量删除文档",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "数据维护"
                ],
                "summary": "从全文索引库中批量删除文档API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据库名称",
                        "name": "database",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "表名称",
                        "name": "table",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "文档id列表,JSON数组格式",
                        "name": "ids",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gosearch/del/database": {
            "post": {
                "description": "删除数据库",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "数据维护"
                ],
                "summary": "删除数据库API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据库名称",
                        "name": "database",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gosearch/del/query": {
            "post": {
                "description": "按条件从全文索引库中删除文档",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "数据维护"
                ],
                "summary": "按条件从全文索引库中删除文档API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据库名称",
                        "name": "database",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "表名称",
                        "name": "table",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "查找相等条件,JSON格式，精确查找",
                        "name": "query",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "一个字段符合多个值之一，支持字符串、int、float类型的字段，相当于SQL的in,类似 AND field in ('value1','value2'),JSON格式，{字段名:[值1,值2...]},可多字段",
                        "name": "in",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gosearch/del/table": {
            "post": {
                "description": "从全文索引库中删除表",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "数据维护"
                ],
                "summary": "从全文索引库中删除表API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据库名称",
                        "name": "database",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "表名称",
                        "name": "table",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gosearch/query": {
            "post": {
                "description": "按组合条件从全文索引库中搜索文档\u003cbr\u003e说明:1.输入中文内容的搜索建议放在mustWildcard中，将同时自动生成一个拼音首字母放在mustWildcard中\u003cbr\u003e2.由于没有分词，所以中文将自动拆字，只要包含中文串中的一个字都会命中，因此同时自动加上简拼搜索会大幅度提升命中率\u003cbr\u003e3.多个组合条件精确查找的也都放在must当中，比如shopId=xxxx and userId=xxxx and data like '%xx%'\u003cbr\u003e4.如果需要更精确的结果，建议在搜索结果中再次进行一遍过滤",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "全文检索"
                ],
                "summary": "按组合条件从全文索引库中搜索文档API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据库名称",
                        "name": "database",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "表名称",
                        "name": "table",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OR相等查找条件,JSON格式，{字段名:值},即field1=value1 OR field2=value2,可多字段",
                        "name": "should",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "OR模糊搜索，相当于",
                        "name": "shouldWildcard",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "AND相等查找条件,JSON格式，{字段名:值},即field1=value1 AND field2=value2,可多字段，若某个字段的值为JSON数组，则相当于in查询",
                        "name": "must",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "AND模糊搜索，相当于*关键字*,类似 AND field like '%keyword%',JSON格式，{字段名:关键字},可多字段",
                        "name": "mustWildcard",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "不相等,JSON格式，{字段名:值},类似 AND field \u003c\u003e value,可多字段",
                        "name": "not",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "不包含，相当于*关键字*,类似 AND NOT (field like '%keyword%'),JSON格式，{字段名:关键字},可多字段",
                        "name": "notWildcard",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "一个字段符合多个值之一，支持字符串、int、float类型的字段，相当于SQL的in,类似 AND field in ('value1','value2'),JSON格式，{字段名:[值1,值2...]},可多字段",
                        "name": "in",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "一个字符串数组字段符合多个值之一,仅支持字符串数组类型字段，原始数据集合与查询数据集合有交集则匹配,JSON格式，{字段名:[值1,值2...]},可多字段",
                        "name": "contains",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "按\u003e/\u003e=/\u003c/\u003c=搜索，类似MongoDB，gt/gte/lt/lte,JSON格式，如:{字段名:{",
                        "name": "range",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "搜索结果排序方式，字段名JSON数组，字段名前加+号代表正排序，加-号代表倒排序",
                        "name": "sort",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "分页偏移量，从第几条开始的偏移量，默认为0",
                        "name": "offset",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "分页的单页记录数，默认不分页",
                        "name": "size",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "是否准精确搜索 1-是 0-否， 默认为1",
                        "name": "almost",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gosearch/suggest": {
            "post": {
                "description": "获取下拉补全列表或热搜词",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "热搜补全"
                ],
                "summary": "获取下拉补全列表或热搜词",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据库名称",
                        "name": "database",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "表名称",
                        "name": "table",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "个性化定制标识编码，如用户id、文件id等，若不传则为同一张表共用的热搜词",
                        "name": "customize",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "关键字前缀，若为空则返回热搜列表",
                        "name": "prefix",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "返回记录数，默认10条",
                        "name": "size",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gosearch/suggest/del": {
            "post": {
                "description": "删除关键字",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "热搜补全"
                ],
                "summary": "删除关键字",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据库名称",
                        "name": "database",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "表名称",
                        "name": "table",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "搜索关键字",
                        "name": "keyword",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "个性化定制标识编码，如用户id、文件id等，若不传则为同一张表共用的热搜词",
                        "name": "customize",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gosearch/suggest/incr": {
            "post": {
                "description": "增加搜索关键字一次或新增关键字",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "热搜补全"
                ],
                "summary": "增加搜索关键字一次或新增关键字",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据库名称",
                        "name": "database",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "表名称",
                        "name": "table",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "搜索关键字",
                        "name": "keyword",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "个性化定制标识编码，如用户id、文件id等，若不传则为同一张表共用的热搜词",
                        "name": "customize",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "关键字增量值",
                        "name": "incr",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gosearch/update": {
            "post": {
                "description": "从全文索引库中更新文档",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "数据维护"
                ],
                "summary": "从全文索引库中更新文档API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据库名称",
                        "name": "database",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "表名称",
                        "name": "table",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "文档id,在插入时返回的id值",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "要更新的字段内容,JSON格式",
                        "name": "update",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gosearch/update/query": {
            "post": {
                "description": "按条件从全文索引库中更新文档",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "数据维护"
                ],
                "summary": "按条件从全文索引库中更新文档API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据库名称",
                        "name": "database",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "表名称",
                        "name": "table",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "查找条件,JSON格式，精确查找",
                        "name": "query",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "要更新的字段内容,JSON格式",
                        "name": "update",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0.2(gosearch)",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}