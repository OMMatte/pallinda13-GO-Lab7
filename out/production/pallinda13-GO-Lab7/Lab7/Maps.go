/**
 * Created with IntelliJ IDEA.
 * User: OMMatte
 * Date: 2013-03-19
 * Time: 14:26
 */
package Lab7

import (
//	"code.google.com/p/go-tour/wc"
	"strings"
//	"fmt"
)

func WordCount(s string) map[string]int {

	wordArray := strings.Split(s, " ")
	returnMap := make(map[string]int, len(wordArray))
	for _,word := range wordArray {
		returnMap[word] += 1
	}
	return  returnMap
}

func RunMaps() {
//	wordArray := WordCount("de de de de de och och")
//	for word := range wordArray{
//		fmt.Print(word + ": ")
//		fmt.Println(wordArray[word])
//

//	wc.Test(WordCount)
}
