package main

import (
	"collector"
	"log"
)

func main() {
	dec := &collector.EventEncodeDecoder{}
	queue := &collector.EventQueue{}

	for {
		if p, err := queue.Read(); err == nil && p != nil {
		   e, event :=	dec.Decode(p)
		   if e == nil {
		   		log.Printf("Recieved : %v \n", event)
		   }
		}
	}
}