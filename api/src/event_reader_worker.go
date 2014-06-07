package main

import (
	"collector"
	"log"
)

func main() {
	codec := &collector.EventEncodeDecoder{}
	queue := &collector.EventQueue{}
	e := queue.Connect("localhost:11300")
	if e != nil {
		log.Fatal(e)
	}

	for {
		if p, err := queue.Read(); err == nil && p != nil {
		   event, e :=	codec.Decode(p)
		   if e == nil {
		   		log.Printf("Recieved : %v \n", event)
		   }
		}
	}
}