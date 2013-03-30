/**
 * Created with IntelliJ IDEA.
 * User: OMMatte
 * Date: 2013-03-21
 * Time: 12:00
 */
package Lab8

import (
	"fmt"
)

// I want this program to print "Hello world!", but it doesn't work.
// We are only using one thread and this causes the ch <- "Hello world!" to block and since we have no other thread able to read from the channel, we are stuck in a deadlock.
// There are 2 ways of solving this. Either make a go routine so that we have the sender and receiver in different threads.
// Or we can make the channel buffered with a big enough buffer, like i have.

func RunBug1() {
	ch := make(chan string, 1)
	ch <- "Hello world!"
	fmt.Println(<-ch)
}


