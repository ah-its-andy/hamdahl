package hamdahl

type QueryTable struct {
	Name    string
	Columns *QueryColumnCollection
	Rows    []*QueryRow
}
