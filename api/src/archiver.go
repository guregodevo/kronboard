package main

import (
//	"log"
	"core"
	"metrics"
	"github.com/guregodevo/gosequel"
	"redigowrapper"
	"time"
	"fmt"
)

func index(now time.Time, db *gosequel.DataB) {
	fmt.Printf("%v: Polling ...", now)

	repo := metrics.MetricRepository{db}

	metrics, err := repo.GetAllMetric()
	
	if err!= nil {
		return
	}

	for _, metric := range metrics {
		repo.IndexMetric(metric)
	}	
}


func ingest(clientIds []string, now time.Time, redisDB *redigowrapper.RedisDB, db *gosequel.DataB) {
	fmt.Printf("%v: Polling ...", now)

	codec := &core.EventEncodeDecoder{}
	
	repo := metrics.MetricRepository{db}

	for _, clientId := range clientIds {
		for {
			eventString, rerr := redisDB.String("LPOP", clientId)
			if rerr != nil {
				break
			}
			event, e :=	codec.DecodeBase64(eventString)
			if e == nil && event !=nil {
				repo.InsertEvent(clientId, event)
			}
		}
	}	
}

func main() {

	redisDB := redigowrapper.NewRedisDB("localhost", ":6379", "tcp")
	fmt.Printf("Redis server - %v\n", redisDB.Url())

    db := gosequel.DataB{"postgres", "localhost", "postgres", "postgres", "miranalytics", nil}
	fmt.Printf("SQL Database - %v\n", db.Url())

	nativedb := db.Opendb()
	defer nativedb.Close()

	t := time.NewTicker(5 * time.Second)
	clients := []string { "123" }
	for now := range t.C {
    	ingest(clients, now, &redisDB, &db)
    	index(now, &db)
	}

}