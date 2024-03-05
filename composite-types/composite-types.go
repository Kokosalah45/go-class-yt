package compositetypes

import "fmt"


func Run () {
	arr := [3]int{1,2,3}
	slice := []int{1,2,3}
	mapp := map[string]int{"one":1, "two":2, "three":3}

	slice = append(slice, 4)
	slice = append(slice, 5)
	slice = append(slice, 6) 
	slice = append(slice, 7)	

	println(len(arr) , cap(arr))
	println(len(mapp))
	println(len(slice) , cap(slice))

	for i , v := range arr {
		println(i , v)
	}

	for i , v := range slice {
		println(i , v)
	
	}
	for k , v := range mapp {
		println(k , v)
	
	}

	fmt.Printf("---------nil---------\n")

	var nilArr []int
	var nilMap map[string]int
	var nilSlice []int

	println(len(nilArr) , cap(nilArr))
	println(len(nilMap))
	println(len(nilSlice) , cap(nilSlice))

	for i , v := range nilArr {
		println(i , v)
	}

	for i , v := range nilSlice {
		println(i , v)
	}
	for k , v := range nilMap {
		println(k , v)
	}	
}