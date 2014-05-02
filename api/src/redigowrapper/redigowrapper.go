//A wrapper of redigo providing higher level instructions to handle connection pools.
package redigowrapper

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

//http://godoc.org/github.com/garyburd/redigo/redis#hdr-Executing_Commands

type RedisDB struct {
	Host      string
	Port      string //":6379"
	Network   string
	RedisPool redis.Pool
}

func NewRedisDB(host string, port string, network string) RedisDB {
	return RedisDB{host, port, network, NewRedisPool(host, network, port)}
}

func NewRedisPool(host string, network string, port string) redis.Pool {
	return redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(network, port)
			if err != nil {
				return nil, err
			}
			//	                 if _, err := c.Do("AUTH", password); err != nil {
			//	                     c.Close()
			//	                     return nil, err
			//	                 }
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

//Redis URL
func (db *RedisDB) Url() string {
	return fmt.Sprintf("%v %v%v\n", db.Network, db.Host, db.Port)
}

//Opens a connection
//func (db *RedisDB) Open() (redis.Conn, error) {
//	conn, err := redis.Dial(db.Network, db.Port)
//	if err != nil {
//    	log.Fatal("Redis connection error",err)
//	}
//	return conn, err
//}

//Executes a Redis command.
func (db *RedisDB) ExecRedis(cmd string, args ...interface{}) (interface{}, error) {
	allArgs := append([]interface{}{cmd}, args...)
	log.Printf("Redis '%v' args=", allArgs...)
	conn := db.RedisPool.Get()
	defer conn.Close()
	n, err := conn.Do(cmd, args...)
	if err != nil {
		log.Fatal("Exec of Redis func failed %v ", cmd, err)
	}
	return n, err
}
