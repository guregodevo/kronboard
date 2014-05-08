package charts


type PackCharts []map[string]interface{}

func Weight(charts PackCharts) int {
	w := 0
	for _, chart := range charts {
		sizeX := chart["sizeX"].(int)
		sizeY := chart["sizeY"].(int)
		w = w + sizeX * sizeY   
	}
	return w
}


func MaxPack(X int, Y int, charts []map[string]interface{}, values []map[string]interface{}) []map[string]interface{} {
 	max := make([]map[string]interface{}) //max capacit for X,Y
 	maxcapacity := make([][]int)
 	maxcapacity[0][0] = 0
 	for n:=0; n < X * Y; n++ {
	 	for x := 0; x <= X ; x++ {
			for y := 0; y <= Y ; y++ { 	
										 	
	 		}
	 	}
 	}
 	return (pack);	
}

