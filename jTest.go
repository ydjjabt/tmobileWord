package main

import(
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


func validWords(startWord string, endWord string) []string{

dictionary := openFile("english.json")

resultArray := []string{};


formatStartWord := strings.ToUpper(string(startWord[0])) + startWord[1:] 
_,exist := dictionary[formatStartWord]
if( exist ){
	resultArray = append(resultArray, startWord);
}
	
	possibleWord := startWord
	for i:= 1; i < len(endWord); i++{
		possibleWord = strings.Replace(possibleWord, startWord[i:i+1], endWord[i:i+1], 1)
		
		formatPossibleWord  :=  strings.ToUpper(string(possibleWord[0])) + possibleWord[1:] 
		_,exist := dictionary[formatPossibleWord]
		if(exist){
			resultArray = append(resultArray, possibleWord);
		}
	}

formatEndWord := strings.ToUpper(string(endWord[0])) + endWord[1:] 
_,existEndWord := dictionary[formatEndWord]
if( existEndWord ){
	resultArray = append(resultArray, endWord);
}

return resultArray;

}

func main(){

	fmt.Println( validWords("bat", "cow"))
	
}

