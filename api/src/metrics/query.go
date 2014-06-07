package metrics

import "time"


type QueryDef struct {
	Dimension string
	Filters map[string]interface{}
	Groups []string
}

func NewFullQuery(dimension string, filters map[string]interface{}, groups []string) *QueryDef {
	return &QueryDef{ dimension , filters, groups }
}

func NewQuery(dimension string) *QueryDef {
	return &QueryDef{ dimension , map[string]interface{}{}, []string{} }
}

func (q *QueryDef) Filter(dimKey string, dimValue interface{}) {
	q.Filters[dimKey] = dimValue
}

func (q *QueryDef) GroupBy(dimKey string) {
	q.Groups = append(q.Groups, dimKey)
}

type QueryRange struct {
	start  time.Time
	end	   time.Time		 	
}