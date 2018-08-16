package main


import(
"testing"
"reflect"
)


func Test_openFile(t *testing.T){
dictionary := openFile("english.json") 
if len(dictionary) == 0{
	t.Error("dictionary is empty")
}
}


func Test_validWords(t *testing.T){
	resultArr := validWords("bat", "cow")
	expectedArr := []string{"bat", "bot", "bow", "cow"}
	if !reflect.DeepEqual(resultArr, expectedArr){
		t.Error("answers is invalid")
	}
}
