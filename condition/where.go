package condition

type ChildQuery struct {
	Select    string      `json:"select"`
	TableName string      `json:"tableName"`
	Where     []WhereItem `json:"where"`
}
type WhereItem struct {
	Field     string      `json:"field"`
	FieldList []string    `json:"fieldList"`
	Condition string      `json:"condition"`
	Value     any         `json:"value"`
	Query     *ChildQuery `json:"query"`
	Child     []WhereItem `json:"child"`
}
