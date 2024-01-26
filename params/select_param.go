package params

import (
	"github.com/zj-kenzhou/web-common/util"
	"gorm.io/gorm"
	"strings"
)

type SelectParam struct {
	Select []string `query:"select" json:"select"`
}

func (query SelectParam) DoSelect(db *gorm.DB) {
	if len(query.Select) > 0 {
		var selectList []string
		for _, selectStr := range query.Select {
			if util.IsNormalStr(strings.ReplaceAll(selectStr, " ", "")) {
				selectList = append(selectList, selectStr)
			}
		}
		db = db.Select(selectList)
	}
}
