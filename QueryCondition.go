package hamdahl

type QueryCondition struct {
	LimitSettings *QueryLimitSettings
	FetchFields   *QueryFieldSettings
	SortSettings  *QuerySortSettings
	Groups        []*QueryConditionGroup
}

func NewQueryCondition() *QueryCondition {
	return &QueryCondition{
		LimitSettings: NewQueryLimitSettings(),
		FetchFields:   NewQueryFieldSettings(),
		SortSettings:  NewQuerySortSettings(),
		Groups:        make([]*QueryConditionGroup, 0),
	}
}

func (c *QueryCondition) And() *QueryConditionGroup {
	return NewQueryConditionGroup(QueryRelation_And)
}

func (c *QueryCondition) Or() *QueryConditionGroup {
	return NewQueryConditionGroup(QueryRelation_Or)
}

func (c *QueryCondition) Not() *QueryConditionGroup {
	return NewQueryConditionGroup(QueryRelation_Not)
}
