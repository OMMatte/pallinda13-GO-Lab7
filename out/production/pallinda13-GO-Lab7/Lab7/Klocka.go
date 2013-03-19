/**
 * Created with IntelliJ IDEA.
 * User: OMMatte
 * Date: 2013-03-19
 * Time: 15:30
 */
package Lab7

import( "time"
		"fmt"
)
 func Remind(text string, paus time.Duration){
	for{
		time.Sleep(paus)
		fmt.Print("Klockan är ")
		fmt.Print(time.Now().Format("15:04"))
		fmt.Println(": " + text)
 	}
}

func RunKlocka(){
	go Remind("Dags att äta", time.Hour*3)
	go Remind("Dags att arbeta", time.Hour*5)
	go Remind("Dags att sova", time.Hour*6)
	select{}
}

