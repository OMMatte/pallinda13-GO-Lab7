/**
 * Created with IntelliJ IDEA.
 * User: OMMatte
 * Date: 2013-03-29
 * Time: 14:14
 */
package Lab9

import (
	"fmt"
	"sync"
)

// This programs demonstrates how a channel can be used for sending and
// receiving by any number of goroutines. It also shows how  the select
// statement can be used to choose one out of several communications.
func RunMatching() {
	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
	match := make(chan string, 1) // Make room for one unmatched send.
	wg := new(sync.WaitGroup)
	wg.Add(len(people))
	for _, name := range people {
		 go Seek(name, match, wg)
	}

	wg.Wait()
	select {
	case name := <-match:
		fmt.Printf("No one received %sâ€™s message.\n", name)
//	default:
		// There was no pending send operation.
	}
}

// Seek either sends or receives, whichever possible, a name on the match
// channel and notifies the wait group when done.
func Seek(name string, match chan string, wg *sync.WaitGroup) {
	select {
	case peer := <-match:
		fmt.Printf("%s sent a message to %s.\n", peer, name)
	case match <- name:
		// Wait for someone to receive my message.
	}
	wg.Done()
}
/*


* Vad händer om man tar bort go-kommandot från Seek-anropet i main-funktionen?
Innan test: Istället för trådar så kommer varje person köra sin Seek en efter en. Eftersom resultated hamna i ordning trots att vi körde med trådar
så kommer output bli likadant. Om man har flera processorer så kommer det gå lite långsammare också då flera Seek-functioner aldrig körs tillsammans.
Fast då det är oerhört små go-rutiner som körs så tar det kanske mer på systemet att starta dom än vad man tjänar på att köra dom.
Efter test: Precis som misstänkt så ändrades ingenting i output. Det enda praktiska som ändrats är att vi har garanterat en viss ordning på namnen.

* Vad händer om man byter deklarationen wg := new(sync.WaitGroup) mot var wg sync.WaitGroup och parametern wg *sync.WaitGroup mot wg sync.WaitGroup?
Innan test: Jag tror ingenting kommer att ändras. Vi ändrar bara wg från att vara en pekare till att vara en variabel.
Efter test: Jag hade fel, det blev deadlock vid wg.wait(). Då jag är lite java-handikappad missade jag att wg blir en lokal kopia i varje
funktion. Deadlocken beror på att wg (första kopian, eller ogrinalet) i RunMatching inte har markerat done() någon gång.

* Vad händer om man tar bort bufferten på kanalen match?
Innan test: Programmet kommer fastna vid sista namnet "Eva" då Seek kommer skriva till match men ingen kommer att läsa.
Programmet kommer dessutom fram till den punkten gå marginellt långsammare eftersom vi hela tiden måste vänta på läsning/skrivning.
Efter test: Det verkade stämma bra med min beskrivning.

* Vad händer om man tar bort default-fallet från case-satsen i main-funktionen?
Innan test: Vi kommer få deadlock om vi har ett jämnt antal namn. Enda alternativet i case-satsen är ju att läsa från kanalen och
har vi inget att läsa så blir det deadlock då inga andra go-rutiner är aktiva. I vårt fall med 5 namn får vi ingen deadlock dock.
Efter test: Fungerade vid ojämnt antal personer. Deadlock vid jämnt antal, precis som misstänkt.

*/
