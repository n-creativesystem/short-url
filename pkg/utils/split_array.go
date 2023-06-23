package utils

func SplitArray[T any](arr []T, splitCount int) [][]T {
	if splitCount == 0 {
		splitCount = 10
	}
	var result [][]T
	for i := 0; i < len(arr); i += splitCount {
		end := i + splitCount
		if end > len(arr) {
			end = len(arr)
		}
		result = append(result, arr[i:end])
	}
	return result
}
