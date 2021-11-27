package main

import "fmt"

func main() {
	abc := []int{1, 2, 3, 4, 5}

	res := sum(abc...) // === | sum := sum(1, 2, 3, 4, 5)

	println(fmt.Sprintf("SUM: %d", res))

	//=========================
	data := map[string]int{"key1": 1, "key2": 2, "key3": 3, "key4": 4}

	showMap(data)

	var a []int

	a = append(a, 1, 2)
	a = append(a, 3)
	a = append(a, 4)
	a = append(a, 5, 6)

	fmt.Println(a)
}

func sum(data ...int) int {
	s := 0

	for _, item := range data {
		s += item
	}

	return s
}

func showMap(m map[string]int) {
	for key, val := range m {
		fmt.Println(fmt.Sprintf("%s => %d", key, val))
	}
}
