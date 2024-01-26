package params

import (
	"github.com/zj-kenzhou/web-common/util"
	"gorm.io/gorm"
)

type GroupParam struct {
	GroupBy []string `query:"groupBy" json:"groupBy"`
}

func (query GroupParam) DoGroup(db *gorm.DB) {
	for _, group := range query.GroupBy {
		if util.IsNormalStr(group) {
			db = db.Group(group)
		}
	}
}
