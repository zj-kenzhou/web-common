package util

import (
	"encoding/json"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/zj-kenzhou/web-common/condition"
)

// 测试条件生成
func TestQuerySql(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(`root:root@tcp(127.0.0.1:3306)/qylcp-backend?charset=utf8mb4&parseTime=True&loc=Local`), &gorm.Config{})
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		tx = tx.Table("SYS_USER")
		queryParam := `[
          { "condition": "notEq", "field": "id", "value": "bbaa"},
          { "condition": "eq", "field": "id", "value": "bbaa"},
          { "condition": "blurry", "fieldList": ["id","bl"], "value": "bbaa"},
          { "condition": ">", "field": "age", "value": 22},
 		  { "condition": ">=", "field": "age", "value": 23},
          { "condition": "<", "field": "age", "value": 40},
          { "condition": "<=", "field": "age", "value": 42},
          { "condition": "between", "field": "create_time", "value": ["23","89"]},
	      { "condition": "notBetween", "field": "create_time", "value": ["23","89"]},
          { "condition": "like", "field": "remarks", "value": "like value"},
          { "condition": "notLike", "field": "remarks", "value": "not like value"},
		  { "condition": "leftLike", "field": "remarks", "value": "not like value"},
		  { "condition": "rightLike", "field": "remarks", "value": "not like value"},
		  { "condition": "isNull", "field": "code"},
          { "condition": "isNotNull", "field": "code"},
          { "condition": "in", "field": "name", "value": ["23","89"]},
          { "condition": "notIn", "field": "name", "value": ["23","89"]},
		  { "condition": "or", "field": "name", "child": [
			{"condition": "eq", "field": "name", "value": "qqq"}
		  ]},
          { "condition": "or", "field": "name", "child": [
			{"condition": "eq", "field": "name", "value": "qqq"},
			{"condition": "eq", "field": "name", "value": "kkk"}
		  ]},
          { "condition": "and", "field": "name", "child": [
			{"condition": "eq", "field": "name", "value": "qqq"}
		  ]},
          { "condition": "and", "field": "name", "child": [
			{"condition": "eq", "field": "name", "value": "qqq"},
			{"condition": "eq", "field": "name", "value": "kkk"}
		  ]}
		]`
		var whereList []condition.WhereItem
		err := json.Unmarshal([]byte(queryParam), &whereList)
		if err != nil {
			t.Error(err)
		}
		err = SetWhere(tx, whereList)
		if err != nil {
			t.Error(err)
		}
		return tx.Scan(map[string]any{})
	})
	t.Log(sql)
}

func TestQuerySql1(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(`root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local`), &gorm.Config{})
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		tx = tx.Table("SYS_USER")
		queryParam := `[
           { "condition": "in", "field": "id", "query": {
				"tableName": "test_table",
				"select": "test_id",
				"where": [
					{ "condition": "eq", "field": "id", "value": "1"}
				]
			   }
           }
		]`
		var whereList []condition.WhereItem
		err := json.Unmarshal([]byte(queryParam), &whereList)
		if err != nil {
			t.Error(err)
		}
		err = SetWhere(tx, whereList)
		if err != nil {
			t.Error(err)
		}
		return tx.Scan(map[string]any{})
	})
	t.Log(sql)
}
