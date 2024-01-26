package params

import (
	"gorm.io/gorm"
)

type PageParam struct {
	PageIndex int64 `query:"pageIndex" json:"pageIndex"`
	PageSize  int64 `query:"pageSize" json:"pageSize"`
}

func (query PageParam) DoPage(db *gorm.DB) {
	if query.PageSize > 0 {
		if query.PageIndex > 0 {
			db = db.Offset(int((query.PageIndex - 1) * query.PageSize))
		}
		db = db.Limit(int(query.PageSize))
	}
}
