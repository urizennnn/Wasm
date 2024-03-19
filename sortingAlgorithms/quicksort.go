package sortingAlgorithms

// import "fmt"
//export Quicksort
func Quicksort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    pivot := arr[0]
    left := []int{}
    right := []int{}

    for _, num := range arr[1:] {
        if num < pivot {
            left = append(left, num)
        } else {
            right = append(right, num)
        }
    }

    left = Quicksort(left)
    right = Quicksort(right)

    left = append(left, pivot)
    return append(left, right...)
}
