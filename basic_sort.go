package basic_algorithms

import "fmt"

// 1) 冒泡，持续比较相邻元素，大的挪到后面 O(n^2)
func BubbleSort(array []int) []int {
	n:=len(array)

	for i:=0;i<n;i++{
		for j:=1;j<n-i;j++{
			if array[j-1]>array[j]{
				array[j-1],array[j]=array[j],array[j-1]
			}
		}
	}
	return array
}

// 2) 选择，不断地选择剩余元素中的最小者(找到最小的元素，和第一个元素交换；从剩下的元素找到最小的元素，和第二个元素交换，以此类推) O(n^2)
func SelectSort(array []int)[]int  {
	n:=len(array)
	
	for i:=0;i<n;i++{
		min:=i //min存放最小值的元素下标

		for j:=i+1;j<n;j++{
			if array[j]<array[min]{
				min=j
			}
		}

		array[i],array[min]=array[min],array[i]
	}

	return array
}



// 3) 插入，对于未排序数据，在已排序数据中，从后往前扫描，插入    1，3，5，6，8 7（8后移一位，插到8的前面）
func InsertSort(array []int)[]int  {
	n:=len(array)

	for i:=1;i<n;i++{
		for j:=i-1;j>=0&&array[j]>array[j+1];j--{
			array[j], array[j+1] = array[j+1],array[j]

		}
	}
	return array
}


// 4)快速排序：找基准值(一般是取第一个或最后一个称为基准)；划分区（所有比基准小的元素置于基准左侧，比基准大的元素置于右侧；左右两个指针;左比基准值大，右比基准值小，交换;相遇停止交换）；递归调用
func Partition(arr []int, low, high int) int {
	baseNum := arr[low] //最左边元素作为基准值，指针从右边开始移动
	for low < high {
		for arr[high] >= baseNum && high > low {
			high--
		}

		arr[low] = arr[high]

		for arr[low] <= baseNum && high > low {
			low++
		}
		arr[high] = arr[low]
	}
	arr[high] = baseNum //第一轮low和high相等时均为5，此时交换停止 [5 1 2 4 3 6 9 7 10 8]
	return high
}

func QuickSort(arr []int, low, high int) []int {
	if low >= high {
		return arr
	}
	mid := Partition(arr, low, high)
	QuickSort(arr, low, mid-1)
	QuickSort(arr, mid+1, high)
	return arr
}


// 5)堆排序 堆排序在 top K 问题中使用比较频繁；父节点的键值总大于等于任何一子节点的键值；每个节点的左右子树都是一个二叉堆（最大堆或最小堆）
// 参考链接：http://wuchong.me/blog/2014/02/09/algorithm-sort-summary/

//最大堆调整：将堆的末端子节点作调整，使得子节点永远小于父节点
func heapify(array []int, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && array[left] > array[largest] {
		largest = left
	}
	if right < arrLen && array[right] > array[largest] {
		largest = right
	}
	if largest != i {
		array[i],array[largest]=array[largest],array[i]
		heapify(array, largest, arrLen)
	}
}


func buildMaxHeap(array []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(array, i, arrLen)
	}
	fmt.Println(array)
}

func HeapSort(arr []int) []int {
	arrLen := len(arr)
	buildMaxHeap(arr, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		arr[0],arr[i]=arr[i],arr[0]
		arrLen -= 1
		heapify(arr, 0, arrLen)
	}
	return arr
}


// 6)归并，两个有序数组归并成更大的有序数据   ???
//迭代法实现
func MergeSort(array []int) []int{
	n := len(array)
	if n < 2 {
		return array
	}
	key := n / 2
	left := MergeSort(array[0:key])
	right := MergeSort(array[key:])
	return merge(left, right)
}

func merge(left []int, right []int) []int{
	newArr := make([]int, len(left)+len(right))
	i, j, index :=0,0,0
	for {
		if left[i] > right[j] {
			newArr[index] = right[j]
			index++
			j++
			if j == len(right) {
				copy(newArr[index:], left[i:])
				break
			}

		}else{
			newArr[index] = left[i]
			index++
			i++
			if i == len(left) {
				copy(newArr[index:], right[j:])
				break
			}
		}
	}
	return newArr
}


//递归法实现
func MergeRecursive(data []int) []int {
	sum := len(data)
	if sum <= 1 {
		return data
	}
	left := data[0 : sum/2]
	lSize := len(left)
	if lSize >= 2 {
		left = MergeRecursive(left)
	}

	right := data[sum/2:]
	rSize := len(right)
	if rSize >= 2 {
		right = MergeRecursive(right)
	}

	j := 0
	t := 0
	arr := make([]int, sum)
	//fmt.Println(left, right, data)
	for i := 0; i < sum; i++ {
		if j < lSize && t < rSize {
			if left[j] <= right[t] {
				arr[i] = left[j]
				j++
			} else {
				arr[i] = right[t]
				t++
			}
		}  else if j >= lSize{
			arr[i] = right[t]
			t++
		}  else if t >= rSize{
			arr[i] = left[j]
			j++
		}
	}
	return arr
}
