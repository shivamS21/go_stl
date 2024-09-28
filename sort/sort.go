package sort

// QuickSort sorts an array of integers in ascending order using the QuickSort algorithm.
func QuickSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }

    left, right := 0, len(arr)-1
    pivot := len(arr) / 2

    // Swap pivot and the right element
    arr[pivot], arr[right] = arr[right], arr[pivot]

    // Partition the array
    for i := range arr {
        if arr[i] < arr[right] {
            arr[i], arr[left] = arr[left], arr[i]
            left++
        }
    }

    // Swap pivot to its correct place
    arr[left], arr[right] = arr[right], arr[left]

    // Recursively sort the left and right partitions
    QuickSort(arr[:left])
    QuickSort(arr[left+1:])

    return arr
}
