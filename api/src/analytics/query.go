package analytics

import "time"

type PersistentQuery struct {
	Id string
	ClientId string	
}

type QueryDefinition struct {
	dimension string
	filter *QueryFilter
	group []string
	granularity string
}

type QueryFilter struct {
	filter map[string]string
	timeRange  *QueryRange		 	
}

type QueryRange struct {
	start  time.Time
	end	   time.Time		 	
}

