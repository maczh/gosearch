package service

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/utils"
	"ququ.im/gosearch/logic"
	"ququ.im/pinyin"
	"strings"
)

func SearchDocument(database, table, should, must, not, shouldWildcard, mustWildcard, notWildcard, rangeQuery, in, contains, sort string, offset, size, almost int) mgresult.Result {
	if database == "" {
		return *mgresult.Error(-1, "数据库名称不可为空")
	}
	//if table == "" {
	//	return *mgresult.Error(-1, "表名称不可为空")
	//}
	shouldMap := make(map[string]interface{})
	mustMap := make(map[string]interface{})
	notMap := make(map[string]interface{})
	shouldWildcardMap := make(map[string]interface{})
	mustWildcardMap := make(map[string]interface{})
	notWildcardMap := make(map[string]interface{})
	rangeQueryMap := make(map[string]interface{})
	inMap := make(map[string][]interface{})
	containsMap := make(map[string][]interface{})
	if should != "" {
		utils.FromJSON(should, &shouldMap)
	}
	if must != "" {
		utils.FromJSON(must, &mustMap)
	}
	if not != "" {
		utils.FromJSON(not, &notMap)
	}
	if shouldWildcard != "" {
		utils.FromJSON(shouldWildcard, &shouldWildcardMap)
	}
	if mustWildcard != "" {
		utils.FromJSON(mustWildcard, &mustWildcardMap)
		for field, keyword := range mustWildcardMap {
			mustWildcardMap[field] = utils.ReplacePunctuation(keyword.(string), "")
			keywords := strings.Split(utils.AddSpaceBetweenCharsAndNumbers(keyword.(string)), " ")
			if almost == 1 { //准精确搜索
				if len(keywords) > 1 {
					keys := []string{strings.ToLower(utils.ReplacePunctuation(pinyin.ToSimplePinYin(keyword.(string), false), ""))}
					for _, k := range keywords {
						if utils.IsChinese(k) {
							keys = append(keys, k)
						}
					}
					mustWildcardMap[field] = keys
				}
			} else {
				if len(keywords) > 1 {
					keys := make([]string, 0)
					for _, k := range keywords {
						if utils.IsChinese(k) {
							keys = append(keys, "*"+k)
						} else {
							keys = append(keys, k)
						}
					}
					mustWildcardMap[field] = keys
				} else {
					if utils.IsChinese(keywords[0]) {
						mustWildcardMap[field] = "*" + keywords[0]
					}
				}
			}
		}
	}
	if notWildcard != "" {
		utils.FromJSON(notWildcard, &notWildcardMap)
	}
	if rangeQuery != "" {
		utils.FromJSON(rangeQuery, &rangeQueryMap)
	}
	if in != "" {
		utils.FromJSON(in, &inMap)
	}
	if contains != "" {
		utils.FromJSON(contains, &containsMap)
	}
	sortFields := make([]string, 0)
	if sort != "" {
		utils.FromJSON(sort, &sortFields)
	}
	payload, err := logic.SearchDocument(database, table, shouldMap, mustMap, notMap, shouldWildcardMap, mustWildcardMap, notWildcardMap, rangeQueryMap, inMap, containsMap, sortFields, offset, size, false)
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	return *mgresult.Success(payload)
}
