/**
 * Created with IntelliJ IDEA.
 * User: OMMatte
 * Date: 2013-03-21
 * Time: 13:05
 */
package Lab8

import "fmt"

// This program should go to 11, but sometimes it only prints 1 to 10.
// Answer: The problem is that we have 2 threads and when the for-loop is finished we do not know if the println or close(ch) gets to go first.
// An easy fix is making another channel that signals when the printing has finished. Sort of like a handshake.

func RunBug2() {
	ch := make(chan int)
	printDone := make(chan bool)
	go Print(ch, printDone)
	for i := 1; i <= 11; i++ {
		ch <- i
		<-printDone
	}

	close(ch)
}


// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, printDone chan<- bool) {
	for n := range ch { // reads from channel until it's closed
		fmt.Println(n)
		printDone <- true
	}
}
