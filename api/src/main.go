package main

import (
	"analytics"
	"auth"
	"charts"
	"fmt"
	"log"
	"redigowrapper"
	"github.com/guregodevo/pastis"
	"scul"
)

func main() {
	redisDB := redigowrapper.NewRedisDB("localhost", ":6379", "tcp")
	fmt.Printf("Redis server - %v\n", redisDB.Url())
	sqlDB := scul.DataB{"postgres", "localhost", "postgres", "postgres", "miranalytics", nil}
	fmt.Printf("SQL Database - %v\n", sqlDB.Url())

	nativedb := sqlDB.Opendb()
	defer nativedb.Close()

	authServiceDomain := new(auth.AuthenticationDomain)
	authServiceDomain.TokenRepository = &auth.TokenRepository{&redisDB, 15}
	authServiceDomain.AccountRepository = &auth.AccountRepository{&sqlDB}

	authResource := new(auth.AuthenticationResource)
	authResource.Service = authServiceDomain

	metricsResource := new(analytics.MetricsResource)

	dashboardsResource := new(charts.DashboardsResource)
	dashboardsResource.Repository = &charts.DashboardRepository{&sqlDB}

	chartsResource := new(charts.ChartsResource)
	chartsResource.Repository = &charts.ChartRepository{&sqlDB}

	var api = pastis.NewAPI()
	api.AddFilter(pastis.LoggingFilter)
	api.AddFilter(pastis.CORSFilter)
	api.AddResource(authResource, "/authenticate")
	api.AddResource(metricsResource, "/metrics")
	api.AddResource(chartsResource, "/charts")
	api.AddResource(dashboardsResource, "/dashboards")

	log.Printf("About to listen on 3000. Go to http://127.0.0.1:3000/")
	err := api.Start(3000)
	if err != nil {
		log.Fatal(err)
	}

}
