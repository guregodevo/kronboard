package main

import (
	"collector"
	"log"
	"github.com/guregodevo/pastis"
)

func main() {

	collectorResource, e := collector.NewCollectorResource("localhost:11300")
	if e != nil {
		log.Fatal(e)
	}

	var api = pastis.NewAPI()
	api.AddResource("/events", collectorResource)

	log.Printf("Listening on 3000. Go to http://127.0.0.1:3000/")
	err := api.Start(4000)
	if err != nil {
		log.Fatal(err)
	}

}
