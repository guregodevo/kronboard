package analytics

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type MetricsResource struct{}

type Metrics struct {
	When int64
	key  string
}

func (Metrics) GetMetrics() string {
	//Turn profile into a token
	return "fake_token"
}

type AnalyticsError struct {
	When time.Time
	What string
}

func (e *AnalyticsError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func GetMetrics(t int64, m string) (Metrics, error) {
	if m == "" {
		return Metrics{12233333, ""}, &AnalyticsError{time.Now(), "Wrong user or password"}
	}
	return Metrics{0, "dau"}, nil
}

func (api MetricsResource) Get(values url.Values) (int, interface{}) {
	//timeWhen := values.Get("time")
	//timestamp, _ := strconv.ParseUint(timeWhen, 10, 64)

	metricType := values.Get("type")

	metrics, err := GetMetrics(1000, metricType)

	if err == nil {
		data := map[string]string{"metrics": metrics.key}
		return http.StatusOK, data
	}
	data := map[string]string{"error": "Metrics failure. Metrics not found"}
	return http.StatusNotFound, data
}
