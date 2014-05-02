package charts

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"scul"
	"time"
)

type DashboardRepository struct {
	Db *scul.DataB
}

type Dashboard struct {
	Id int64
	Name string
	Charts []map[string]interface{}
	Created     time.Time
}

const (
	DASHBOARD_UPDATE                = "UPDATE DASHBOARD SET CHARTS = $1 WHERE ID = $1;"
	DASHBOARD_INSERT                = "INSERT INTO DASHBOARD (NAME, CHARTS, CREATED) values ($1, $2, $3) RETURNING ID"
	DASHBOARD_SELECT_ID             = "SELECT NAME, CHARTS, CREATED FROM DASHBOARD WHERE ID = $1"
)

func (repo *DashboardRepository) Create(name string, charts []map[string]interface{}) (int64, error) {
	created := time.Now()
	var id int64
	hstores := repo.Db.HStoresToString(charts) 
	error := repo.Db.QueryRow(DASHBOARD_INSERT, 3, name, hstores, created, &id)
	return id, error
}

func (repo *DashboardRepository) Update(id int64, charts []map[string]interface{}) error {
	created := time.Now()
	hstores := repo.Db.HStoresToString(charts)
	fmt.Printf(hstores)
	error := repo.Db.QueryRow(DASHBOARD_UPDATE, 2, hstores, created)
	return error
}

func (repo *DashboardRepository) FindId(id int64) (*Dashboard, error) {
	var created time.Time
	var charts string
	var name string
	err := repo.Db.QueryRow(DASHBOARD_SELECT_ID, 1, id, &name, &charts, &created)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		hstores := repo.Db.StringToHStores(charts)
		return &Dashboard{id, name, hstores, created}, nil
	}
}

func (charts *Dashboard) String() string {
	return fmt.Sprintf("Dashboard [name='%v',charts='%v',created='%v']", charts.Name, charts.Charts, charts.Created)
}
