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

func (resource DashboardChartsResource) validate(id string) (*Dashboard, int64, string) {
	if id == "" {		
		return nil, http.StatusBadRequest, "Missing dashboard id parameter"
	}

	idInt, errConv := strconv.ParseInt(id,10, 64);
	if  errConv != nil {
		return nil, http.StatusBadRequest, "Wrong dashboard id parameter type"
	}

	//get dashboard
	dashboard, err := resource.DashboardRepository.FindId(idInt)
	if err != nil {
		return nil, http.StatusInternalServerError, "Technical error while getting chart record"
	}
	if  dashboard == nil {		
		return nil, http.StatusNotFound, "Could not create chart"
	}

	return dashboard, http.StatusOK, ""
}

func (resource *DashboardChartsResource) Post(values url.Values, chart Chart) (int64, interface{}) {
	id := values.Get("dashboardid")


	dashboard, status, msg := resource.validate(id)
	if msg != "" {
		error := &ChartsError{time.Now(), msg}
		return status, pastis.ErrorResponse(error)	
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
	err = resource.DashboardRepository.Update(dashboard.Id, 10, 5, charts)
	if err!=nil {
        e := &ChartsError{time.Now(), "Could not save packed chart."}		
		return http.StatusInternalServerError, pastis.ErrorResponse(e)
	}
	return http.StatusOK, chart
}

func (resource *DashboardChartsResource) Delete(values url.Values, chart Chart) (int64, interface{}) {
	id := values.Get("dashboardid")
	chartid := values.Get("chartid")

	chartidInt, errConv := strconv.ParseInt(chartid,10, 64);
	if  errConv != nil {
		error := &ChartsError{time.Now(), "Wrong chart id parameter type"}
		return http.StatusBadRequest, pastis.ErrorResponse(error) 
	}

	dashboard, status, msg := resource.validate(id)
	if msg != "" {
		error := &ChartsError{time.Now(), msg}
		return status, pastis.ErrorResponse(error)	
	}

	//Delete chart into dashboard
	charts := []map[string]interface{} {} 
		
	for _, c := range dashboard.Charts {
		if c["id"].(int64) != chartidInt {
			charts = append(charts, c)
		}
	}
	err := resource.DashboardRepository.Update(dashboard.Id, 10, 5, charts)
	if err!=nil {
        e := &ChartsError{time.Now(), "Could not delete chart."}		
		return http.StatusInternalServerError, pastis.ErrorResponse(e)
	}
	return http.StatusOK, charts
}
