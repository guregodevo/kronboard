package charts

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/guregodevo/gosequel"
	"time"
)

type DashboardRepository struct {
	Db *gosequel.DataB
}

type Dashboard struct {
	Id int64
	Name string
	Width int
	Height int
	Charts []map[string]interface{}
	Created	time.Time
}

const (
	DASHBOARD_UPDATE                = "UPDATE DASHBOARD SET CHARTS = $1, WIDTH=$2, HEIGHT=$3 WHERE ID = $4"
	DASHBOARD_INSERT                = "INSERT INTO DASHBOARD (NAME, WIDTH, HEIGHT, CHARTS, CREATED) values ($1, $2, $3, $4, $5) RETURNING ID"
	DASHBOARD_SELECT_ID             = "SELECT NAME, WIDTH, HEIGHT, CHARTS, CREATED FROM DASHBOARD WHERE ID = $1"
)

func (repo *DashboardRepository) Create(name string, width int, height int, charts []map[string]interface{}) (int64, error) {
	created := time.Now()
	var id int64
	hstores := repo.Db.HStoresToString(charts) 
	error := repo.Db.QueryRow(DASHBOARD_INSERT, 5, name, width, height, hstores, created, &id)
	return id, error
}

func (repo *DashboardRepository) Update(id int64, width int, height int, charts []map[string]interface{}) error {
	hstores := repo.Db.HStoresToString(charts)
	//fmt.Printf(hstores)
	error := repo.Db.QueryRow(DASHBOARD_UPDATE, 4, hstores, width, height, id)
	if error == sql.ErrNoRows{
		return nil
	}
	return error
}

func (repo *DashboardRepository) FindId(id int64) (*Dashboard, error) {
	var created time.Time
	var charts string
	var name string
	var width int
	var height int
	err := repo.Db.QueryRow(DASHBOARD_SELECT_ID, 1, id, &name, &width, &height, &charts, &created)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		hstores := repo.Db.StringToHStores(charts)
		return &Dashboard{id, name, width, height, hstores, created}, nil
	}
}

func (charts *Dashboard) String() string {
	return fmt.Sprintf("Dashboard [name='%v',charts='%v',created='%v']", charts.Name, charts.Charts, charts.Created)
}
