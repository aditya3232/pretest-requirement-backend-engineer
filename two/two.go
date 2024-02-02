package main

func findMax(arr []uint64) uint64 {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func main() {
	arr := []uint64{1, 2, 3, 4, 15}
	println(findMax(arr))
}
