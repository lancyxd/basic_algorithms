package basic_algorithms

//1）二分查找  迭代方法,返回目标元素的下标
func BinarySearch(array []int,target int) int {
	high:=len(array)-1
	low:=0

	for low<=high{
		mid:=low + (high-low)/2
		midValue:=array[mid]
		if target==midValue{
			return mid
		}else if target>midValue{
			low=mid+1
		}else if target<midValue{
			high=mid-1
		}
	}
	return -1
}



//递归方法，已排序好的数组(数组是升序的)，起始位置，结束位置n-1 目标数字(当边界条件不满足时，递归前进；当边界条件满足时，递归返回)
//递归升序查找   start为0，end为len(arr)-1
func BinarySearchRecursive(array []int,target,start,end  int)  int  {
	mid:=(start+end)/2

	if target<array[start]||target>array[end]||start>end{
		return -1
	}

	if target<array[mid]{
		return BinarySearchRecursive(array,target,start,mid-1)
	}else if target>array[mid]{
		return BinarySearchRecursive(array,target,mid+1,end)
	}else {
		return mid
	}
}








