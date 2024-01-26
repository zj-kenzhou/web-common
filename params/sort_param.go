package params

import (
	"github.com/zj-kenzhou/web-common/util"
	"gorm.io/gorm"
	"strings"
)

type SortParam struct {
	Sort string `query:"sort" json:"sort"`
}

func (query SortParam) DoSort(db *gorm.DB) {
	if query.Sort != "" {
		if strings.Contains(query.Sort, ";") {
			sort := strings.Split(query.Sort, ";")
			for _, item := range sort {
				db = doSortOne(db, item)
			}
		}
		db = doSortOne(db, query.Sort)
	}
}

func doSortOne(db *gorm.DB, sort string) *gorm.DB {
	if strings.Contains(sort, ",") {
		split := strings.Split(sort, ",")
		field := util.CameCaseToUnderscore(split[0])
		if util.IsNormalStr(field) {
			if "descending" == split[1] || "desc" == split[1] {
				return db.Order(field + " desc")
			} else {
				return db.Order(field)
			}
		}
	} else if util.IsNormalStr(sort) {
		return db.Order(util.CameCaseToUnderscore(sort))
	}
	return db
}
