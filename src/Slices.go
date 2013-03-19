/**
 * Created with IntelliJ IDEA.
 * User: OMMatte
 * Date: 2013-03-19
 * Time: 12:42
 */
package main

import (
	"code.google.com/p/go-tour/pic"

)

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)

	for y := range image {
		image[y] = make([]uint8, dx)
		for x := range image[y] {
			image[y][x] = uint8(x * y)
		}
	}
	return image

}

func main() {
	pic.Show(Pic)
}


