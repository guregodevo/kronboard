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

func poll(now time.Time, redisDB *redigowrapper.RedisDB, db *gosequel.DataB) {
	fmt.Printf("%v: Polling ...", now)

	codec := &core.EventEncodeDecoder{}
	
	repo := metrics.MetricRepository{db}

	metrics, err := repo.GetAllMetric()
	
	if err!= nil {
		return
	}

	for _, metric := range metrics {
		for {
			eventString, rerr := redisDB.String("LPOP", metric.ClientId)
			if rerr != nil {
				break
			}
			//fmt.Printf("Decode %v",eventAstring )
			event, e :=	codec.DecodeBase64(eventString)
			if e == nil && event !=nil {
				repo.InsertEvent(metric, event)
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
	for now := range t.C {
    	poll(now, &redisDB, &db)
	}

}