package piscine

func Max(arr []int) int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				smaller := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = smaller
			}
		}
	}
	return arr[len(arr)-1]
}
