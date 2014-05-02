// Repository providing CRUD operations to manage User Account tokens.
package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"redigowrapper"
)

//http://godoc.org/github.com/garyburd/redigo/redis#hdr-Executing_Commands

//A definition of a Token repository
type TokenRepository struct {
	Db                  *redigowrapper.RedisDB
	ExpirationInSeconds int
}

//A token is an expirable hash value of User Account attributes.
type Token struct {
	Id    int64
	Email string
}

//Factory method of Token
func TokenOf(acc *Account) Token {
	return Token{acc.Id, acc.Email}
}

//Encode a token
func (token *Token) Encode() string {
	return fmt.Sprintf("%v:%v", token.Id, token.Email)
}

//Encrypt a token
func (token *Token) Encrypt() string {
	hash := sha256.New()
	data := token.Encode()
	hash.Write([]byte(data))
	hashSum := hash.Sum(nil)
	str := hex.EncodeToString(hashSum)
	return str
}

//Store a token of a User account. Returns the expiration time left of the token and the error if an error occurred.
func (repo *TokenRepository) Put(acc *Account) (string, error) {
	token := TokenOf(acc)
	hash := token.Encrypt()
	repo.Db.ExecRedis("SET", acc.Id, hash)
	_, err := repo.Db.ExecRedis("EXPIRE", acc.Id, repo.ExpirationInSeconds)
	return hash, err
}

//Get a token given a User account. Returns the hash value of the token or an error if an error occcured.
func (repo *TokenRepository) Get(id int64, expire bool) (string, error) {
	value, err := repo.Db.ExecRedis("GET", id)
	if err != nil {
		return "", err
	}
	arr, ok := value.([]uint8)
	if ok {
		hash := string(arr)
		log.Printf("Redis GET key value of '%v' : '%v'", id, hash)
		if expire == true {
			repo.Db.ExecRedis("EXPIRE", id, repo.ExpirationInSeconds)
		}
		return hash, err
	}
	log.Printf("Redis GET No key value of '%v'", id)
	return "", nil
}
