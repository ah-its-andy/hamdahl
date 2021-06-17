package hamdahl

type QuerySortSettings struct {
	Fields        []string
	SortDirection int
}

func NewQuerySortSettings() *QuerySortSettings {
	return &QuerySortSettings{
		Fields:        make([]string, 0),
		SortDirection: SortDirection_Asc,
	}
}

func (s *QuerySortSettings) SetFields(fieldNames []string) *QuerySortSettings {
	s.Fields = append(s.Fields, fieldNames...)
	return s
}
