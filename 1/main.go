package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i:=0 ; i<n; i++{
		fmt.Print(i*(i+1)/2+1)
		if (i<n-1){
			fmt.Print("-")
		}
	}
}