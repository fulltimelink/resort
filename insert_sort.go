package sort

/*
	我们将数组中的数据分为两个区间，已排序区间和未排序区间。
	初始已排序区间只有一个元素，就是数组的第一个元素。
	插入算法的核心思想是取未排序区间中的元素，
	在已排序区间中找到合适的插入位置将其插入，并保证已排序区间数据一直有序。
	重复这个过程，直到未排序区间中元素为空，算法结束。
*/

// InsertSort
// --  @# annotated by fulltimelink @ 2021-05-28 10:24:16
// --  @#
func InsertSort(slice []int) {
	if 1 >= len(slice) {
		return
	}
	for k, v := range slice {
		ip := k - 1
		for ; ip >= 0; ip-- {
			if slice[ip] > v {
				slice[ip+1] = slice[ip]
				continue
			}
			break
		}
		slice[ip+1] = v
	}
}
