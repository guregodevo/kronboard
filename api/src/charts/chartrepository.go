package charts

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/guregodevo/gosequel"
	"time"
)

type ChartRepository struct {
	Db *gosequel.DataB
}

type Chart struct {
	Id int64
	Interval int64
	Type 	string
	Description string	
	Created     time.Time
}


const (
	CHART_UPDATE                = "UPDATE CHART SET INTERVAL = $1, TYPE = $2 WHERE ID = $2;"
	CHART_INSERT                = "INSERT INTO CHART (INTERVAL, TYPE, DESCRIPTION, CREATED) values ($1, $2, $3, $4) RETURNING ID"
	CHART_SELECT_ID             = "SELECT INTERVAL, TYPE, DESCRIPTION, CREATED FROM CHART WHERE ID = $1"
)

func (repo *ChartRepository) Create(interval int64, typeString string, description string) (int64, error) {
	created := time.Now()
	var id int64
	error := repo.Db.QueryRow(CHART_INSERT, 4, interval, typeString, description, created, &id)
	return id, error
}

func (repo *ChartRepository) Update(interval int64, typeString string) (int64, error) {
	created := time.Now()
	var id int64
	error := repo.Db.QueryRow(CHART_INSERT, 3, interval, typeString, created, &id)
	return id, error
}

func (repo *ChartRepository) FindId(id int64) (*Chart, error) {
	var created time.Time
	var interval int64
	var typeString string
	var description string
	err := repo.Db.QueryRow(CHART_SELECT_ID, 1, id, &interval, &typeString, &description, &created)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &Chart{id, interval, typeString, description, created}, nil
	}
}

func (chart *Chart) String() string {
	return fmt.Sprintf("Chart [interval='%v',desc=%v, type='%v',created='%v']", chart.Interval, chart.Description, chart.Type, chart.Created)
}
