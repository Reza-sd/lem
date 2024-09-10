package main

func main() {
	println(slicesEqual([]int{3400, 2, 6}, []int{3, 2, 6}))
	println(slicesEqual([]int{3, 2, 6}, []int{3, 2, 6}))

}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
