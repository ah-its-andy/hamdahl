package hamdahl

type BinaryQueryCondition struct {
	FieldFullName  string
	Values         []string
	BinaryOperator int
}

func NewBinaryQueryCondition() *BinaryQueryCondition {
	return &BinaryQueryCondition{}
}

func (c *BinaryQueryCondition) Operator(binaryOperator int) *BinaryQueryCondition {
	c.BinaryOperator = binaryOperator
	return c
}

func (c *BinaryQueryCondition) Field(fieldName string) *BinaryQueryCondition {
	c.FieldFullName = fieldName
	return c
}

func (c *BinaryQueryCondition) Constants(values []string) *BinaryQueryCondition {
	c.Values = values
	return c
}

func (c *BinaryQueryCondition) Constant(value string) *BinaryQueryCondition {
	c.Values = []string{value}
	return c
}
