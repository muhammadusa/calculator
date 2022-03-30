package basic_test

import (
	"testing"
	basic "day24/basic"
)

func test_Add(t *testing.T){
	if basic.Add(3,5) == 8{
		t.Error("True")
	}else{
		t.Error("False")
	}
}


func test_Sub(t *testing.T){
	if basic.Sub(2,5) == 10{
		t.Error("True")
	}else{
		t.Error("False")
	}
}


func test_Minus(t *testing.T){
	if basic.Minus(15,5) == 10{
		t.Error("True")
	}else{
		t.Error("False")
	}
}


func test_Tobe(t *testing.T){
	if basic.Tobe(15,5) == 3{
		t.Error("True")
	}else{
		t.Error("False")
	}
}


//////////////////////



