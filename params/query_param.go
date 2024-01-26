package params

import (
	"gorm.io/gorm"
)

type BaseQuery struct {
	PageParam
	SortParam
	GroupParam
	SelectParam
	WhereParam
}

func (query BaseQuery) DoAllParam(db *gorm.DB) {
	query.DoPageAndSort(db)
	query.DoGroup(db)
	query.DoSelectAndWhere(db)
}

func (query BaseQuery) DoPageAndSort(db *gorm.DB) {
	query.DoPage(db)
	query.DoSort(db)
}

func (query BaseQuery) DoGroup(db *gorm.DB) {
	query.DoGroup(db)
}

func (query BaseQuery) DoSelectAndWhere(db *gorm.DB) {
	query.DoSelect(db)
	query.DoWhere(db)
}
