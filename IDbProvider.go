package hamdahl

type IDbProvider interface {
	Get(tableName string, id string) (*DataSet, error)
	GetAll(tableName string, ids []string) (*DataSet, error)
	Query(tableName string, condition *QueryCondition) (*DataSet, error)

	Add(tableName string, row *DataRow)
	AddRange(tableName string, rows []*DataRow)

	Update(tableName string, row *DataRow)
	UpdateRange(tableName string, rows []*DataRow)

	Remove(tableName string, id string)
	RemoveRange(tableName string, ids []string)
}
