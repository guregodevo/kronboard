package main

import (
	"collector"
	"redigowrapper"
	"log"
	"fmt"
	"runtime"
	"github.com/guregodevo/pastis"
)

func index(out chan collector.Event, redisDB *redigowrapper.RedisDB, codec *collector.EventEncodeDecoder) {
	for event := range out {
	    fmt.Println("%v to indexToRedis", event)
		data, errEnc := codec.EncodeBase64(event)
		if errEnc != nil {
			fmt.Print("Error encoding base 64")
			continue
		}
		redisDB.ExecRedis("ZADD", event["clientid"], event["timestamp"], data)
    }
}

func S3(out chan collector.Event) {
	for event := range out {
	    fmt.Println("%v toS3", event)
    }	
}

func broadcast(out chan collector.Event, db *redigowrapper.RedisDB, codec *collector.EventEncodeDecoder) {
	chanIndex := make(chan collector.Event, 5)
	chanS3 := make(chan collector.Event, 5)

	go index(chanIndex, db, codec)
	go S3(chanS3)
	for res := range out {
	    chanIndex <- res
		chanS3 <- res	    
    }	
}

func main() {
	runtime.GOMAXPROCS(1)
	c := collector.NewChannelCollectorResource()

	codec := &collector.EventEncodeDecoder{}
    //db := gosequel.DataB{"postgres", "localhost", "postgres", "postgres", "miranalytics", nil}
	//fmt.Printf("SQL Database - %v\n", db.Url())

	redisDB := redigowrapper.NewRedisDB("localhost", ":6379", "tcp")
	fmt.Printf("Redis server - %v\n", redisDB.Url())

	var api = pastis.NewAPI()
	api.AddResource("/events", c)
	api.SetLevel("ERROR")
	log.Printf("Listening on 4000. Go to http://127.0.0.1:4000/")
	go c.Run()	
    go broadcast(c.Out, &redisDB, codec)

	err := api.Start(4000)
	if err != nil {
		log.Fatal(err)
	}

}
