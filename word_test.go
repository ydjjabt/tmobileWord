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
//checking file type in Go
// check if os.Open() work


func Test_validWords_1(t *testing.T){
	resultArr := validWords("bat", "cow")
	expectedArr := []interface{}{"bat", "bot", "bow", "cow"}
	if !reflect.DeepEqual(resultArr, expectedArr){
		t.Error("answers is invalid")
		t.Error("your answer:", resultArr)
		t.Error("expected answer:", expectedArr)
	}
}


func Test_validWords_2(t *testing.T){
	resultArr := validWords("cat", "dog")
	expectedArr := []reflect.Value{"cat", "cog", "cot", "dog"}
	if !reflect.DeepEqual(resultArr, expectedArr){
		t.Error("answers is invalid")
		t.Error("your answer:", resultArr)
		t.Error("expected answer:", expectedArr)
	}
}


func Test_validWords_SwitchOrder(t *testing.T){
	resultArr := validWords("dog", "cat")
	expectedArr := []reflect.Value{"cat", "cog", "cot", "dog"}
	if !reflect.DeepEqual(resultArr, expectedArr){
		t.Error("answers is invalid")
		t.Error("your answer:", resultArr)
		t.Error("expected answer:", expectedArr)
	}
}


func Test_validWords_UnequalStrings(t *testing.T){
	resultArr := validWords("battle", "dog")
	expectedArr := []reflect.Value{"battle", "bottle", "dog"}
	if !reflect.DeepEqual(resultArr, expectedArr){
		t.Error("answers is invalid. ")
		t.Error("your answer:", resultArr)
		t.Error("expected answer:", expectedArr)
	}
}


func Test_validWords_UnequalStrings_2(t *testing.T){
	resultArr := validWords("dog", "horse")
	expectedArr := []reflect.Value{"dog", "dor", "horse"}
	if !reflect.DeepEqual(resultArr, expectedArr){
		t.Error("answers is invalid")
		t.Error("your answer:", resultArr)
		t.Error("expected answer:", expectedArr)
	}
}


func Test_validWords_UnequalStrings_more(t *testing.T){
	resultArr := validWords("cat", "hopeful")
	expectedArr := []reflect.Value{"cat", "cop", "cot", "hopeful"}
	if !reflect.DeepEqual(resultArr, expectedArr){
		t.Fatalf("answers is invalid")
		t.Error("your answer:", resultArr)
		t.Error("expected answer:", expectedArr)
	}
}