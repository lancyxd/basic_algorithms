package basic_algorithms

import (
	"fmt"
	"testing"
)

var array=[]int{6, 5, 3, 1, 8, 7, 2, 4}


func TestBubbleSort(t *testing.T)  {
	rlist:=BubbleSort(array)
	fmt.Println("TestBubbleSort result:",rlist)
}

func TestSelectSort(t *testing.T)  {
	rlist:=SelectSort(array)
	fmt.Println("TestSelectSort result:",rlist)
}

func TestInsertSort(t *testing.T)  {
	rlist:=InsertSort(array)
	fmt.Println("TestInsertSort result:",rlist)
}

func TestMergeSort(t *testing.T)  {
	rlist:=MergeSort(array)
	fmt.Println("TestMergeSort result:",rlist)
}

func TestMergeRecursive(t *testing.T)  {
	rlist:=MergeRecursive(array)
	fmt.Println("TestMergeRecursive result:",rlist)
}

func TestQuickSort(t *testing.T) {
  rlist:=QuickSort(array,0,len(array)-1)
fmt.Println("TestMergeRecursive result:",rlist)
}

func TestHeapSort(t *testing.T)  {
	fmt.Println(HeapSort(array))

}