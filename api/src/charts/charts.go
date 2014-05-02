package charts

import (
	"fmt"
	"net/http"
	"net/url"
	"time"	
	"github.com/guregodevo/pastis"
	"strconv"
	"log"
)


type ChartsResource struct{
	Repository *ChartRepository
}

type ChartsError struct {
	When time.Time
	What string
}

func (e *ChartsError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func (resource ChartsResource) GET(values url.Values) (int64, interface{}) {
	id := values.Get("id")

	if id == "" {
        err := &ChartsError{time.Now(), "Missing chart id parameter"}		
		return http.StatusBadRequest, pastis.ErrorResponse(err)
	}

	idInt, errConv := strconv.ParseInt(id,10, 64);
	if  errConv != nil {
        err := &ChartsError{time.Now(), "Wrong chart id parameter type"}		
		return http.StatusBadRequest, pastis.ErrorResponse(err)
	}
	
	chart, err := resource.Repository.FindId(idInt)
	if err != nil {
		log.Fatal("Could not get chart '%v' \n", id, err)
		return http.StatusInternalServerError, &ChartsError{time.Now(), "Technical error while getting chart record"}
	}
	if  chart == nil {
        err := &ChartsError{time.Now(), "Could not find chart"}		
		return http.StatusNotFound, pastis.ErrorResponse(err)
	}	
	return http.StatusOK, chart
}

func (resource ChartsResource) Post(values url.Values) (int64, interface{}) {
	id := values.Get("id")

	if id == "" {
        err := &ChartsError{time.Now(), "Missing chart id parameter"}		
		return http.StatusBadRequest, pastis.ErrorResponse(err)
	}

	idInt, errConv := strconv.ParseInt(id,10, 64);
	if  errConv != nil {
        err := &ChartsError{time.Now(), "Wrong chart id parameter type"}		
		return http.StatusBadRequest, pastis.ErrorResponse(err)
	}
	
	chart, err := resource.Repository.FindId(idInt)
	if err != nil {
		log.Fatal("Could not get chart '%v' \n", id, err)
		return http.StatusInternalServerError, &ChartsError{time.Now(), "Technical error while getting chart record"}
	}
	if  chart == nil {
        err := &ChartsError{time.Now(), "Could not find chart"}		
		return http.StatusNotFound, pastis.ErrorResponse(err)
	}	
	return http.StatusOK, chart
}





