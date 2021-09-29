package logic

import (
	"errors"
	"ququ.im/gosearch/elastic"
)

func AddDocument(database, table string, doc map[string]interface{}, searchFields []string) (string, error) {
	//data := make(map[string]interface{})
	//utils.FromJSON(utils.ToJSON(doc), &data)
	//for k, v := range data {
	//	if k == "id" {
	//		continue
	//	}
	//	switch v.(type) {
	//	case string:
	//		if !((len(k) > 3 && k[len(k)-3:] == "Spy") || (len(k) > 5 && k[len(k)-5:] == "SpyNp")) {
	//			data[k+"Spy"] = strings.ToLower(utils.ToSimplePinYin(data[k].(string), false))
	//			data[k+"SpyNp"] = utils.ReplacePunctuation(data[k+"Spy"].(string), "")
	//		}
	//	case []interface{}:
	//		fieldStrSlice := v.([]interface{})
	//		fieldSpy := make([]string, 0)
	//		fieldSpyNp := make([]string, 0)
	//		for _, str := range fieldStrSlice {
	//			spy := strings.ToLower(utils.ToSimplePinYin(str.(string), false))
	//			fieldSpy = append(fieldSpy, spy)
	//			fieldSpyNp = append(fieldSpyNp, utils.ReplacePunctuation(spy, ""))
	//		}
	//		if !((len(k) > 3 && k[len(k)-3:] == "Spy") || (len(k) > 5 && k[len(k)-5:] == "SpyNp")) {
	//			data[k+"Spy"] = fieldSpy
	//			data[k+"SpyNp"] = fieldSpyNp
	//		}
	//	}
	//}
	switch SEARCH_ENGINE {
	case "elasticsearch":
		return elastic.AddDocument(database, table, doc, searchFields)
	}
	return "", nil
}

func AddDocuments(database, table string, docs []map[string]interface{}, searchFields []string) ([]string, error) {
	//for i, data := range docs {
	//	for k, v := range data {
	//		if k == "id" {
	//			continue
	//		}
	//		switch v.(type) {
	//		case string:
	//			if !((len(k) > 3 && k[len(k)-3:] == "Spy") || (len(k) > 5 && k[len(k)-5:] == "SpyNp")) {
	//				data[k+"Spy"] = strings.ToLower(utils.ToSimplePinYin(data[k].(string), false))
	//				data[k+"SpyNp"] = utils.ReplacePunctuation(data[k+"Spy"].(string), "")
	//			}
	//		case []interface{}:
	//			fieldStrSlice := v.([]interface{})
	//			fieldSpy := make([]string, 0)
	//			fieldSpyNp := make([]string, 0)
	//			for _, str := range fieldStrSlice {
	//				spy := strings.ToLower(utils.ToSimplePinYin(str.(string), false))
	//				fieldSpy = append(fieldSpy, spy)
	//				fieldSpyNp = append(fieldSpyNp, utils.ReplacePunctuation(spy, ""))
	//			}
	//			if !((len(k) > 3 && k[len(k)-3:] == "Spy") || (len(k) > 5 && k[len(k)-5:] == "SpyNp")) {
	//				data[k+"Spy"] = fieldSpy
	//				data[k+"SpyNp"] = fieldSpyNp
	//			}
	//		}
	//	}
	//	docs[i] = data
	//}
	switch SEARCH_ENGINE {
	case "elasticsearch":
		return elastic.AddDocuments(database, table, docs, searchFields)
	}
	return nil, nil
}

func DeleteDocument(database, table string, id string) (bool, error) {
	switch SEARCH_ENGINE {
	case "elasticsearch":
		return elastic.DeleteDocument(database, table, id)
	}
	return false, nil
}

func DeleteDocuments(database, table string, ids []string) (bool, error) {
	switch SEARCH_ENGINE {
	case "elasticsearch":
		return elastic.DeleteDocuments(database, table, ids)
	case "searchx":
		return false, errors.New("暂不支持此功能")
	}
	return false, nil
}

func UpdateDocument(database, table, id string, updateData map[string]interface{}) (bool, error) {
	//for k, v := range updateData {
	//	switch v.(type) {
	//	case string:
	//		if !((len(k) > 3 && k[len(k)-3:] == "Spy") || (len(k) > 5 && k[len(k)-5:] == "SpyNp")) {
	//			updateData[k+"Spy"] = strings.ToLower(utils.ToSimplePinYin(updateData[k].(string), false))
	//			updateData[k+"SpyNp"] = utils.ReplacePunctuation(updateData[k+"Spy"].(string), "")
	//		}
	//	case []interface{}:
	//		fieldStrSlice := v.([]interface{})
	//		fieldSpy := make([]string, 0)
	//		fieldSpyNp := make([]string, 0)
	//		for _, str := range fieldStrSlice {
	//			spy := strings.ToLower(utils.ToSimplePinYin(str.(string), false))
	//			fieldSpy = append(fieldSpy, spy)
	//			fieldSpyNp = append(fieldSpyNp, utils.ReplacePunctuation(spy, ""))
	//		}
	//		if !((len(k) > 3 && k[len(k)-3:] == "Spy") || (len(k) > 5 && k[len(k)-5:] == "SpyNp")) {
	//			updateData[k+"Spy"] = fieldSpy
	//			updateData[k+"SpyNp"] = fieldSpyNp
	//		}
	//	}
	//}
	switch SEARCH_ENGINE {
	case "elasticsearch":
		return elastic.UpdateDocument(database, table, id, updateData)
	default:
		return false, errors.New("暂无此功能")
	}
}

func DeleteTable(database, table string) (bool, error) {
	switch SEARCH_ENGINE {
	case "elasticsearch":
		return elastic.DeleteTable(database, table)
	default:
		return false, errors.New("暂无此功能")
	}
}

func DeleteDatabase(database string) (bool, error) {
	switch SEARCH_ENGINE {
	case "elasticsearch":
		return elastic.DeleteDatabase(database)
	default:
		return false, errors.New("暂无此功能")
	}
}
