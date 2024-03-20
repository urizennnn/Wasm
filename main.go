package main 

import (
	"fmt"
	"github.com/urizennnn/Wasm/sortingAlgorithms"

)

func main() {
	// fmt.Println("Hello World")
	arr:=[]int{10,6,5,7,8,0}
	// fmt.Print(sortingAlgorithms.Quicksort(arr))
	//fmt.Print(sortingAlgorithms.Bubblesort(arr))
	fmt.Print(sortingAlgorithms.Insertionsort(arr))
}
