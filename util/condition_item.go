package util

import (
	"fmt"
	"log"

	"github.com/spf13/cast"
	"gorm.io/gorm"

	"github.com/zj-kenzhou/web-common/condition"
)

type conditionFunc func(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error

var _conditionFuncMap = make(map[string]conditionFunc)

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

func notEq(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" <> ?", whereItem.Value)
	return nil
}

func eq(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	strValue, isStr := whereItem.Value.(string)
	if isStr && strValue == "" {
		return nil
	}
	validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" = ?", whereItem.Value)
	return nil
}

func eqWithEmpty(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" = ?", whereItem.Value)
	return nil
}

func blurry(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
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
			return nil
		}
	}
	return nil
}

func gt(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" > ?", whereItem.Value)
	return nil
}

func gte(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" >= ?", whereItem.Value)
	return nil
}

func lt(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" < ?", whereItem.Value)
	return nil
}

func lte(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" =< ?", whereItem.Value)
	return nil
}

func between(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	valueList, ok := whereItem.Value.([]any)
	if ok && len(valueList) > 1 {
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" BETWEEN ? AND ?", valueList[0], valueList[1])
	} else {
		log.Println(fmt.Sprintf("属性:%s::的值不合规范::%v", whereItem.Field, whereItem.Value))
	}
	return nil
}

func notBetween(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	valueList, ok := whereItem.Value.([]any)
	if ok && len(valueList) > 1 {
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" NOT BETWEEN ? AND ?", valueList[0], valueList[1])
	} else {
		log.Println(fmt.Sprintf("属性:%s::的值不合规范::%v", whereItem.Field, whereItem.Value))
	}
	return nil
}

func like(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" like ?", "%"+cast.ToString(whereItem.Value)+"%")
	return nil
}

func notLike(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" not like ?", "%"+cast.ToString(whereItem.Value)+"%")
	return nil
}

func leftLike(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" like ?", cast.ToString(whereItem.Value)+"%")
	return nil
}

func rightLike(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" like ?", "%"+cast.ToString(whereItem.Value))
	return nil
}

func isNull(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" IS NULL")
	return nil
}

func isNotNull(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" IS NOT NULL")
	return nil
}

func in(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	valueList, ok := whereItem.Value.([]any)
	if ok && len(valueList) > 0 {
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" IN ?", valueList)
		return nil
	}
	if whereItem.Query != nil {
		childDb, err := buildChildSql(db, *whereItem.Query)
		if err != nil {
			return err
		}
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" IN (?)", childDb)
	} else {
		log.Println(fmt.Sprintf("condition:%s::query and value is nil", whereItem.Condition))
	}
	return nil
}

func notIn(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	valueList, ok := whereItem.Value.([]any)
	if ok && len(valueList) > 0 {
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" NOT IN ?", valueList)
		return nil
	}
	if whereItem.Query != nil {
		childDb, err := buildChildSql(db, *whereItem.Query)
		if err != nil {
			return err
		}
		validAndSetCondition(whereItem.Condition, whereItem.Field, db, isAnd, whereItem.Field+" NOT IN (?)", childDb)
	} else {
		log.Println(fmt.Sprintf("condition:%s::query and value is nil", whereItem.Condition))
	}
	return nil
}

func combo(db *gorm.DB, isAnd bool, whereItem condition.WhereItem) error {
	if len(whereItem.Combo) > 0 {
		childDb := newDb(db)
		err := SetWhere(childDb, whereItem.Combo)
		if err != nil {
			return err
		}
		if isAnd {
			db = db.Where(childDb)
		} else {
			db = db.Or(childDb)
		}
	}
	return nil
}

func init() {
	_conditionFuncMap["notEq"] = notEq
	_conditionFuncMap["eq"] = eq
	_conditionFuncMap["eqWithEmpty"] = eqWithEmpty
	_conditionFuncMap["blurry"] = blurry
	_conditionFuncMap["gt"] = gt
	_conditionFuncMap[">"] = gt
	_conditionFuncMap["gte"] = gte
	_conditionFuncMap[">="] = gte
	_conditionFuncMap["=>"] = gte
	_conditionFuncMap["lt"] = lt
	_conditionFuncMap["<"] = lt
	_conditionFuncMap["lte"] = lte
	_conditionFuncMap["<="] = lte
	_conditionFuncMap["=<"] = lte
	_conditionFuncMap["between"] = between
	_conditionFuncMap["notBetween"] = notBetween
	_conditionFuncMap["like"] = like
	_conditionFuncMap["notLike"] = notLike
	_conditionFuncMap["leftLike"] = leftLike
	_conditionFuncMap["rightLike"] = rightLike
	_conditionFuncMap["isNull"] = isNull
	_conditionFuncMap["isNotNull"] = isNotNull
	_conditionFuncMap["in"] = in
	_conditionFuncMap["notIn"] = notIn
	_conditionFuncMap["combo"] = combo
}
