# resort
review &amp; achieve sort algorithm

**主要使用for-range实现，从目标元素、目标位置的视角来进行排序**
* 目标元素 当前排序应该被操作的元素
* 目标位置 当前目标元素应该被放置的位置

> 经过benchmark 发现 for-range 比普通的for loop  op/ns 要低...

## bubble sort 冒泡排序

> 冒泡排序只会操作相邻的两个数据。  
每次冒泡操作都会对相邻的两个元素进行比较，看是否满足大小关系要求，如果不满足就让它俩互换。  
一次冒泡会让至少一个元素移动到它应该在的位置，重复 n 次，就完成了 n 个数据的排序工作。
> ***
> **循环中两两交换移动目标元素到目标位置，循环结束排序结束，循环元素个数和循环次数成反比**

```go
package sort
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
```

## insert sort 插入排序

> 我们将数组中的数据分为两个区间，已排序区间和未排序区间。  
初始已排序区间只有一个元素，就是数组的第一个元素。  
插入算法的核心思想是取未排序区间中的元素，  
在已排序区间中找到合适的插入位置将其插入，并保证已排序区间数据一直有序。  
重复这个过程，直到未排序区间中元素为空，算法结束。 
> ***
> **正向循环目标元素，反向循环移动目标位置，循环结束排序结束**

```go
package sort
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
```

## select sort 选择排序

> 选择排序算法的实现思路有点类似插入排序，也分已排序区间和未排序区间。  
但是选择排序每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾。  
这样一来，当遍历完未排序区间，就意味着已经完成整个序列的排序了。  
> ***
> **正向循环目标位置，子循环查找目标元素，循环结束排序结束**

```go
package sort
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
```

