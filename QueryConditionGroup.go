package hamdahl

type QueryConditionGroup struct {
	QueryRelation int
	Conditions    []*BinaryQueryCondition
}

func NewQueryConditionGroup(queryRelation int) *QueryConditionGroup {
	return &QueryConditionGroup{
		QueryRelation: queryRelation,
		Conditions:    make([]*BinaryQueryCondition, 0),
	}
}

func (c *QueryConditionGroup) Binary() *BinaryQueryCondition {
	binaryCondition := NewBinaryQueryCondition()
	c.Conditions = append(c.Conditions, binaryCondition)
	return binaryCondition
}
