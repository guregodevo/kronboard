package charts

import "github.com/guregodevo/stripack"

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
	case int64:
	    return t
	case int:
	    return int64(t) 
	}
}

func ToInt(t interface{}) int {
	switch t := t.(type) {
	default:
		return t.(int)
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



