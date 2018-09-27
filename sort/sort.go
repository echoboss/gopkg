package main

import (
	"sort"
	"fmt"
)

func main() {
	nums := []int{-10,100,3,9,-2,6,1,}
	sort.Ints(nums)
	fmt.Println(nums,sort.IntsAreSorted(nums))

	fs := []float64{1.1,-100.1,3.14,66.6,-1.3}
	sort.Float64s(fs)
	fmt.Println(fs)

	fmt.Println(sort.IsSorted(sort.IntSlice(nums)))
	fmt.Println(sort.Reverse(sort.IntSlice(nums)))
}