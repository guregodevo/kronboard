package collector

import (
	"github.com/iwanbk/gobeanstalk"
	//"code.google.com/p/goprotobuf/proto"	
	"log"
)


type EventQueue struct {
}


func (q *EventQueue) Write(p []byte) (id uint64, err error) {
	conn, err := gobeanstalk.Dial("localhost:11300")
    if err != nil {
        return 0, err
    }

    id, err = conn.Put(p, 0, 0, 10)
    return id, err
}

func (q *EventQueue) Read() (p []byte, err error) {
	conn, err := gobeanstalk.Dial("localhost:11300")
    if err != nil {
        log.Printf("connect failed")
        return nil, err
    }

    j, err := conn.Reserve()
    if err != nil {
        log.Println("reserve failed")
        return nil, err
    }
    log.Printf("id:%d \n", j.Id)
    err = conn.Delete(j.Id)
    if err != nil {
        return nil, err
    }
    return j.Body, nil
}