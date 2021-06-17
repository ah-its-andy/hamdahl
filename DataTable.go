package hamdahl

type DataTable struct {
	Name    string
	Columns *QueryColumnCollection
	Rows    []*DataRow
}
