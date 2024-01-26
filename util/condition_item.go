package util

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/zj-kenzhou/web-common/condition"
	"gorm.io/gorm"
	"log"
	"strings"
)

type conditionFunc func(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error)

var _conditionFuncList []conditionFunc

func oneStrToInt(arg any) any {
	resStr, isStr := arg.(string)
	if isStr && len(resStr) > 15 && IsInt(resStr) {
		return cast.ToInt64(resStr)
	}
	return arg
}

func strArgToInt(arg any) any {
	arg = oneStrToInt(arg)
	resArray, isStrArray := arg.([]any)
	if isStrArray {
		argArrayRes := make([]any, 0)
		for _, item := range resArray {
			argArrayRes = append(argArrayRes, oneStrToInt(item))
		}
		return argArrayRes
	}
	return arg
}

func setCondition(db *gorm.DB, and bool, condition any, args ...any) {
	for argIndex, arg := range args {
		args[argIndex] = strArgToInt(arg)
	}
	if and {
		db = db.Where(condition, args...)
	} else {
		db = db.Or(condition, args...)
	}
}

func validAndSetCondition(conditionName string, field string, db *gorm.DB, and bool, condition any, args ...any) {
	if IsNormalStr(field) {
		setCondition(db, and, condition, args...)
	} else {
		log.Println(fmt.Sprintf("where field is ignore：condition:%s value: %s", conditionName, field))
	}
}

func notEq(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == "notEq" {
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" <> ?", whereItem.Value)
		return true, nil
	}
	return false, nil
}

func eq(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == "eq" {
		strValue, isStr := whereItem.Value.(string)
		if isStr && strValue == "" {
			return true, nil
		}
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" = ?", whereItem.Value)
		return true, nil
	}
	return false, nil
}

func eqWithEmpty(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == "eqWithEmpty" {
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" = ?", whereItem.Value)
		return true, nil
	}
	return false, nil
}

func blurry(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == "blurry" {
		valueStr, ok := whereItem.Value.(string)
		if ok && valueStr != "" {
			if len(whereItem.FieldList) > 0 {
				andDb := newDb(db)
				for index := range whereItem.FieldList {
					field := whereItem.FieldList[index]
					validAndSetCondition(whereItem.Condition, field, andDb, false, field+" like ?", "%"+valueStr+"%")
				}
				if isAnd {
					db = db.Where(andDb)
				} else {
					db = db.Or(andDb)
				}
				return true, nil
			}
		}
	}
	return false, nil
}

func gt(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == ">" {
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" > ?", whereItem.Value)
		return true, nil
	}
	return false, nil
}

func gte(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == ">=" || whereItem.Condition == "=>" {
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" >= ?", whereItem.Value)
		return true, nil
	}
	return false, nil
}

func lt(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == "<" {
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" < ?", whereItem.Value)
		return true, nil
	}
	return false, nil
}

func lte(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == "=<" || whereItem.Condition == "<=" {
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" =< ?", whereItem.Value)
		return true, nil
	}
	return false, nil
}

func between(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == "between" {
		valueList, ok := whereItem.Value.([]any)
		if ok && len(valueList) > 1 {
			validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" BETWEEN ? AND ?", valueList[0], valueList[1])
		} else {
			log.Println(fmt.Sprintf("属性:%s::的值不合规范::%v", whereItem.Field, whereItem.Value))
		}
		return true, nil
	}
	return false, nil
}

func notBetween(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == "notBetween" {
		valueList, ok := whereItem.Value.([]any)
		if ok && len(valueList) > 1 {
			validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" NOT BETWEEN ? AND ?", valueList[0], valueList[1])
		} else {
			log.Println(fmt.Sprintf("属性:%s::的值不合规范::%v", whereItem.Field, whereItem.Value))
		}
		return true, nil
	}
	return false, nil
}

func like(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == "like" {
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" like ?", "%"+cast.ToString(whereItem.Value)+"%")
		return true, nil
	}
	return false, nil
}

func notLike(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == "notLike" {
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" not like ?", "%"+cast.ToString(whereItem.Value)+"%")
		return true, nil
	}
	return false, nil
}

