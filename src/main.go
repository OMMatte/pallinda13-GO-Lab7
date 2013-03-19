/**
 * Created with IntelliJ IDEA.
 * User: OMMatte
 * Date: 2013-03-19
 * Time: 12:42
 */
package main

import (
	"fmt"
	"math"
)


func main() {
	x := 10.0
	delta := 0.0001

	root := Sqrt(x, delta)
	fmt.Print("Approx: ")
	fmt.Println(root)
	fmt.Print("Real: ")
	fmt.Println(math.Sqrt(x))

}

func Sqrt(x float64, delta float64) float64 {
	Z := x/2
	calcDelta := delta + 1.0
	oldValue := calcDelta
	loops := 0
	for (calcDelta > delta){
		Z = ApproxSqrt(x, Z)
		calcDelta = math.Abs(oldValue-Z)
		oldValue = Z
		loops += 1
	}
	fmt.Print("Loops: ")
	fmt.Println(loops)
	return Z
}

func ApproxSqrt(x float64, z float64) float64 {
	r := (z - (z*z-x)/(2.0*z))
	return r
}


