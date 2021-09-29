package logic

import (
	"github.com/maczh/utils"
	"strings"
	"unicode/utf8"
)

func HighLightingBySpy(cnString, spyString, keyword string) string {
	l := utf8.RuneCountInString(keyword)
	le := len(keyword)
	highlight := ""
	spyString = strings.ToLower(spyString)
	keyword = strings.ToLower(keyword)
	idx := strings.Index(spyString, keyword)
	for idx >= 0 {
		idxCN := utf8.RuneCountInString(spyString[:idx])
		spyString = spyString[idx+le:]
		highlight = highlight + utils.SubChineseString(cnString, 0, idxCN) + "<em>" + utils.SubChineseString(cnString, idxCN, l) + "</em>"
		cnString = utils.SubChineseString(cnString, idxCN+l, -1)
		idx = strings.Index(spyString, keyword)
	}
	if highlight != "" {
		cnString = highlight + cnString
	}
	return cnString
}
