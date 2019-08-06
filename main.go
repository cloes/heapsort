package main

import "fmt"

func heapSort(arr []int) []int {
	len := len(arr)
	half := len / 2

	for i := half; i > -1; i-- {
		//构建堆
		heap(arr, i, len-1)
	}

	for j := len - 1; j > 0; j-- {
		arr[j], arr[0] = arr[0], arr[j]
		heap(arr, 0, j-1)
	}
	return arr
}

func heap(array []int, i, end int) {
	left := 2*i + 1
	//如果左边的下标已经大于最大下标，表明没有左子树，那么直接返回
	if left > end {
		return
	}
	tmp := left
	right := 2*i + 2

	//如果右边的值最大，那么把最大值的下标赋给tmp变量
	if right <= end && array[tmp] < array[right] {
		tmp = right
	}

	if array[i] < array[tmp] {
		//交换数据，使得顶端的值为最大值
		array[i], array[tmp] = array[tmp], array[i]
		heap(array, tmp, end)
	}
}

func main() {
	origin := []int{4, 8, 5, 7, 20, 6, 25, 15, 9}
	origin2 := make([]int,len(origin))
	//origin2 := make([]int,len(origin),len(origin))
	copy(origin2,origin)

	fmt.Println(origin)
	result := heapSort(origin)
	fmt.Println(result)

	fmt.Println(origin2)
	result = heapSort2(origin2)
	fmt.Println(result)
}
