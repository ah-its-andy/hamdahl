package hamdahl

type QueryLimitSettings struct {
	Offsets int
	Limits  int
}

func NewQueryLimitSettings() *QueryLimitSettings {
	return &QueryLimitSettings{
		Offsets: -1,
		Limits:  -1,
	}
}
