package main

import (
	"log"
	"core"
	"metrics"
	"github.com/guregodevo/gosequel"
	"redigowrapper"
	"time"
	"fmt"
)

func poll(now time.Time, redisDB *redigowrapper.RedisDB, db *gosequel.DataB) {
	fmt.Printf("%v: Polling ...", now)

	repo := metrics.MetricRepository{db}

	metrics, err := repo.GetAllMetric()
	
	if err!= nil {
		return
	}

	for _, metric := range metrics {
		repo.
		if rerr != nil {
    		log.Printf("Error : %v \n", rerr)
			continue
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

   	poll(time.Now(), &redisDB, &db)

}