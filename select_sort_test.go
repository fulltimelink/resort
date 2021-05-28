package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	slices := []int{11, 2, 4, 5, 8, 7, 9}
	fmt.Println(slices)
	SelectSort(slices)
	fmt.Println(slices)
}
