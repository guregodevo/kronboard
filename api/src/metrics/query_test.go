package metrics




import (
	"testing"
	"reflect"
	"core"
)

/* Test Helpers */
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func refute(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}


func TestQueryMatching(t *testing.T)  {
	event := core.Event  {
		"id" : "1",
		"timestamp" : "1402068401",
		"deviceType" : "Samsung",
		"browser" : "firefox",
		"OS" : "Android",
		"type" : "social_action",
	}
	
	query := NewQuery("visitid")
	query.Filter("type","social_action")
	query.GroupBy("browser")

	expect(t, query.match(event),true)
}