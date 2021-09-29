package model

type Suggest struct {
	Database  string `json:"database"`  //数据库名
	Table     string `json:"table"`     //表名
	Customize string `json:"customize"` //个性化编码
	Keyword   string `json:"keyword"`   //关键字原文
	//SpyNp       string `json:"spyNp"`       //关键字简拼无标点
	SearchCount int  `json:"searchCount"` //搜索次数
	Deleted     bool `json:"deleted"`     //删除状态
}

func NewSuggest(database, table, customize, keyword string) *Suggest {
	suggest := &Suggest{
		Database:    database,
		Table:       table,
		Customize:   customize,
		Keyword:     keyword,
		SearchCount: 1,
		Deleted:     false,
		//SpyNp:       utils.ReplacePunctuation(strings.ToLower(utils.ToSimplePinYin(keyword, false)), ""),
	}
	return suggest
}
