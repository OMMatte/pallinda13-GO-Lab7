/**
 * Created with IntelliJ IDEA.
 * User: OMMatte
 * Date: 2013-03-19
 * Time: 19:38
 */
package Lab7


import (
	"fmt"
)

// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan<- int) {
	sum := 0
	for _,x := range a{
		sum += x
	}
	res <- sum
}

func RunSumma() {
	a := []int{1, 2, 3, 4, 5, 6, 7}

	n := len(a)
	ch := make(chan int)
	go Add(a[:n/2], ch)
	go Add(a[n/2:], ch)

	<-ch
	sum := <-ch
	fmt.Println(sum)

	// TODO: Get the subtotals from the channel and print their sum.
}
