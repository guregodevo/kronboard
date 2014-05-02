package charts

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"scul"
	"time"
)

type ChartRepository struct {
	Db *scul.DataB
}

type Chart struct {
	Id int64
	Interval int64
	Type 	string
	Created     time.Time
}


const (
	CHART_UPDATE                = "UPDATE CHART SET INTERVAL = $1, TYPE = $2 WHERE ID = $1;"
	CHART_INSERT                = "INSERT INTO CHART (INTERVAL, TYPE, CREATED) values ($1, $2, $3) RETURNING ID"
	CHART_SELECT_ID             = "SELECT INTERVAL, TYPE, CREATED FROM CHART WHERE ID = $1"
)

func (repo *ChartRepository) Create(interval int64, typeString string) (int64, error) {
	created := time.Now()
	var id int64
	error := repo.Db.QueryRow(CHART_INSERT, 3, interval, typeString, created, &id)
	return id, error
}

func (repo *ChartRepository) FindId(id int64) (*Chart, error) {
	var created time.Time
	var interval int64
	var typeString string
	err := repo.Db.QueryRow(CHART_SELECT_ID, 1, id, &interval, &typeString, &created)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &Chart{id, interval, typeString, created}, nil
	}
}


func (chart *Chart) String() string {
	return fmt.Sprintf("Chart [interval='%v',type='%v',created='%v']", chart.Interval, chart.Type, chart.Created)
}
