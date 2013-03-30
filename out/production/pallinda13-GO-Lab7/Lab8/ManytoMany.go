/**
 * Created with IntelliJ IDEA.
 * User: OMMatte
 * Date: 2013-03-21
 * Time: 13:58
 */
// This is a testbed to help you understand channels better.
package Lab8

// Stefan Nilsson 2013-03-13



import (
"fmt"
"math/rand"
"strconv"
"sync"
"time"
)

func RunManytoMany() {
	// Use different random numbers each time this program is executed.
//	rand.Seed(time.Now().Unix())

	const strings = 32
	const producers = 4
	const consumers = 2

	before := time.Now()
	ch := make(chan string)
	wgp := new(sync.WaitGroup)
	wgc := new(sync.WaitGroup)
	// Sets how many go-routines for the WaitGroup to wait for
	wgp.Add(producers)
	wgc.Add(strings)
	for i := 0; i < producers; i++ {
		// Create the producers and send the waitgroup to them. Each producer get a number of strings
		go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
	}
	for i := 0; i < consumers; i++ {
		// Create the consumers and set their name to c0, c1 etc...
		go Consume("c"+strconv.Itoa(i), ch, wgc)
	}
	wgp.Wait() // Wait for all producers to finish.
	//TODO: Not good! Could actually have a datarace here. It is possible that the channel closes before the consumers reads.
	wgc.Wait()
	close(ch)
	fmt.Println("time:", time.Now().Sub(before))
}

// Produce sends n different strings on the channel and notifies wg when done.
func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		RandomSleep(100) // Simulate time to produce data.
		// Write each string together with the name of the producer
		// Will block until read by a consumer
		ch <- id + ":" + strconv.Itoa(i)
	}
	wg.Done()

}

// Consume prints strings received from the channel until the channel is closed.
func Consume(id string, ch <-chan string, wg *sync.WaitGroup) {
	for s := range ch {
		// Reads from the channel and prints the string
		fmt.Println(id, "received", s)
		wg.Done()
		RandomSleep(100) // Simulate time to consume data.

	}

}

// RandomSleep waits for x ms, where x is a random number, 0 â‰¤ x < n,
// and then returns.
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}
/*


* Vad händer om man byter plats på satserna wg.Wait() och close(ch) i slutet av main-funktionen?
Vi kommer stänga kanalen direkt vilket innebär att (troligen alla) läsningarna kommer renturnera null eller tomma strängar.
Troligtvis kommer programmet även krasha då vi försker skriva till en stängd kanal. Test visade att så var fallet.

* Vad händer om man flyttar close(ch) från main-funktionen och i stället stänger kanalen i slutet av funktionen Produce?
Vi kommer stänga kanalen innan vi vet att alla andra producenter är klara. Har vi rikligt med tur så kommer det gå vägen men
troligtvis krashar programmet då någon annan producent som inte är klar ska försöka skriva till kanalen. Ett sista fall som kan hända
är att producenterna hunnit skriva men alla konsumenter är inte klara med att läsa så dom läser in tomma strängar.
Sedan är det säkert inte bra att stänga en stängd channel. Test visade att det blev krash då någon producent försökte skriva till kanalen.

* Vad händer om man tar bort satsen close(ch) helt och hållet?
Vi stänger aldrig kanalen vilket inte är ett jättestort problem eftersom en skräpsamlare kan hantera sådant.
Problemet är dock att konumenterna är skapta på sådant sätt att de aldrig kommer försvinna ur minnet om vi inte stänger kanalen.
Faktum är att det faktiskt är bra i det här fallet då vi har ett datarace och det är troligen bättre att programmet beter sig som det ska än
det äter för mycket minne. Se TODO

* Vad händer om man ökar antalet konsumenter från 2 till 4?
Det kommer vara 4 konsumenter som tar från producenterna istället för 2. Annars ingenting speciellt. Minskar något problemet med datarace som finns

* Kan man vara säker på att alla strängar blir utskrivna innan programmet stannar?
Nej. Har nämnt detta flera gånger och hoppades på att denna fråga inte skulle komma och att det istället var en miss som inte alla upptäcker... =)

*/
