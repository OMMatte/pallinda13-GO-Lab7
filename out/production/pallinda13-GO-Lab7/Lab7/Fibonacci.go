/**
 * Created with IntelliJ IDEA.
 * User: OMMatte
 * Date: 2013-03-19
 * Time: 15:08
 */
package Lab7

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	bFib := 1
	bbFib := 0
	return func() int{
		bbFib,bFib = bFib,bbFib+bFib
		return bbFib
	}
}

func RunFibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
