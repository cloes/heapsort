package main

func heapSort2(arr []int) []int {
	len := len(arr)
	half := len / 2

	for i := half; i > -1; i-- {
		//构建堆
		heap2(arr, i, len-1)
	}

	for j := len - 1; j > 0; j-- {
		arr[j], arr[0] = arr[0], arr[j]
		heap2(arr, 0, j-1)
	}
	return arr
}

func heap2(array []int, i, end int) {
	for i < end {
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
			//交换数据，使得顶端的值为最大值(因为原来已经是堆结构，现在只是换掉了root最大值，把较小的值赋值给root)
			array[i], array[tmp] = array[tmp], array[i]
			//交换后，恢复堆结构
			i = tmp
		} else {
			return
		}
	}
}
