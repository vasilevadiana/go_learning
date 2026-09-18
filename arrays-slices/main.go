package main

import "fmt"

func main() {
	/* little Summary:
	Declaring and inializing syntax:
	arr_name := [arr_size]type{values}
	len(arr) -> len of arr
	cap(arr) -> capacity of arr
	*/
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)
	fmt.Println((len(arr)))
}
