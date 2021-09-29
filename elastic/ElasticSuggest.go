package elastic

import (
	"context"
	"errors"
	"github.com/maczh/logs"
	config "github.com/maczh/mgconfig"
	"github.com/maczh/utils"
	"github.com/olivere/elastic"
	"ququ.im/gosearch/model"
)

const SuggestIndex = "suggestion"

func AddSuggestion(id string, suggest *model.Suggest) error {
	resp, err := config.Elastic.Index().Index(SuggestIndex).Id(id).BodyJson(*suggest).Do(context.TODO())
	if err != nil {
		logs.Debug("插入文档错误:{}", err.Error())
		return err
	}
	if resp.Result == "created" || resp.Result == "updated" {
		return nil
	} else {
		return errors.New(resp.Result)
	}
}

func GetKeywordSuggest(id string) (*model.Suggest, error) {
	resp, err := config.Elastic.Get().Index(SuggestIndex).Id(id).Do(context.Background())
	if err != nil {
		logs.Error("获取文档错误:{}", err.Error())
		return nil, err
	}
	if resp.Found {
		var suggest model.Suggest
		data, _ := resp.Source.MarshalJSON()
		utils.FromJSON(string(data), &suggest)
		return &suggest, nil
	} else {
		return nil, errors.New("无此记录")
	}
}

func DeleteSuggest(id string) (bool, error) {
	resp, err := config.Elastic.Update().Index(SuggestIndex).Id(id).Doc(map[string]bool{"deleted": true}).Do(context.Background())
	if err != nil {
		logs.Error("删除热搜词错误:{}", err.Error())
		return false, err
	}
	if resp.Result == "updated" {
		return true, nil
	} else {
		return false, errors.New(resp.Result)
	}
}

func UpdateSuggestCount(id string, count int) error {
	searchCount := map[string]interface{}{"searchCount": count}
	resp, err := config.Elastic.Update().Index(SuggestIndex).Id(id).Doc(searchCount).Do(context.TODO())
	if err != nil {
		logs.Error("更新文档错误:{}", err.Error())
		return err
	}
	if resp.Result == "updated" {
		return nil
	} else {
		return errors.New(resp.Result)
	}
}

func SearchSuggestions(database, table, customize, prefix string) ([]model.Suggest, error) {
	searchService := config.Elastic.Search(SuggestIndex)
	boolQuery := elastic.NewBoolQuery().
		Must(elastic.NewTermQuery("database.keyword", database)).
		Must(elastic.NewTermQuery("table.keyword", table)).
		Must(elastic.NewTermQuery("deleted", false))
	if !utils.IsChinese(prefix) {
		boolQuery = boolQuery.Must(elastic.NewPrefixQuery("keyword.spy", prefix))
	} else {
		boolQuery = boolQuery.Must(elastic.NewPrefixQuery("keyword.wildcard", prefix))
	}
	if customize != "" {
		boolQuery = boolQuery.Must(elastic.NewTermQuery("customize.keyword", customize))
	}
	searchService.Query(boolQuery)
	sorter := elastic.NewFieldSort("searchCount").Desc()
	searchService.SortBy(sorter)
	resp, err := searchService.Do(context.TODO())
	logs.Debug("搜索结果:{}", resp)
	if err != nil {
		logs.Error("ElasticSearch查询错误:{}", err.Error())
		return nil, err
	}
	if resp.Status > 0 {
		logs.Error("搜索错误:{}", resp.Error)
		return nil, errors.New(resp.Error.Type + ":" + resp.Error.Reason)
	}
	payload := new(model.Payload)
	payload.Totals = resp.TotalHits()
	suggests := make([]model.Suggest, 0)
	for _, hit := range resp.Hits.Hits {
		jsonBytes, _ := hit.Source.MarshalJSON()
		var doc model.Suggest
		utils.FromJSON(string(jsonBytes), &doc)
		suggests = append(suggests, doc)
	}
	return suggests, nil
}
