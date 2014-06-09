package metrics

type Metric struct {
	Id int64
	ClientId string
	Query *QueryDef
}

func NewMetric(id int64, clientId string, query *QueryDef) Metric {
	return Metric{id, clientId, query}
}
