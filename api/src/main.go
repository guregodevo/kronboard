package main

import (
	"auth"
	"charts"
	"fmt"
	"log"
	"redigowrapper"
	"github.com/guregodevo/pastis"
	"github.com/guregodevo/gosequel"
)

func main() {
	redisDB := redigowrapper.NewRedisDB("localhost", ":6379", "tcp")
	fmt.Printf("Redis server - %v\n", redisDB.Url())
	sqlDB :=gosequel .DataB{"postgres", "localhost", "postgres", "postgres", "miranalytics", nil}
	fmt.Printf("SQL Database - %v\n", sqlDB.Url())

	nativedb := sqlDB.Opendb()
	defer nativedb.Close()

	authServiceDomain := new(auth.AuthenticationDomain)
	authServiceDomain.TokenRepository = &auth.TokenRepository{&redisDB, 15}
	authServiceDomain.AccountRepository = &auth.AccountRepository{&sqlDB}

	authResource := new(auth.AuthenticationResource)
	authResource.Service = authServiceDomain

	dashboardsResource := new(charts.DashboardsResource)
	dashboardsResource.Repository = &charts.DashboardRepository{&sqlDB}

	chartsResource := new(charts.ChartsResource)
	chartsResource.Repository = &charts.ChartRepository{&sqlDB}

	dashboardChartsResource := new(charts.DashboardChartsResource)
	dashboardChartsResource.ChartRepository = &charts.ChartRepository{&sqlDB}
	dashboardChartsResource.DashboardRepository = &charts.DashboardRepository{&sqlDB}

	var api = pastis.NewAPI()
	api.AddFilter(pastis.LoggingFilter)
	api.AddFilter(pastis.CORSFilter)
	api.AddResource("/authenticate", authResource)
	api.AddResource("/charts/:id", chartsResource)
	api.AddResource("/dashboards/:id", dashboardsResource)
	api.AddResource("/dashboards/:dashboardid/chart/(?P<chartid>[0-9]*)$", dashboardChartsResource)

	log.Printf("Listening on 3000. Go to http://127.0.0.1:3000/")
	err := api.Start(3000)
	if err != nil {
		log.Fatal(err)
	}

}
