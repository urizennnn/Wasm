package sortingAlgorithms



//export Bubblesort
func Bubblesort (arr[]int)[]int{
	length := len(arr)	
	if length < 1 {
		return arr
	}

	for i:=0; i<length - 1; i++ {
		for j:=0; j<length-1-i;j++{
			if arr[j] > arr[j+1]{
				temp:=arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}

	}
	return arr
}
