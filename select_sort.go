package sort

/*
	选择排序算法的实现思路有点类似插入排序，也分已排序区间和未排序区间。
	但是选择排序每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾。
	这样一来，当遍历完未排序区间，就意味着已经完成整个序列的排序了。
*/

// SelectSort
// --  @# annotated by fulltimelink @ 2021-05-28 14:01:29
// --  @# select sort
func SelectSort(list []int) {
	if 1 >= len(list) {
		return
	}
	for k := range list[:len(list)-1] {
		minIndex := 0
		for kk, vv := range list[k:] {
			if vv < list[minIndex+k] {
				minIndex = kk
			}
		}
		if 0 != minIndex {
			list[minIndex+k], list[k] = list[k], list[minIndex+k]
		}
	}
}
