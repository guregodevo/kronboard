package main

import (
	"github.com/guregodevo/strippacking"
	"math/rand"
	"time"
	"fmt"
	"flag"
	"math/cmplx"
)

func main() {
	prender := flag.Bool("r", true, "Render resulting alignment of all the rectangles")
	strippacking.Prenderbins = flag.Bool("rb", false, "Render bins")
	strippacking.Pnonsolid = flag.Bool("ns", false, "Non solid rendering of rectangles")
	pn := flag.Int("n", 20, "Number of rectangles")
	pm := flag.Int("m", 1, "Number of strips")
	pvalidate := flag.Bool("v", false, "Validate resulting alignment")
	palgo := flag.String("a", "2d", "Type of algorithm")
	ptimes := flag.Int("t", 30, "Number of tests")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	println("Number of rectangles = ", *pn)
	fmt.Printf("N^(2/3) = %0.9v\n\n", real(cmplx.Pow(complex(float64(*pn), 0), (2.0/3))))

	var coef_s float64 = 0
	for y := 0; y < *ptimes; y++ {
		since := time.Now()
		coef := strippacking.Run(*pn, *prender, *pvalidate, *palgo, *pm)
		fmt.Printf("\nTime elapsed  = %v\n", time.Now().Sub(since))	
		coef_s += coef
	}
	fmt.Printf("\nAverage coefficient = %0.9v\n", coef_s/float64(*ptimes))
}
