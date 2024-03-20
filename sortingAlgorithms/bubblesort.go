package sortingAlgorithms



//export Bubblesort
func Bubblesort (arr[]int)[]int{
	length := len(arr)	
	if length < 1 {
		return arr
	}

	for i:=0; i<length - 1; i++ {
		flag := 0
		for j:=0; j<length-1-i;j++{
			flag = 1
			if arr[j] > arr[j+1]{
				temp:=arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
		if flag == 0{ break}
	}
	return arr
}
