package analytics

import (
	//"database/sql"
	"fmt"
	"time"
	_ "github.com/lib/pq"
	"github.com/guregodevo/gosequel"
)

type EventRepository struct {
	Db *gosequel.DataB
}

type Event struct {
	visitId string
	dimensions  map[string]string
	date	time.Time
}

const (
	CREATE_METRICS_TABLE  				= "CREATE TABLE %s_%s_hll (date date, visit_id hll)"
	METRICS_VISIT_INSERT                = "INSERT INTO %s_%s_hll values ($1, hll_add_agg(hll_hash_text($2)) "
	METRICS_VISIT_SELECT_TIMESERIES     = "SELECT %s, hll_cardinality(hll_union_agg(visit_id)) FROM %s_%s_hll GROUP BY 1"
	METRICS_VISIT_SELECT     			= "SELECT %s hll_cardinality(hll_union_agg(visit_id)) FROM %s_%s_hll %s GROUP BY %d"
)

func (repo *EventRepository) Create(clientId int, metricsId int) (error) {
	query := fmt.Sprintf(CREATE_METRICS_TABLE,string(clientId), string(metricsId))
	return repo.Db.QueryRow(query, 2, clientId, metricsId)
}

func (repo *EventRepository) Insert(clientId int, metricsId int, visitId string) (error) {
	insertQuery := fmt.Sprintf(METRICS_VISIT_INSERT,string(clientId), string(metricsId))

	err := repo.Db.QueryRow(insertQuery, 2, time.Now(), visitId)
	return err
}

func (repo *EventRepository) Select(clientId int, metricsId int, query QueryDefinition) (error) {
	timeWindow := ""
	querySQL := fmt.Sprintf(METRICS_VISIT_SELECT, timeWindow, string(clientId), string(metricsId))
	err := repo.Db.QueryRow(querySQL, 1, metricsId)
	return err
}
