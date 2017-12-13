package bees

import "fmt"

func SliceReverse(arr []*interface{}) {
	if arr == nil {
		panic(fmt.Errorf("the arr can't be empty!"))
	}
	count := len(arr)
	mid := count / 2
	for i := 0; i < mid; i++ {
		tmp := arr[i]
		arr[i] = arr[count-1]
		arr[count-1] = tmp
		count--
	}
}
