package main

import (
	"collector"
	"log"
	"fmt"
	"github.com/guregodevo/pastis"
)

func index(out chan collector.Event) {
	for event := range out {
	    fmt.Println("%v to PostgresSQL", event)
    }	
}

func S3(out chan collector.Event) {
	for event := range out {
	    fmt.Println("%v toS3", event)
    }	
}

func broadcast(out chan collector.Event) {
	chanIndex := make(chan collector.Event, 5)
	chanS3 := make(chan collector.Event, 5)

	go index(chanIndex)
	go S3(chanS3)
	for res := range out {
	    chanIndex <- res
		chanS3 <- res	    
    }	
}

func main() {

	collector := collector.NewChannelCollectorResource()

	var api = pastis.NewAPI()
	api.AddResource("/events", collector)
	api.SetLevel("ERROR")
	log.Printf("Listening on 4000. Go to http://127.0.0.1:3000/")
	go collector.Run()	
    go broadcast(collector.Out)

	err := api.Start(4000)
	if err != nil {
		log.Fatal(err)
	}


}
