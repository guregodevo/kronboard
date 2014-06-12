package metrics

import "core"
import "bytes"

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

func (q *QueryDef) HasGroup() bool {
	return len(q.Groups) > 0
}

func (q *QueryDef) GroupBy(dimKey string) {
	q.Groups = append(q.Groups, dimKey)
}

func (q *QueryDef) GroupByValue(event core.Event) string {
	groups := q.matchingGroups(event)
	var buffer bytes.Buffer
	buffer.WriteString("")
	for i := 0; i < len(groups); i++ {
		buffer.WriteString(groups[i])
		if i < len(groups) - 2  {
			buffer.WriteString("_")
		}
	}
	return buffer.String()
}

func (q *QueryDef) matchingGroups(event core.Event) []string {
	groups := []string{}
	for k, _ := range event {
		for _, kG := range q.Groups {
			if k == kG {
				println(event[k])
				groups = append(groups, event[k])
			}
		}
	}
	return groups
}

func (q *QueryDef) matchGroup(event core.Event) bool {
	if !q.HasGroup() {
		return true
	}
	for k, _ := range event {
		for _, kG := range q.Groups {
			if k == kG {
				return true
			}
		}
	}
	return false
}

func (q *QueryDef) matchFilter(event core.Event) bool {
	for k, v := range event {
		for kF, vF := range q.Filters {
			if k == kF && v != vF {
				return false
			}
		}
	}
	return true
}

func (q *QueryDef) match(event core.Event) bool {
	return q.matchFilter(event) && q.matchGroup(event)
}