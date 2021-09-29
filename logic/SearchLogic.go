package logic

import (
	"github.com/maczh/logs"
	"ququ.im/gosearch/elastic"
	"ququ.im/gosearch/model"
)

var SEARCH_ENGINE string

func SearchDocument(database, table string, should, must, notMap, shouldWildcard, mustWildcard, notWildcard, rangeQuery map[string]interface{}, in, contains map[string][]interface{}, sortFields []string, offset, size int, outJpy bool) (*model.Payload, error) {
	var payload *model.Payload
	var err error
	switch SEARCH_ENGINE {
	case "elasticsearch":
		payload, err = elastic.SearchDocument(database, table, should, must, notMap, shouldWildcard, mustWildcard, notWildcard, rangeQuery, in, contains, sortFields, offset, size, outJpy)
	}
	if err != nil {
		logs.Error("搜索错误:{}", err.Error())
		return nil, err
	}
	return payload, nil
}
