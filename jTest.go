package main

import(
"strings"
"fmt")


 func binarySearch(target string, wordList []string, startIndex int, endIndex int) int {
 	//goLang doesnt have indexOf item in an array. so i need to implement this. 
 	if( startIndex <= endIndex){ 		
 		middleIndex := (startIndex + endIndex ) / 2;
 		// fmt.Println();
 		//  		fmt.Println("wordList", wordList)
 		//  		fmt.Println("middleIndex", middleIndex, ". middle value is", wordList[middleIndex], ". target", target)
 		//  		fmt.Println("startIndex", startIndex, ". endIndex", endIndex)

 		if( wordList[middleIndex] == target ){
 			return middleIndex;
 		} else if(wordList[middleIndex] < target){
 			return binarySearch(target, wordList, middleIndex + 1, endIndex);
 		} else{
 			return binarySearch(target, wordList, startIndex , middleIndex - 1);
 		}
 	}else{ return -1;}
 	
 }
func main(){
	// small sample list of valid words. This list have to SORTED alphabetical order from A to Z
	// use a dictionary API if this was for commercial operation to check if a word is a valid English word
	wordList := []string{"bat", "bog", "bot", "bow", "cat", "cog", "cot", "cow", "dog"};


    inputA := "bat";
    inputB := "dog";
	
	startWord := inputA;
	endWord := inputB;
	resultArray := []string{};

	//check if the startWord is valid English word.
	// use a dictionary API if this was for commercial operation to check if a word is a valid English word
	if( binarySearch(startWord, wordList, 0, len(wordList) - 1) >= 0){
		resultArray = append( resultArray, startWord);
	}

	//check in between startWord and endWord
	// use a dictionary API if this was for commercial operation to check if a word is a valid English word
	possibleWord := startWord
	for i:= 1; i < len(endWord); i++{
		possibleWord = strings.Replace(possibleWord, startWord[i:i+1], endWord[i:i+1], 1)
		if( binarySearch(possibleWord, wordList, 0, len(wordList)  - 1) >= 0){
			resultArray = append(resultArray, possibleWord);
		}
	}

	//check if endWord is valid English word
	// use a dictionary API if this was for commercial operation to check if a word is a valid English word
	if( binarySearch(endWord, wordList, 0, len(wordList) - 1) >= 0){
		resultArray = append( resultArray, endWord);

	}

	fmt.Println("resultArray ", resultArray)
}

