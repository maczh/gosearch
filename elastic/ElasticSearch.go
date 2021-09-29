package elastic

import (
	"context"
	"errors"
	"fmt"
	"github.com/maczh/logs"
	config "github.com/maczh/mgconfig"
	"github.com/maczh/utils"
	"github.com/olivere/elastic"
	"ququ.im/gosearch/model"
	"strings"
)

func SearchDocument(database, table string, should, must, notMap, shouldWildcard, mustWildcard, notWildcard, rangeQuery map[string]interface{}, in, contains map[string][]interface{}, sortFields []string, offset, size int, outJpy bool) (*model.Payload, error) {
	indexName := fmt.Sprintf("%s_%s", database, table)
	if table == "" {
		indexName = database
	}
	searchService := config.Elastic.Search(indexName)
	boolQuery := elastic.NewBoolQuery()
	if must != nil && len(must) > 0 {
		for k, v := range must {
			switch v.(type) {
			case string:
				boolQuery = boolQuery.Must(elastic.NewTermQuery(k+".keyword", v))
			case float64, int, int64:
				if k == "id" {
					boolQuery = boolQuery.Must(elastic.NewTermQuery(k+".keyword", fmt.Sprintf("%v", v)))
				} else {
					boolQuery = boolQuery.Must(elastic.NewTermQuery(k, v))
				}
			case []interface{}:
				orQueyy := elastic.NewBoolQuery()
				for _, value := range v.([]interface{}) {
					switch value.(type) {
					case string:
						orQueyy = orQueyy.Should(elastic.NewTermQuery(k+".keyword", value))
					case float64, int, int64:
						orQueyy = orQueyy.Should(elastic.NewTermQuery(k, value))
					}
				}
				boolQuery = boolQuery.Must(orQueyy)
			}
		}
	}
	if mustWildcard != nil && len(mustWildcard) > 0 {
		for k, v := range mustWildcard {
			logs.Debug("正在处理mustWildcard,{}==>{}", k, v)
			switch v.(type) {
			case []string:
				keys := v.([]string)
				for _, key := range keys {
					if key == "" || key == " " {
						continue
					}
					if key[:1] == "*" {
						boolQuery = boolQuery.Must(elastic.NewMatchQuery(k+".iks", key[1:]))
					} else {
						if utils.IsChinese(key) {
							boolQuery = boolQuery.Must(elastic.NewWildcardQuery(k+".wildcard", fmt.Sprintf("*%s*", key)))
						} else {
							boolQuery = boolQuery.Must(elastic.NewWildcardQuery(k+"Jpy.wildcard", fmt.Sprintf("*%s*", strings.ToLower(key))))
						}
					}
				}
			case []interface{}:
				keys := v.([]interface{})
				for _, key := range keys {
					if key.(string) == "" || key.(string) == " " {
						continue
					}
					if key.(string)[:1] == "*" {
						boolQuery = boolQuery.Must(elastic.NewMatchQuery(k+".iks", key.(string)[1:]))
					} else {
						if utils.IsChinese(key.(string)) {
							boolQuery = boolQuery.Must(elastic.NewWildcardQuery(k+".wildcard", fmt.Sprintf("*%s*", key)))
						} else {
							boolQuery = boolQuery.Must(elastic.NewWildcardQuery(k+"Jpy.wildcard", fmt.Sprintf("*%s*", strings.ToLower(key.(string)))))
						}
					}
				}
			case string:
				if v.(string)[:1] == "[" {
					var keys []string
					utils.FromJSON(v.(string), &keys)
					for _, key := range keys {
						if key == "" || key == " " {
							continue
						}
						if key[:1] == "*" {
							boolQuery = boolQuery.Must(elastic.NewMatchQuery(k+".iks", key[1:]))
						} else {
							if utils.IsChinese(key) {
								boolQuery = boolQuery.Must(elastic.NewWildcardQuery(k+".wildcard", fmt.Sprintf("*%s*", key)))
							} else {
								boolQuery = boolQuery.Must(elastic.NewWildcardQuery(k+"Jpy.wildcard", fmt.Sprintf("*%s*", strings.ToLower(key))))
							}
						}
					}
				} else {
					if v.(string)[:1] == "*" {
						boolQuery = boolQuery.Must(elastic.NewMatchQuery(k+".iks", v.(string)[1:]))
					} else {
						if utils.IsChinese(v.(string)) {
							boolQuery = boolQuery.Must(elastic.NewWildcardQuery(k+".wildcard", fmt.Sprintf("*%s*", v)))
							//boolQuery = boolQuery.Must(elastic.NewWildcardQuery(k+"SpyNp.keyword", fmt.Sprintf("*%s*", utils.ToSimplePinYin(v.(string), false))))
						} else {
							boolQuery = boolQuery.Must(elastic.NewWildcardQuery(k+"Jpy.wildcard", fmt.Sprintf("*%s*", strings.ToLower(v.(string)))))
						}
					}
				}
			case int, int64, float64:
				return nil, errors.New("数值字段不可用模糊查询")
			}
		}
	}
	if should != nil && len(should) > 0 {
		for k, v := range should {
			switch v.(type) {
			case string:
				boolQuery = boolQuery.Should(elastic.NewTermQuery(k+".keyword", v))
			case float64, int64, int:
				boolQuery = boolQuery.Should(elastic.NewTermQuery(k, v))
			}
		}
	}
	if shouldWildcard != nil && len(shouldWildcard) > 0 {
		for k, v := range shouldWildcard {
			if utils.IsChinese(v.(string)) {
				boolQuery = boolQuery.Should(elastic.NewMatchQuery(k+".iks", v))
			} else {
				boolQuery = boolQuery.Should(elastic.NewWildcardQuery(k+"Jpy.wildcard", fmt.Sprintf("*%s*", v)))
			}
		}
	}
	if notMap != nil && len(notMap) > 0 {
		for k, v := range notMap {
			switch v.(type) {
			case float64, int, int64:
				boolQuery = boolQuery.MustNot(elastic.NewTermQuery(k, v))
			case string:
				boolQuery = boolQuery.MustNot(elastic.NewTermQuery(k+".keyword", v))
			}
		}
	}
	if notWildcard != nil && len(notWildcard) > 0 {
		for k, v := range notWildcard {
			if utils.IsChinese(v.(string)) {
				boolQuery = boolQuery.MustNot(elastic.NewMatchQuery(k+".iks", v))
			} else {
				boolQuery = boolQuery.MustNot(elastic.NewWildcardQuery(k+"Jpy.wildcard", fmt.Sprintf("*%s*", v)))
			}
		}
	}
	if rangeQuery != nil && len(rangeQuery) > 0 {
		for k, v := range rangeQuery {
			ranQuery := elastic.NewRangeQuery(k)
			rq := v.(map[string]interface{})
			for a, b := range rq {
				switch a {
				case "gte":
					ranQuery = ranQuery.Gte(b)
				case "gt":
					ranQuery = ranQuery.Gt(b)
				case "lte":
					ranQuery = ranQuery.Lte(b)
				case "lt":
					ranQuery = ranQuery.Lt(b)
				}
			}
			boolQuery = boolQuery.Must(ranQuery)
		}
	}
	if in != nil && len(in) > 0 {
		for k, v := range in {
			termsQuery := elastic.NewTermsQuery(k, v...)
			boolQuery = boolQuery.Must(termsQuery)
		}
	}
	if contains != nil && len(contains) > 0 {
		for k, v := range contains {
			termsSetQuery := elastic.NewTermsSetQuery(k+".keyword", v...).
				MinimumShouldMatchScript(elastic.NewScript("Math.min(params.num_terms,1)"))
			boolQuery = boolQuery.Must(termsSetQuery)
		}
	}
	query, _ := boolQuery.Source()
	logs.Debug("搜索语句:{}", query)
	searchService.Query(boolQuery)
	if sortFields != nil && len(sortFields) > 0 {
		for _, field := range sortFields {
			var sorter elastic.Sorter
			if field[:1] == "+" {
				sorter = elastic.NewFieldSort(field[1:]).Asc()
			} else if field[:1] == "-" {
				sorter = elastic.NewFieldSort(field[1:]).Desc()
			} else {
				sorter = elastic.NewFieldSort(field).Asc()
			}
			searchService.SortBy(sorter)
		}
	}
	if size > 0 {
		searchService.Size(size)
	}
	if offset > 0 {
		searchService.From(offset)
	}
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
	docs := make([]map[string]interface{}, 0)
	for _, hit := range resp.Hits.Hits {
		jsonBytes, _ := hit.Source.MarshalJSON()
		doc := make(map[string]interface{})
		utils.FromJSON(string(jsonBytes), &doc)
		if !outJpy {
			for k, _ := range doc {
				if len(k) > 3 && utils.Right(k, 3) == "Jpy" {
					delete(doc, k)
				}
			}
		}
		docs = append(docs, doc)
	}
	payload.Docs = docs
	return payload, nil
}
