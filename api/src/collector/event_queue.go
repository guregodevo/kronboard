package collector

import (
	"github.com/iwanbk/gobeanstalk"
	//"code.google.com/p/goprotobuf/proto"	
	"log"
)


type EventQueue struct {
    conn *gobeanstalk.Conn
}

func (q *EventQueue) Connect(url string) (err error) {
    conn, err := gobeanstalk.Dial(url)
    if err != nil {
        log.Printf("connect failed")
    } else {
        q.conn = conn
    }
    return err
}

func (q *EventQueue) Write(p []byte) (id uint64, err error) {
    id, err = q.conn.Put(p, 0, 0, 10)
    return id, err
}

func (q *EventQueue) Read() (p []byte, err error) {
    j, err := q.conn.Reserve()
    if err != nil {
        log.Println("reserve failed")
        return nil, err
    }
    //log.Printf("id:%d \n", j.Id)
    err = q.conn.Delete(j.Id)
    if err != nil {
        return nil, err
    }
    return j.Body, nil
}