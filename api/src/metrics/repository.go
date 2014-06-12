package metrics

import (
	"fmt"
	"time"
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
	SELECT_FACT_TABLE     				= "SELECT count(*) FROM information_schema.tables WHERE table_name = 'fact_%s'"	
	CREATE_METRICS_TABLE  				= "CREATE TABLE hll_%s_%d (date date, visit_id hll);"
    CREATE_FACT_TABLE  				    = "CREATE TABLE fact_%s (date date, dimensions text);"
    DROP_METRICS_TABLE  				= "DROP TABLE IF EXISTS hll_%s_%d ;"	
	METRICS_EVENT_INSERT                = "INSERT INTO hll_%s_%d(date, visit_id) SELECT $1, hll_add_agg(hll_hash_text($2) )"
	FACT_EVENT_INSERT_GROUP_BY           = "INSERT INTO fact_%s(date, dimensions) values ($1,$2)"
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
	repo.createFactTable(clientId, query)
	err = repo.createMetricTable(id, clientId, query)	
	return err
}

func (repo *MetricRepository) createFactTable(clientId string, query *QueryDef) error {
	var count int
	selectTable := fmt.Sprintf(SELECT_FACT_TABLE, clientId)
	err := repo.Db.QueryRow(selectTable, 0, &count)
	if count == 0 {
		createTable := fmt.Sprintf(CREATE_FACT_TABLE, clientId)
		_, err = repo.Db.Exec(createTable)
	}
	return err
}

func (repo *MetricRepository) createMetricTable(id int64, clientId string, query *QueryDef) error {
	var count int

	selectTable := fmt.Sprintf(SELECT_METRICS_TABLE, clientId, id)
	err := repo.Db.QueryRow(selectTable, 0, &count)
	if count == 0 {
		createTable := fmt.Sprintf(CREATE_METRICS_TABLE, clientId, id)
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

func (repo *MetricRepository) IndexMetric(cm Metric) {
	println("TODO INDEX METRIC")
}


func (repo *MetricRepository) InsertEvent(clientId string, event core.Event) {
	insertQuery := fmt.Sprintf(FACT_EVENT_INSERT_GROUP_BY, clientId)
	hstore := repo.Db.StringMapToHStore(event)
	repo.Db.Exec(insertQuery, time.Now(), hstore)
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
