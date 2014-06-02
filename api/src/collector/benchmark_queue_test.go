package collector

import (
	"testing"
	"net/url"
	"net/http"
	//"time"
	"log"
	"bytes"
)

func BenchmarkQueueWrites(b *testing.B) {
	c, e := NewCollectorResource("localhost:11300")
	if e != nil {
		log.Fatal(e)
	}
	values, _ := url.ParseQuery("hello.html")
	event := map[string]string {"userId":"12345","timestamp":"1401212227","type":"init_session"}

    for i := 0; i < b.N; i++ {
        code, _ := c.Post(values, event)
    	if (http.StatusOK != code) {
    		log.Fatal("POST error");
    	}
    }
}

func BenchmarkChannelWrites(b *testing.B) {
	c := NewChannelCollectorResource()
	values, _ := url.ParseQuery("hello.html")
	event := map[string]string {"userId":"12345","timestamp":"1401212227","type":"init_session"}
	go c.Run()
    for i := 0; i < b.N; i++ {
        code, _ := c.Post(values, event)
    	if (http.StatusOK != code) {
    		log.Fatal("POST error");
    	}
    }
    go func() {
    	for res := range c.Out {
    	    log.Println(res)
	    }
    }()

/*	for {
		select {
	    case res := <-c.Out:
	        log.Println(res)
	    case <-time.After(time.Second * 1):
	        log.Println("timeout 1")
	    }
	}*/


}


func BenchmarkChannels2Consumers(b *testing.B) {
	dataSize := 256
	data := make([]byte, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = byte(i % 256)
	}

	ch := make(chan []byte, 128)
	ch2 := make(chan []byte, 128)
	go func() {
		for i := 0; i < b.N; i++ {
			tmp := make([]byte, dataSize)
			copy(tmp, data)
			ch <- tmp
			ch2 <- tmp
		}
	}()

	go func() {
		for i := 0; i < b.N; i++ {
			out := <-ch
			if !bytes.Equal(out, data) {
				b.Fatalf("bytes not the same")
			}
		}		
	}()

	for i := 0; i < b.N; i++ {
		out := <-ch2
		if !bytes.Equal(out, data) {
			b.Fatalf("bytes not the same")
		}
	}
}