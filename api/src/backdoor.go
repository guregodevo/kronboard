// These examples demonstrate more intricate uses of the flag package.
package main

import (
	"auth"
	"metrics"
	"charts"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"redigowrapper"
	"github.com/guregodevo/gosequel"
	"time"
)

var dbAction string

func init() {
	flag.StringVar(&dbAction, "db", "l", "Backdoor commands to operate on the application storage components")
}

func main() {
	redisDB := redigowrapper.NewRedisDB("localhost", ":6379", "tcp")
	fmt.Printf("Redis server - %v\n", redisDB.Url())

    db := gosequel.DataB{"postgres", "localhost", "postgres", "postgres", "miranalytics", nil}
	fmt.Printf("SQL Database - %v\n", db.Url())
	flag.Parse()

	fmt.Printf("[Action=%v]\n", dbAction)
	switch dbAction {
	case "h":
		fmt.Printf("Help to get started...\n")
		fmt.Printf("How to connect to %v? \n", db)
		fmt.Printf(">sudo -u postgres \n")
		fmt.Printf(">\\connect %v \n\n", db.Databasename)
		fmt.Printf("How to create database to %v? \n", db.Databasename)
		fmt.Printf("sudo -u %v psql \n", db.User)
		fmt.Printf("%v=#create database %v \n", db.User, db.Databasename)
		fmt.Printf("How to update '%v' password? \n", db.User)
		fmt.Printf("ALTER USER %v WITH PASSWORD '<new Password>'; \n", db.User)

		//	fmt.Printf("Tips: Password = postgres / sudo -u postgres createdb -h localhost -U postgres miranalytics")
	case "l":
		fmt.Printf("listing the databases table available...\n")
		// "sudo -u postgres psql -l"
		nativedb := db.Opendb()
		defer nativedb.Close()

		rows, _ := db.Query("SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname = 'public'")
		for rows.Next() {
			var tablename string
			err := rows.Scan(&tablename)
			if err == nil {
				fmt.Printf("** %v\n", tablename)
			}
		}
	case "c":
		fmt.Printf("creating tables... \n")
		sqldb := db.Opendb()
		defer sqldb.Close()
		fmt.Printf("creating USER_ACCOUNT... \n")
		db.Exec("CREATE TABLE USER_ACCOUNT(ID SERIAL PRIMARY KEY NOT NULL, EMAIL VARCHAR(50) NOT NULL, COMPANY_NAME VARCHAR(50) NOT NULL, PASSWORD VARCHAR(15) NOT NULL, CREATED TIMESTAMP NOT NULL);")
		fmt.Printf("creating DASHBOARD_ACCOUNT... \n")
		db.Exec("CREATE TABLE DASHBOARD(ID SERIAL, WIDTH SMALLINT, HEIGHT SMALLINT, NAME VARCHAR(100), CHARTS hstore[], CREATED TIMESTAMP NOT NULL );")
		fmt.Printf("creating CHART_ACCOUNT... \n")
		db.Exec("CREATE TABLE CHART(ID SERIAL, INTERVAL SMALLINT, TYPE VARCHAR(100), DESCRIPTION VARCHAR(100), CREATED TIMESTAMP NOT NULL );")
		fmt.Printf("creating METRIC TABLE... \n")
		db.Exec("CREATE TABLE METRICS(ID SERIAL, CLIENTID VARCHAR(100), DIMENSION VARCHAR(100), FILTERS HSTORE, GROUPS TEXT[]);")	

	case "d":
		fmt.Printf("dropping tables... \n")
		nativedb := db.Opendb()
		defer nativedb.Close()
		//db.Exec("DROP TABLE METRICS;")
		db.Exec("DROP TABLE USER_ACCOUNT;")
		db.Exec("DROP TABLE DASHBOARD;")
		db.Exec("DROP TABLE CHART;")
		repoMetrics := metrics.MetricRepository{&db}
		log.Printf("Drop all Metrics ")
		repoMetrics.DeleteAllMetrics()
		db.Exec("DROP TABLE METRICS;")
		
	case "p":
		fmt.Printf("provisioning tables... \n")
		sqldb := db.Opendb()
		defer sqldb.Close()
		repo := &auth.AccountRepository{&db}
		acc, err := repo.Create("jeff.bezos@gmail.com", "companyName", "password")
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Created %v", acc)
		repoDashboard := &charts.DashboardRepository{&db}
		c1 := map[string]interface{} {
			"id" : 1,
			"sizeX" : 3,
			"sizeY" : 2,
			"row" : 0,
			"col" : 0,
			"type" : "line",
		}

		c2 := map[string]interface{} {
			"id" : 2,
			"sizeX" : 1,
			"sizeY" : 0,
			"row" : 0,
			"col" : 3,
			"type" : "circle",
		}

		c3 := map[string]interface{} {
			"id" : 3,
			"sizeX" : 2,
			"sizeY" : 0,
			"row" : 3,
			"col" : 3,
			"type" : "circle",
		}

		c4 := map[string]interface{} {
			"id" : 4,
			"sizeX" : 5,
			"sizeY" : 2,
			"row" : 2,
			"col" : 0,
			"type" : "index",
		}

		c5 := map[string]interface{}  {
			"id" : 5,
			"sizeX" : 2,
			"sizeY" : 1,
			"row" : 1,
			"col" : 3,
			"type" : "bar",
		}
		data := []map[string]interface{} {c1,c2, c3, c4, c5}
		id, errD := repoDashboard.Create("MyDashboard", 6, 5, data)
		if errD != nil {
			log.Fatal(errD)
		}
		log.Printf("Created Dashboard [id=%v]", id)

		repoChart := &charts.ChartRepository{&db}
		repoChart.Create(2,"line","My Line chart")
		repoChart.Create(1,"bar", "My Bar chart")

		query := metrics.NewQuery("visitid")
		query.Filter("type","social_action")
		query.GroupBy("browser")

		repoMetrics := metrics.MetricRepository{&db}
		log.Printf("Created Metric [query=%v]", query)
		repoMetrics.CreateMetric("123", query)

	case "q":
		fmt.Printf("running a few queries... \n")
		sqldb := db.Opendb()
		defer sqldb.Close()
		repo := &auth.AccountRepository{&db}
		acc, err := repo.FindEmail("jeff.bezos@gmail.com")
		if err == nil {
			log.Printf(acc.Tostring())
		}
		acc, err = repo.FindEmailAndPassword("jeff.bezos@gmail.com", "password")
		if err == nil {
			log.Printf(acc.Tostring())
		}
		repoDashboard := &charts.DashboardRepository{&db}
		dashboard, errD := repoDashboard.FindId(1)
		if errD == nil {
			log.Printf(dashboard.String())
		}
		repoChart := &charts.ChartRepository{&db}
		chartOne, errD := repoChart.FindId(1)
		if errD == nil {
			log.Printf(chartOne.String())
		}
	case "r":
		acc := new(auth.Account)
		acc.Id = 1
		fmt.Printf("testing Redis token... \n")
		repoToken := &auth.TokenRepository{&redisDB, 15}
		exp, err := repoToken.Put(acc)
		if err != nil {
			log.Fatal("Could not Put account token %v ", acc, err)
		}
		log.Printf("Put token: expiration = %v  ", exp)
		token, _ := repoToken.Get(acc.Id, true)
		if token != "" {
			log.Printf("Got token: %v  ", token)
		}
		timer := time.NewTimer(11 * time.Second)
		<-timer.C
		log.Printf("expired")
		repoToken.Get(acc.Id, true)
	}
}
