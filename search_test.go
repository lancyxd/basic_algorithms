package basic_algorithms

import (
	"fmt"
	"testing"
)

var alist=[]int{1,2,3,4,5,6,7}




func TestBinarySearch(t *testing.T) {
	target:=6
	targetIdx:=BinarySearch(alist,target)
	fmt.Println("target index:",targetIdx)
}

func TestBinarySearchRecursive(t *testing.T) {
	target:=7
	targetIdx:=BinarySearchRecursive(alist,target,0,6)
	fmt.Println("target index:",targetIdx)
}


