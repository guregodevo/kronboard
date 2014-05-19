package charts


import (
	"testing"
	"reflect"
	"github.com/guregodevo/strippacking"
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


func getCharts() []map[string]interface{} {
	c1 := map[string]interface{} {
		"id" : 1,
		"sizeX" : 3,
		"sizeY" : 2,
		"row" : 0,
		"col" : 0,
		"type" : "line",
	}

	c2 := map[string]interface{} {
		"id" : 2,
		"sizeX" : 1,
		"sizeY" : 0,
		"row" : 0,
		"col" : 5,
		"type" : "circle",
	}

	c3 := map[string]interface{} {
		"id" : 1,
		"sizeX" : 2,
		"sizeY" : 0,
		"row" : 3,
		"col" : 3,
		"type" : "circle",
	}

	c4 := map[string]interface{} {
		"id" : 1,
		"sizeX" : 6,
		"sizeY" : 2,
		"row" : 2,
		"col" : 0,
		"type" : "index",
	}

	c5 := map[string]interface{}  {
		"id" : 1,
		"sizeX" : 2,
		"sizeY" : 1,
		"row" : 1,
		"col" : 3,
		"type" : "bar",
	}
	data := []map[string]interface{} {c1,c2, c3, c4, c5}
	return data
}

func Test_Pack(t *testing.T) {
	c1 := map[string]interface{} {
		"id" : 1,
		"sizeX" : 1,
		"sizeY" : 1,
		"row" : 0,
		"col" : 0,
		"type" : "line",
	}

	c2 := map[string]interface{} {
		"id" : 2,
		"sizeX" : 1,
		"sizeY" : 1,
		"row" : 0,
		"col" : 0,
		"type" : "circle",
	}

	c3 := map[string]interface{} {
		"id" : 3,
		"sizeX" : 1,
		"sizeY" : 2,
		"row" : 2,
		"col" : 0,
		"type" : "circle",
	}
	data := []map[string]interface{} {c1,c2, c3}
	rects := ToRects(data)
	expect(t, len(rects), 3)
	algo := new(strippacking.TdAlgo)

	//rects[0].PrintInfo()
	//rects[1].PrintInfo()
	algo.Pack(rects, 0, 0, 10)
	expect(t, len(algo.Rects), 3)
	rectOne := algo.Rects[0]
	expect(t, rectOne.X, float64(1))
	expect(t, rectOne.Y, float64(0))
	expect(t, rectOne.H, float64(1))
	expect(t, rectOne.W, float64(0))
	rectTwo := algo.Rects[1]
	expect(t, rectTwo.X, float64(1))
	expect(t, rectTwo.Y, float64(1))
	expect(t, rectTwo.H, float64(1))
	expect(t, rectTwo.W, float64(0))
}