func leftLike(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == "leftLike" {
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" like ?", "%"+cast.ToString(whereItem.Value))
		return true, nil
	}
	return false, nil
}

func rightLike(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == "rightLike" {
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" like ?", cast.ToString(whereItem.Value)+"%")
		return true, nil
	}
	return false, nil
}

func isNull(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == "isNull" {
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" IS NULL")
		return true, nil
	}
	return false, nil
}

func isNotNull(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == "isNotNull" {
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" IS NOT NULL")
		return true, nil
	}
	return false, nil
}

func in(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == "in" {
		valueList, ok := whereItem.Value.([]any)
		if ok && len(valueList) > 0 {
			validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" IN ?", valueList)
			return true, nil
		}
		if whereItem.Query != nil {
			childDb, err := buildChildSql(db, *whereItem.Query)
			if err != nil {
				return false, err
			}
			validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" IN (?)", childDb)
		} else {
			log.Println(fmt.Sprintf("condition:%s::query and value is nil", whereItem.Condition))
		}
		return true, nil
	}
	return false, nil
}

func notIn(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) (bool, error) {
	if whereItem.Condition == "notIn" {
		valueList, ok := whereItem.Value.([]any)
		if ok && len(valueList) > 0 {
			validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" NOT IN ?", valueList)
			return true, nil
		}
		if whereItem.Query != nil {
			childDb, err := buildChildSql(db, *whereItem.Query)
			if err != nil {
				return false, err
			}
			validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" NOT IN (?)", childDb)
		} else {
			log.Println(fmt.Sprintf("condition:%s::query and value is nil", whereItem.Condition))
		}
		return true, nil
	}
	return false, nil
}

func or(db *gorm.DB, _ bool, whereItem condition.WhereItem) (bool, error) {
	if strings.EqualFold(whereItem.Condition, "or") {
		if len(whereItem.Child) > 0 {
			if len(whereItem.Child) == 1 {
				err := putCondition(db, false, whereItem.Child[0])
				if err != nil {
					return false, err
				}
				return true, nil
			}
			childDb := newDb(db)
			err := SetWhere(childDb, whereItem.Child)
			if err != nil {
				return true, err
			}
			db = db.Or(childDb)
		}
		return true, nil
	}
	return false, nil
}

func and(db *gorm.DB, _ bool, whereItem condition.WhereItem) (bool, error) {
	if strings.EqualFold(whereItem.Condition, "and") {
		if len(whereItem.Child) > 0 {
			if len(whereItem.Child) == 1 {
				err := putCondition(db, true, whereItem.Child[0])
				if err != nil {
					return false, err
				}
				return true, nil
			}
			childDb := newDb(db)
			err := SetWhere(childDb, whereItem.Child)
			if err != nil {
				return true, err
			}
			db = db.Where(childDb)
		}
		return true, nil
	}
	return false, nil
}

func init() {
	_conditionFuncList = append(_conditionFuncList, notEq)
	_conditionFuncList = append(_conditionFuncList, eq)
	_conditionFuncList = append(_conditionFuncList, eqWithEmpty)
	_conditionFuncList = append(_conditionFuncList, blurry)
	_conditionFuncList = append(_conditionFuncList, gt)
	_conditionFuncList = append(_conditionFuncList, gte)
	_conditionFuncList = append(_conditionFuncList, lt)
	_conditionFuncList = append(_conditionFuncList, lte)
	_conditionFuncList = append(_conditionFuncList, between)
	_conditionFuncList = append(_conditionFuncList, notBetween)
	_conditionFuncList = append(_conditionFuncList, like)
	_conditionFuncList = append(_conditionFuncList, notLike)
	_conditionFuncList = append(_conditionFuncList, leftLike)
	_conditionFuncList = append(_conditionFuncList, rightLike)
	_conditionFuncList = append(_conditionFuncList, isNull)
	_conditionFuncList = append(_conditionFuncList, isNotNull)
	_conditionFuncList = append(_conditionFuncList, in)
	_conditionFuncList = append(_conditionFuncList, notIn)
	_conditionFuncList = append(_conditionFuncList, or)
	_conditionFuncList = append(_conditionFuncList, and)
}
