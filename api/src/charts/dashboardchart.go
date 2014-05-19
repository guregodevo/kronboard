package charts

import (
	"net/http"
	"net/url"
	"time"
	"github.com/guregodevo/stripack"
	"github.com/guregodevo/pastis"
	"strconv"
	"log"
)


type DashboardChartsResource struct{
	ChartRepository *ChartRepository
	DashboardRepository *DashboardRepository
}


func (resource DashboardChartsResource) Post(values url.Values, chart Chart) (int64, interface{}) {
	id := values.Get("dashboardid")

	if id == "" {
        err := &ChartsError{time.Now(), "Missing dashboard id parameter"}		
		return http.StatusBadRequest, pastis.ErrorResponse(err)
	}

	idInt, errConv := strconv.ParseInt(id,10, 64);
	if  errConv != nil {
        err := &ChartsError{time.Now(), "Wrong dashboard id parameter type"}		
		return http.StatusBadRequest, pastis.ErrorResponse(err)
	}

	//get dashboard
	dashboard, err := resource.DashboardRepository.FindId(idInt)
	if err != nil {
		log.Fatal("Could not get dashboard '%v' \n", id, err)
		return http.StatusInternalServerError, &ChartsError{time.Now(), "Technical error while getting chart record"}
	}
	if  dashboard == nil {
        err := &ChartsError{time.Now(), "Could not create chart"}		
		return http.StatusNotFound, pastis.ErrorResponse(err)
	}

	//Create Chart
	chartid, err := resource.ChartRepository.Create(60, chart.Type, chart.Description)
	if err != nil {
		log.Fatal("Could not get chart '%v' \n", id, err)
		return http.StatusInternalServerError, &ChartsError{time.Now(), "Technical error while getting chart record"}
	}
	chart.Id = chartid
	//Insert chart into dashboard
	rects := ToRects(dashboard.Charts)
	algo := new(stripack.GreedyOnlineAlgo)
	rect := &stripack.Rect{Id:chart.Id, H:1, W:1}
	isPacked, packedRect := algo.Pack(6,5, rects, rect)
	if  !isPacked {
        err := &ChartsError{time.Now(), "Could not pack your chart due to lack of space."}		
		return http.StatusBadRequest, pastis.ErrorResponse(err)
	}
	charts := append(dashboard.Charts, ToJSON(chart, packedRect))
	err = resource.DashboardRepository.Update(idInt, 10, 5, charts)
	if err!=nil {
        e := &ChartsError{time.Now(), "Could not save packed chart."}		
		return http.StatusInternalServerError, pastis.ErrorResponse(e)
	}
	return http.StatusOK, chart
}
