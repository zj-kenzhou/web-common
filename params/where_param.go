package params

import (
	"github.com/zj-kenzhou/web-common/condition"
	"github.com/zj-kenzhou/web-common/util"
	"gorm.io/gorm"
	"log"
)

type WhereParam struct {
	Where []condition.WhereItem `json:"where"`
}

func (query WhereParam) DoWhere(db *gorm.DB) {
	err := util.SetWhere(db, query.Where)
	if err != nil {
		log.Println(err)
	}
}
