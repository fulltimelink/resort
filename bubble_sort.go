package sort

/*
	冒泡排序只会操作相邻的两个数据。
	每次冒泡操作都会对相邻的两个元素进行比较，看是否满足大小关系要求，如果不满足就让它俩互换。
	一次冒泡会让至少一个元素移动到它应该在的位置，重复 n 次，就完成了 n 个数据的排序工作。
*/

// BubbleSort
// --  @# annotated by fulltimelink @ 2021-05-28 08:26:44
// --  @# bubble sort
func BubbleSort(list []int) {
	if 1 > len(list) {
		return
	}
	for k := range list {
		// --  @# flag to check sort is ok
		flag := false
		for kk, vv := range list[:len(list)-k-1] {
			if vv > list[kk+1] {
				list[kk], list[kk+1] = list[kk+1], list[kk]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

// BubbleSortFor
// --  @# annotated by fulltimelink @ 2021-05-28 08:26:44
// --  @# bubble sort use `for`
func BubbleSortFor(list []int) {
	if 1 > len(list) {
		return
	}
	for i := 0; i < len(list); i++ {
		// --  @# flag to check sort is ok
		flag := false
		for j := 0; j < len(list)-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
