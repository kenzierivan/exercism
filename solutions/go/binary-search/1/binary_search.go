package binarysearch

import "slices"

func SearchInts(list []int, key int) int {
    slices.Sort(list)

    left := 0
    right := len(list)-1
    
    for left <= right {
        middle := left + (right-left)/2
        if list[middle] == key {
        	return middle
    	} else if list[middle] > key {
        	right = middle - 1
    	} else if list[middle] < key {
            left = middle + 1
        }
    }
    return -1
}
