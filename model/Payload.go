package model

type Payload struct {
	Totals int64                    `json:"totals"`
	Time   int64                    `json:"time"`
	Docs   []map[string]interface{} `json:"docs"`
}
