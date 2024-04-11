package params

import (
	"gorm.io/gorm"

	"github.com/zj-kenzhou/web-common/util"
)

type SortItem struct {
	Field     string `json:"field" query:"field"`
	Direction string `json:"direction" query:"direction"`
}

type SortParam struct {
	Sort []SortItem `query:"sort" json:"sort"`
}

func (query SortParam) DoSort(db *gorm.DB) {
	if len(query.Sort) > 0 {
		for _, item := range query.Sort {
			db = doSortOne(db, item.Field, item.Direction)
		}
	}
}

func doSortOne(db *gorm.DB, field, direction string) *gorm.DB {
	if field != "" && util.IsNormalStr(field) {
		field = util.CameCaseToUnderscore(field)
		if "descending" == direction || "desc" == direction || "descend" == direction {
			return db.Order(field + " desc")
		} else {
			return db.Order(field)
		}
	}
	return db
}
