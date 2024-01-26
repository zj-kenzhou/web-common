package util

import (
	"errors"
	"github.com/zj-kenzhou/web-common/condition"
	"gorm.io/gorm"
)

func newDb(db *gorm.DB) *gorm.DB {
	return db.Session(&gorm.Session{NewDB: true, Initialized: true})
}

func SetWhere(db *gorm.DB, whereList []condition.WhereItem) error {
	if len(whereList) == 0 {
		return nil
	}
	for _, whereItem := range whereList {
		err := putCondition(db, true, whereItem)
		if err != nil {
			return err
		}
	}
	return nil
}

func putCondition(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	for _, conditionFunc := range _conditionFuncList {
		match, err := conditionFunc(db, isAnd, whereItem)
		if err != nil {
			return err
		}
		if match {
			return nil
		}
	}
	return nil
}

func buildChildSql(originDb *gorm.DB, param condition.ChildQuery) (*gorm.DB, error) {
	db := newDb(originDb)
	if param.TableName == "" {
		return nil, errors.New("child sql tableName is empty")
	}
	if !IsNormalStr(param.TableName) {
		return nil, errors.New("child sql tableName is not normal string")
	}
	db.Table(param.TableName)
	if param.Select == "" {
		return nil, errors.New("child sql select is empty")
	}
	if !IsNormalStr(param.Select) {
		return nil, errors.New("child sql select is not normal string")
	}
	db.Select(param.Select)
	err := SetWhere(db, param.Where)
	if err != nil {
		return nil, err
	}
	return db, nil
}
