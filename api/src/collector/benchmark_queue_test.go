package collector

import (
	"testing"
	"net/url"
	"log"
)

func BenchmarkQueueWrites(b *testing.B) {
	c, e := NewCollectorResource("localhost:11300")
	if e != nil {
		log.Fatal(e)
	}
	values, _ := url.ParseQuery("hello.html")
	event := map[string]string {"userId":"12345","timestamp":"1401212227","type":"init_session"}

    for i := 0; i < b.N; i++ {
        c.Post(values, event)
    }
}
