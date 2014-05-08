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


type DashboardsResource struct{
	Repository *DashboardRepository
}

type DashboardsError struct {
	When time.Time
	What string
}

func (e *DashboardsError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func (api DashboardsResource) Get(values url.Values) (int, interface{}) {
	id := values.Get("id")

	if id == "" {
        err := &DashboardsError{time.Now(), "Missing dashboard id parameter"}		
		return http.StatusBadRequest, pastis.ErrorResponse(err)
	}

	idInt, errConv := strconv.ParseInt(id,10, 64);
	if  errConv != nil {
        err := &DashboardsError{time.Now(), "Missing dashboard id parameter type"}		
		return http.StatusBadRequest, pastis.ErrorResponse(err)
	}
	
	dashboard, err := api.Repository.FindId(idInt)
	if err != nil {
		log.Fatal("Could not get dashboard '%v' \n", id, err)
		return http.StatusInternalServerError, &DashboardsError{time.Now(), "Technical error while getting dashboard record"}
	}
	if  dashboard == nil {
        err := &DashboardsError{time.Now(), "Could not find dashboard"}		
		return http.StatusNotFound, pastis.ErrorResponse(err)
	}	
	return http.StatusOK, dashboard.Charts
}

func (api DashboardsResource) Put(values url.Values, charts []map[string]interface{}) (int64, interface{}) {
		id := values.Get("id")
		fmt.Printf("Hey Ho %v ",charts)
		if id == "" {
	        err := &DashboardsError{time.Now(), "Missing dashboard id parameter"}		
			return http.StatusBadRequest, pastis.ErrorResponse(err)
		}

		idInt, errConv := strconv.ParseInt(id,10, 64);
		if errConv != nil {
	        err := &DashboardsError{time.Now(), "Missing dashboard id parameter type"}		
			return http.StatusBadRequest, pastis.ErrorResponse(err)
		}
		fmt.Printf("Hey Ho %v %v",idInt, charts)
		err := api.Repository.Update(idInt, charts)
		if err != nil {
			log.Fatal("Could not get dashboard '%v' \n", id, err)
			return http.StatusInternalServerError, &DashboardsError{time.Now(), "Technical error while getting dashboard record"}
		}	
		return http.StatusOK, nil
}

func (api DashboardsResource) Delete(values url.Values) (int, interface{}) {
	data := []map[string]interface{} {};
	return http.StatusOK, data
}