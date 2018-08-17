package main

import(
"reflect"
"sync"
"sort"
"encoding/json"
"io/ioutil"
"os"
"strings"
"fmt")

func openFile(fileName string) map[string]interface{}{
	jsonFile, err := os.Open(fileName)
	if err != nil{
		 panic(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil{
		panic(err)
	}
	var dictionary map[string]interface{}

	json.Unmarshal(byteValue, &dictionary)
	return dictionary
}


func validWords(startWord string, endWord string) []reflect.Value{
var wg sync.WaitGroup
dictionary := openFile("english.json")

// switch startWord and endWord if endWord is alphabetical order is smaller than startWord
if endWord < startWord{
	swap := startWord
	startWord = endWord
	endWord = swap
}

resultMap := map[string]interface{}{}
formatStartWord := strings.ToUpper(string(startWord[0])) + startWord[1:] 
_,exist := dictionary[formatStartWord]
if( exist ){
	resultMap[startWord] = nil;
}

formatEndWord := strings.ToUpper(string(endWord[0])) + endWord[1:] 
_,existEndWord := dictionary[formatEndWord]
if( existEndWord ){
	resultMap[endWord] = nil;
}
	
	possibleWord := startWord
	startWordLength := len(strings.Split(startWord, "") ) //using len([]rune(startWord)) doesnt give correct 'end result'. same thing with utf8.runecountsomething
	endWordLength := len( strings.Split(endWord, "") )//therefore i split string to array and it work on equal length and unequal length of two strings.
	var loopEnd int;
		if startWordLength >= endWordLength{ 
			loopEnd = endWordLength
		}else{ 
			loopEnd = startWordLength
		}
		ch := make( chan string, loopEnd )
		for i:= 1; i < loopEnd; i++{
			possibleWord = strings.Replace(possibleWord, startWord[i:i+1], endWord[i:i+1], 1)
			ch <- possibleWord
			wg.Add(1)
			go func(ch chan string){
				defer wg.Done()
				pWord := <- ch
				formatPossibleWord  :=  strings.ToUpper(string(pWord[0])) + pWord[1:] 
				_,exist := dictionary[formatPossibleWord]
				if(exist){
					resultMap[pWord] = nil;
				}
			}(ch)
		}
	
// when main is run, this validWords function is in main as part of main goroutine
// then just put the waitGroup.Wait() here for all those goroutine returning valid words
wg.Wait()

//put only keys of map into array to get unique valid words only. no duplicate
resultArray  := reflect.ValueOf(resultMap).MapKeys()
// sort it from a to z alphabetical order
sort.Slice(resultArray, func(i, j int) bool {
    return resultArray[i].Interface().(string) < resultArray[j].Interface().(string) 
})

return resultArray ;

}

func main(){
	fmt.Println( validWords("cat", "dog"))
}

