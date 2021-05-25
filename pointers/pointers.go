package main

import "fmt"

func main() {
	var p *[]int = new([]int)
	var p2 []int = make([]int, 5, 10)
	fmt.Println(p == nil)	
	*p = make([]int, 10)		
	fmt.Println(p2)
}
