package metrics

import (
	"fmt"
	"time"
	"log"
	"core"
	_ "github.com/lib/pq"
	"github.com/guregodevo/gosequel"
)

type MetricRepository struct {
	Db *gosequel.DataB
}

type Event struct {
	visitId string
	dimensions  map[string]string
	date	time.Time
}

const (
	INSERT_METRIC		  				= "INSERT INTO METRICS(CLIENTID, DIMENSION, FILTERS, GROUPS) values ($1, $2, $3, $4 ) RETURNING ID"	
	SELECT_METRIC		  				= "SELECT ID, CLIENTID, DIMENSION, FILTERS, GROUPS FROM METRICS"										   
	HLL_TABLE_NAME						= "hll_%s_%d"
	SELECT_METRICS_TABLE  				= "SELECT count(*) FROM information_schema.tables WHERE table_name = 'hll_%s_%d'"	
	CREATE_METRICS_TABLE  				= "CREATE TABLE hll_%s_%d (date date, visit_id hll);"
	CREATE_METRICS_TABLE_GROUP_BY		= "CREATE TABLE hll_%s_%d (date date, groups text, visit_id hll);"	
    DROP_METRICS_TABLE  				= "DROP TABLE IF EXISTS hll_%s_%d ;"	
	METRICS_EVENT_INSERT                = "INSERT INTO hll_%s_%d(date, visit_id) SELECT $1, hll_add_agg(hll_hash_text($2) )"
	METRICS_EVENT_INSERT_GROUP_BY       = "INSERT INTO hll_%s_%d(date, visit_id, groups) SELECT $1, hll_add_agg(hll_hash_text($2)), $3"
	//METRICS_VISIT_SELECT_TIMESERIES     = "SELECT %s, hll_cardinality(hll_union_agg(visit_id)) FROM hll_%s_%d GROUP BY 1"
	METRICS_VISIT_SELECT     			= "SELECT %s hll_cardinality(hll_union_agg(visit_id)) FROM hll_%s_%d %s"
)

func TableName(m Metric) string {
	return fmt.Sprintf(HLL_TABLE_NAME,m.ClientId, m.Id)	
}

func (repo *MetricRepository) DeleteAllMetrics() (error) {
	metrics, e := repo.GetAllMetric()
	for _, m := range metrics {

		dropTable := fmt.Sprintf(DROP_METRICS_TABLE,m.ClientId, m.Id)
		_, e = repo.Db.Exec(dropTable)
	}
	return e
}


func (repo *MetricRepository) GetAllMetric() ([]Metric, error) {
	metrics := []Metric {}		
	rows, err := repo.Db.Query(SELECT_METRIC)
    if err != nil {
        return []Metric {} , err
    }	
    for rows.Next() {
            var id int64
            var pclientId, pdimension, pfilters, pgroups string
            if err := rows.Scan(&id, &pclientId, &pdimension, &pfilters, &pgroups); err != nil {
                return []Metric {} , err
            }
            query := NewFullQuery(pdimension, repo.Db.StringToHStore(pfilters), repo.Db.StringToArray(pgroups))
            metric := NewMetric(id, pclientId, query)
            metrics = append(metrics, metric)
    }
	return metrics, rows.Err()
}

func (repo *MetricRepository) CreateMetric(clientId string, query *QueryDef) (error) {
	id, err := repo.insertMetric(clientId, query)
	if (err != nil) {
		return err
	}
	err = repo.createMetricTable(id, clientId, query)
	return err
}

func (repo *MetricRepository) createMetricTable(id int64, clientId string, query *QueryDef) error {
	var count int

	selectTable := fmt.Sprintf(SELECT_METRICS_TABLE, clientId, id)
	err := repo.Db.QueryRow(selectTable, 0, &count)
	//fmt.Printf("count : %s \n", count)
	if count == 0 {
		var createTable string
		if query.HasGroup() {
			createTable = fmt.Sprintf(CREATE_METRICS_TABLE_GROUP_BY, clientId, id)
		} else {
			createTable = fmt.Sprintf(CREATE_METRICS_TABLE, clientId, id)	
		}
		_, err = repo.Db.Exec(createTable)
	}
	return err
}

func (repo *MetricRepository) insertMetric(clientId string, query *QueryDef) (int64, error) {
	filters := repo.Db.HStoreToString( query.Filters )
	groups := repo.Db.ArrayToString(query.Groups)
	var id int64
	err := repo.Db.QueryRow(INSERT_METRIC, 4, clientId, query.Dimension, filters, groups, &id)
	return id, err
}

func (repo *MetricRepository) InsertEvent(m Metric, event core.Event) {
	log.Printf("INSERT EVENT: %v \n",event)

	if m.Query.match(event) == true {
		var insertQuery string
		dimensionValue := event[m.Query.Dimension]
		if m.Query.HasGroup() {
			groupByValue := m.Query.GroupByValue(event)
			insertQuery = fmt.Sprintf(METRICS_EVENT_INSERT_GROUP_BY, m.ClientId, m.Id)
			log.Printf("INSERT QUERY: %v \n",insertQuery)
			repo.Db.Exec(insertQuery, time.Now(), dimensionValue, groupByValue)			
		} else {
			insertQuery = fmt.Sprintf(METRICS_EVENT_INSERT, m.ClientId, m.Id)	
			log.Printf("INSERT QUERY: %v \n",insertQuery)
			repo.Db.Exec(insertQuery, time.Now(), dimensionValue)
		}		
	}
}

func (repo *MetricRepository) MetricQuery(m Metric) (map[string]string, error) {
	SQLtimeWindow := ""
	SQLgroupBy := ""
	querySQL := fmt.Sprintf(METRICS_VISIT_SELECT, SQLtimeWindow, m.ClientId, m.Id, SQLgroupBy)
	rows, err := repo.Db.Query(querySQL)
    result := map[string]string {}
    if err != nil {
        return map[string]string {} , err
    }	
    for rows.Next() {
    	if SQLtimeWindow == "" {
    		var r string
            if err := rows.Scan(&r); err != nil {
                //map[string]string {} , err
            }
    	}
    }
	return result , err
}
