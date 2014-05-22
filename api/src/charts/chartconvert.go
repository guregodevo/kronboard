package charts

import (
"github.com/guregodevo/stripack"
"strconv"
)

func ToJSON(chart Chart, rect *stripack.Rect) map[string]interface{} {
	json := map[string]interface{} {
		"id" : chart.Id,
		"sizeX" : rect.W,
		"sizeY" : rect.H,
		"row" : rect.Y,
		"col" : rect.X,
		"type" : chart.Type,
	}
	return json
}

func ToInt64(t interface{}) int64 {
	switch t := t.(type) {
	default:
		return t.(int64)
	case string:
		value, err := strconv.ParseInt(t, 10, 64)
	    if err!=nil {
	    	return value
	    } else {
	    	return 0
	    }	    
	case float64:
	    return int64(t)
	case int:
	    return int64(t)
	case int64:
	    return t 
	}
}

func ToInt(t interface{}) int {
	switch t := t.(type) {
	default:
		return t.(int)
	case string:
		value, err := strconv.Atoi(t)
	    if err!=nil {
	    	return value
	    } else {
	    	return 0
	    }	    
	case float64:
	    return int(t)
	case int:
	    return t
	case int64:
	    return int(t) 
	}
}

func ToRect(json map[string]interface{}) *stripack.Rect {
	res := &stripack.Rect{}
	if len(json) > 0 {
		res.X = ToInt(json["col"])
		res.Y = ToInt(json["row"])
		res.H = ToInt(json["sizeY"])
		res.W = ToInt(json["sizeX"])
		res.Id = ToInt64(json["id"])
		return res
	}
	return nil
}


func ToRects(json []map[string]interface{}) []*stripack.Rect {
	res := []*stripack.Rect{}
	for _, rectJson := range json {
		rect := ToRect(rectJson)
		if rect != nil {
			res = append(res, rect)
		}
	}
	return res
}



