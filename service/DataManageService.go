package service

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/logs"
	"github.com/maczh/utils"
	"ququ.im/gosearch/logic"
	"ququ.im/pinyin"
	"strings"
)

func AddDocument(database, table, doc, fields string) mgresult.Result {
	if database == "" {
		return *mgresult.Error(-1, "数据库名称不可为空")
	}
	//if table == "" {
	//	return *mgresult.Error(-1, "表名称不可为空")
	//}
	if doc == "" {
		return *mgresult.Error(-1, "记录不可为空")
	}
	data := make(map[string]interface{})
	utils.FromJSON(doc, &data)
	var searchFields []string
	if fields != "" {
		utils.FromJSON(fields, &searchFields)
	}
	id, err := logic.AddDocument(database, table, data, searchFields)
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	result := make(map[string]string)
	result["id"] = id
	return *mgresult.Success(result)
}

func AddDocuments(database, table, docs, fields string) mgresult.Result {
	if database == "" {
		return *mgresult.Error(-1, "数据库名称不可为空")
	}
	//if table == "" {
	//	return *mgresult.Error(-1, "表名称不可为空")
	//}
	if docs == "" {
		return *mgresult.Error(-1, "记录表不可为空")
	}
	if docs[:1] != "[" {
		return *mgresult.Error(-1, "记录表不是JSON数组格式")
	}
	data := make([]map[string]interface{}, 0)
	utils.FromJSON(docs, &data)
	var searchFields []string
	if fields != "" {
		utils.FromJSON(fields, &searchFields)
	}
	ids, err := logic.AddDocuments(database, table, data, searchFields)
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	result := make(map[string]interface{})
	result["ids"] = ids
	return *mgresult.Success(result)
}

func DeleteDocument(database, table, id string) mgresult.Result {
	if database == "" {
		return *mgresult.Error(-1, "数据库名称不可为空")
	}
	//if table == "" {
	//	return *mgresult.Error(-1, "表名称不可为空")
	//}
	if id == "" {
		return *mgresult.Error(-1, "记录id不可为空")
	}
	_, err := logic.DeleteDocument(database, table, id)
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	return *mgresult.Success(nil)
}

func DeleteDocuments(database, table, ids string) mgresult.Result {
	if database == "" {
		return *mgresult.Error(-1, "数据库名称不可为空")
	}
	//if table == "" {
	//	return *mgresult.Error(-1, "表名称不可为空")
	//}
	if ids == "" || ids[:1] != "[" {
		return *mgresult.Error(-1, "记录id列表不可为空或格式不正确")
	}
	docIds := make([]string, 0)
	utils.FromJSON(ids, &docIds)
	_, err := logic.DeleteDocuments(database, table, docIds)
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	return *mgresult.Success(nil)
}

func DeleteDocumentByQuery(database, table, must, in string) mgresult.Result {
	if database == "" {
		return *mgresult.Error(-1, "数据库名称不可为空")
	}
	//if table == "" {
	//	return *mgresult.Error(-1, "表名称不可为空")
	//}
	mustMap := make(map[string]interface{})
	inMap := make(map[string][]interface{})
	if must != "" {
		utils.FromJSON(must, &mustMap)
	}
	if in != "" {
		utils.FromJSON(in, &inMap)
	}
	ids := make([]string, 0)
	payload, err := logic.SearchDocument(database, table, nil, mustMap, nil, nil, nil, nil, nil, inMap, nil, nil, 0, 100, false)
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	i := 0
	for len(payload.Docs) > 0 {
		for _, doc := range payload.Docs {
			ids = append(ids, doc["id"].(string))
		}
		i++
		if len(payload.Docs) == 100 {
			payload, _ = logic.SearchDocument(database, table, nil, mustMap, nil, nil, nil, nil, nil, nil, nil, nil, i*100, 100, false)
		} else {
			break
		}
	}
	if len(ids) > 0 {
		logs.Debug("正在批量删除以下文档:{}", ids)
		_, err = logic.DeleteDocuments(database, table, ids)
		if err != nil {
			return *mgresult.Error(-1, err.Error())
		}
	}
	return *mgresult.Success(nil)
}

func UpdateDocument(database, table, id, updateData string) mgresult.Result {
	if database == "" {
		return *mgresult.Error(-1, "数据库名称不可为空")
	}
	//if table == "" {
	//	return *mgresult.Error(-1, "表名称不可为空")
	//}
	if id == "" {
		return *mgresult.Error(-1, "记录id不可为空")
	}
	ud := make(map[string]interface{})
	utils.FromJSON(updateData, &ud)
	_, err := logic.UpdateDocument(database, table, id, ud)
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	return *mgresult.Success(nil)
}

func UpdateDocumentByQuery(database, table, must, updateData string) mgresult.Result {
	if database == "" {
		return *mgresult.Error(-1, "数据库名称不可为空")
	}
	//if table == "" {
	//	return *mgresult.Error(-1, "表名称不可为空")
	//}
	mustMap := make(map[string]interface{})
	utils.FromJSON(must, &mustMap)
	payload, err := logic.SearchDocument(database, table, nil, mustMap, nil, nil, nil, nil, nil, nil, nil, nil, 0, 0, true)
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	ud := make(map[string]interface{})
	utils.FromJSON(updateData, &ud)
	if payload.Totals > 0 {
		for k, v := range ud {
			if utils.Existi(payload.Docs[0], k+"Jpy") {
				ud[k+"Jpy"] = utils.ReplacePunctuation(strings.ToLower(pinyin.ToSimplePinYin(v.(string), false)), "")
			}
		}
	}
	for i := 0; i < int(payload.Totals); i++ {
		_, err = logic.UpdateDocument(database, table, payload.Docs[i]["id"].(string), ud)
		if err != nil {
			logs.Error("更新文档id={}错误:{}", payload.Docs[0]["id"].(string), err.Error())
		} else {
			logs.Debug("更新文档id={}成功", payload.Docs[0]["id"].(string))
		}
	}
	return *mgresult.Success(nil)
}

func DeleteTable(database, table string) mgresult.Result {
	if database == "" {
		return *mgresult.Error(-1, "数据库名称不可为空")
	}
	//if table == "" {
	//	return *mgresult.Error(-1, "表名称不可为空")
	//}
	_, err := logic.DeleteTable(database, table)
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	return *mgresult.Success(nil)
}

func DeleteDatabase(database string) mgresult.Result {
	if database == "" {
		return *mgresult.Error(-1, "数据库名称不可为空")
	}
	_, err := logic.DeleteDatabase(database)
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	return *mgresult.Success(nil)
}
