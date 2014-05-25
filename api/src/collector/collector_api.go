package collector

import (
	"fmt"
	"net/http"
	"net/url"
	"time"	
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

func (api *CollectorResource) Get(values url.Values, event map[string]string) (int, interface{}) {
	fmt.Printf("\n", event)
	return http.StatusOK, nil
}