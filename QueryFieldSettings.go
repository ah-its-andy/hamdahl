package hamdahl

type QueryFieldSettings struct {
	Fields []string
}

func NewQueryFieldSettings() *QueryFieldSettings {
	return &QueryFieldSettings{
		Fields: make([]string, 0),
	}
}

func (s *QueryFieldSettings) SetFields(fieldNames []string) *QueryFieldSettings {
	s.Fields = append(s.Fields, fieldNames...)
	return s
}
