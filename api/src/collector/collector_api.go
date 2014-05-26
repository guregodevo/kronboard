package collector

import (
	//"code.google.com/p/goprotobuf/proto"	
	"fmt"
	"net/http"
	"net/url"
	"time"
	"log"
)

type CollectorResource struct {
}

type CollectorError struct {
	When time.Time
	What string
}

func (e *CollectorError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func (api *CollectorResource) Post(values url.Values, event map[string]string) (int, interface{}) {

	//Encode
	// Create an encoder and send a value.
	enc := &EventEncodeDecoder{}
	errEnc, eventAsBytes := enc.Encode(event)
	if errEnc != nil {
		return http.StatusBadRequest, nil
	}

	queue := &EventQueue{}
    id, err := queue.Write(eventAsBytes)

    if err != nil {
        log.Printf("ERROR: ", err)
        return http.StatusInternalServerError, nil
    }
    log.Printf("Job id %d inserted\n", id)	
	return http.StatusOK, nil
}