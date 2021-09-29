package service

import (
	"github.com/go-redis/redis"
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/logs"
	config "github.com/maczh/mgconfig"
	"github.com/maczh/utils"
	"ququ.im/gosearch/elastic"
	"ququ.im/gosearch/model"
)

func IncrSuggest(database, table, customize, keyword string, incr int) mgresult.Result {
	if database == "" {
		return *mgresult.Error(-1, "数据库名称不可为空")
	}
	if table == "" {
		return *mgresult.Error(-1, "表名称不可为空")
	}
	if keyword == "" {
		return *mgresult.Error(-1, "关键字不可为空")
	}
	deleted := false
	if customize == "" {
		deleted = config.Redis.ZScore("gosearch:suggest:"+database+table, keyword).Val() < 0
	} else {
		deleted = config.Redis.ZScore("gosearch:suggest:"+database+table+":"+customize, keyword).Val() < 0
	}
	if deleted {
		return *mgresult.Error(-1, "已删除的热搜词")
	}
	id := utils.MD5Encode(database + table + customize + keyword)
	suggest, err := elastic.GetKeywordSuggest(id)
	if err != nil {
		suggest = model.NewSuggest(database, table, customize, keyword)
		err = elastic.AddSuggestion(id, suggest)
		if err != nil {
			logs.Error("新增关键字错误:{}", err.Error())
			return *mgresult.Error(-1, "新增关键字错误:"+err.Error())
		}
	} else {
		suggest.SearchCount += incr
		err = elastic.UpdateSuggestCount(id, suggest.SearchCount)
		if err != nil {
			logs.Error("关键字次数增加错误:{}", err.Error())
			return *mgresult.Error(-1, "关键字次数增加错误:"+err.Error())
		}
	}
	if customize == "" {
		config.Redis.ZIncr("gosearch:suggest:"+database+table, &redis.Z{Score: float64(incr), Member: keyword})
	} else {
		config.Redis.ZIncr("gosearch:suggest:"+database+table+":"+customize, &redis.Z{Score: float64(incr), Member: keyword})
	}
	return *mgresult.Success(suggest)
}

func DeleteSuggest(database, table, customize, keyword string) mgresult.Result {
	if database == "" {
		return *mgresult.Error(-1, "数据库名称不可为空")
	}
	if table == "" {
		return *mgresult.Error(-1, "表名称不可为空")
	}
	if keyword == "" {
		return *mgresult.Error(-1, "关键字不可为空")
	}
	id := utils.MD5Encode(database + table + customize + keyword)
	deleted, err := elastic.DeleteSuggest(id)
	if err != nil {
		logs.Error("删除热搜词错误:{}", err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	if deleted {
		if customize == "" {
			config.Redis.ZAdd("gosearch:suggest:"+database+table, &redis.Z{Score: float64(-100), Member: keyword})
		} else {
			config.Redis.ZAdd("gosearch:suggest:"+database+table+":"+customize, &redis.Z{Score: float64(-100), Member: keyword})
		}
		return *mgresult.Success(nil)
	} else {
		return *mgresult.Error(-1, "无此热搜词或已删除")
	}
}

func ListSearchSuggest(database, table, customize, prefix string, size int) mgresult.Result {
	if database == "" {
		return *mgresult.Error(-1, "数据库名称不可为空")
	}
	if table == "" {
		return *mgresult.Error(-1, "表名称不可为空")
	}
	if prefix == "" { //全部排序，前size个
		key := "gosearch:suggest:" + database + table
		if customize != "" {
			key = "gosearch:suggest:" + database + table + ":" + customize
		}
		result := config.Redis.ZRevRangeByScore(key,
			&redis.ZRangeBy{Count: int64(size),
				Max:    "+inf",
				Min:    "(1",
				Offset: 0,
			}).Val()
		return *mgresult.Success(result)
	} else { //按输入内容前缀获取前size个
		result := make([]string, 0)
		suggests, err := elastic.SearchSuggestions(database, table, customize, prefix)
		if err != nil {
			return *mgresult.Error(-1, err.Error())
		}
		total := len(suggests)
		for i := 0; i < size && i < total; i++ {
			result = append(result, suggests[i].Keyword)
		}
		return *mgresult.Success(result)
	}
}
