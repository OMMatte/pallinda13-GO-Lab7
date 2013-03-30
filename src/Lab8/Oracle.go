/**
 * Created with IntelliJ IDEA.
 * User: OMMatte
 * Date: 2013-03-21
 * Time: 15:00
 */
// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package Lab8

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func RunOracle() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	oracle := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		oracle <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	QAndA := make(chan string)

	go parseQuestions(questions, QAndA)
	go prophecy(QAndA)
	go printAnswers(QAndA)

	return questions
}

func printAnswers(QAndA <- chan string) {
	for q := range QAndA {
		fmt.Println(q)
		fmt.Print("> ")
	}
}

func parseQuestions(questions <- chan string, QAndA chan<- string) {
	answers := []string{
		"Inshallah.",
		"What do you think?.",
		"Maybe.",
		"I have to consult with the gods. Come back later.",
		"How the **** should i know?",
		"If you belive in yourself, i believe in you.",
		"I wonder that myself.",
		"You will find your answer soon enough.",
	}

	for q := range questions {
		go answerQuestion(q, QAndA, answers)
	}
}

// This is my "prophecy" function. I thought the name and what it actually did not work together.
// So the method signatures are the same except the last one. But it would be unwise to create the array of strings each
// time the function is generated. If this was totally forbidden, then it would work just fine to create the array inside.
func answerQuestion(question string, QAndA chan<- string, answers[] string){
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	qAndA := "Your question was: " +question + " | Pynthia answers: " + answers[rand.Intn(len(answers))]
	QAndA <- qAndA
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// Changed. I know we are not allowed to change the method signatures but i renamed it to answerQuestion() and changed the
// behaviour on this fucntion. Honestly though, i did change the signature a bit but just for efficiency reasons, it does not
// change the logic one bit.
func prophecy(QAndA chan<- string) {

	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
		"I am a false prohet.",
		"Do not listen to me.",
		"My advices suck.",
		"To be or not to be.",
		"After rain comes something.",
		"I sense you do not believe me.",
	}

for{
	time.Sleep(time.Duration(5+rand.Intn(10)) * time.Second)
	// Cook up some pointless nonsense.

	QAndA <- "... " + nonsense[rand.Intn(len(nonsense))]
}
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}

/*


Gör klart Oracle-metoden. Du får inte ändra i main-metoden och du får inte heller ändra metodsignaturerna. Observera att svaren inte ska komma direkt, utan med fördröjning.
Glöm inte heller att oraklet ska skriva ut meddelanden även om det inte kommer några frågor. Du får gärna dela upp din lösning på flera metoder.
Ditt program ska innehålla två stycken kanaler: en kanal för frågor samt en kanal för svar och förutsägelser. I Oracle-metoden ska du starta tre stycken permanenta gorutiner:

En gorutin som tar emot alla frågor och för varje inkommande fråga skapar en separat gorutin som besvarar frågan.
En gorutin som genererar förutsägelser.
En gorutin som tar emot alla svar och förutsägelser och skriver ut dem på stdout.
Oracle-metoden är den viktigaste delen av uppgiften. Om du vill får du också förbättra svarsalgoritmen. Även här får gärna dela upp algoritmen på flera metoder. Här är några tips:

Paketen strings och regexp kan vara användbara.
Programmet kan verka mera mänskligt om oraklet skriver ut sina svar en bokstav i taget.
Ta en titt på ELIZA, det första programmet av det här slaget.
*/
