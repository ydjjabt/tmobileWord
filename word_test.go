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



func Test_validWords_equalStrings1(t *testing.T){
	resultArr := validWords("bat", "cow")
	expectedArr := []string{"bat", "bot", "bow", "cow"}
	if !reflect.DeepEqual(resultArr, expectedArr){
		t.Error("answers is invalid")
		t.Error("your answer:", resultArr)
		t.Error("expected answer:", expectedArr)
	}
}


func Test_validWords_equalStrings2(t *testing.T){
	resultArr := validWords("cat", "dog")
	expectedArr := []string{"cat", "cog", "cot", "dog"}
	if !reflect.DeepEqual(resultArr, expectedArr){
		t.Error("answers is invalid")
		t.Error("your answer:", resultArr)
		t.Error("expected answer:", expectedArr)
	}
}


func Test_validWords_SwitchAlphabeticalOrder(t *testing.T){
	resultArr := validWords("dog", "cat")
	expectedArr := []string{"cat", "cog", "cot", "dog"}
	if !reflect.DeepEqual(resultArr, expectedArr){
		t.Error("answers is invalid")
		t.Error("your answer:", resultArr)
		t.Error("expected answer:", expectedArr)
	}
}


func Test_validWords_UnequalStrings_1(t *testing.T){
	resultArr := validWords("battle", "dog")
	expectedArr := []string{"battle", "bottle", "dog"}
	if !reflect.DeepEqual(resultArr, expectedArr){
		t.Error("answers is invalid. ")
		t.Error("your answer:", resultArr)
		t.Error("expected answer:", expectedArr)
	}
}


func Test_validWords_UnequalStrings_2(t *testing.T){
	resultArr := validWords("dog", "horse")
	expectedArr := []string{"dog", "dor", "horse"}
	if !reflect.DeepEqual(resultArr, expectedArr){
		t.Error("answers is invalid")
		t.Error("your answer:", resultArr)
		t.Error("expected answer:", expectedArr)
	}
}


func Test_validWords_UnequalStrings_3(t *testing.T){
	resultArr := validWords("cat", "hopeful")
	expectedArr := []string{"cat", "cop", "cot", "hopeful"}
	if !reflect.DeepEqual(resultArr, expectedArr){
		t.Fatalf("answers is invalid")
		t.Error("your answer:", resultArr)
		t.Error("expected answer:", expectedArr)
	}
}