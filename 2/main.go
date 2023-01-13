package main

import "fmt"

func removeDuplicateValues(intSlice []int) []int {
    keys := make(map[int]bool)
    list := []int{}
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

func main() {
	var n,x,play int
	var arr,score,res []int
	fmt.Scan(&n)
	for i:=0 ; i<n ;i++ {
		fmt.Scan(&x)
		arr = append(arr, x)
	}
	score = removeDuplicateValues(arr)
	// fmt.Print(res)
	fmt.Scan(&play)
	for i:=0 ; i<play ;i++ {
		val := 1
		fmt.Scan(&x)
		for _, v := range score {
			if v > x{
				val++
			}
		}
		res = append(res, val)
	}
	for i,v := range res{
		fmt.Print(v)
		if i < len(res)-1 {
			fmt.Print(" ")
		}
	}
}