package main

import "fmt"

func sortByHeight(a []int) []int {
	return bubbleSort(a)
}

func bubbleSort(a []int) []int {
	sorted := false

	for !sorted {
		sorted = true

	label:
		for i := 0; i < len(a)-1; i++ {
			if a[i] == -1 {
				continue
			}

			nextIndex := i + 1
			tmp := a[nextIndex]

			for tmp == -1 {
				nextIndex++
				if nextIndex >= len(a) {
					i++
					continue label
				}
				tmp = a[nextIndex]
			}

			if a[i] > tmp {
				a[nextIndex] = a[i]
				a[i] = tmp

				sorted = false
			}
		}
	}

	return a
}

func main() {
	test1 := []int{-1, 150, 190, 170, -1, -1, 160, 180}
	test2 := []int{-1, 150, 190, 170, -1, -1, 160, -1}
	test3 := []int{-1, 150, 190, 170, -1, -1, 160, 180, -1, 190}
	test4 := []int{-1, 150, 190, 170, -1, -1, -1, -1}

	fmt.Println(sortByHeight(test1))
	fmt.Println(sortByHeight(test2))
	fmt.Println(sortByHeight(test3))
	fmt.Println(sortByHeight(test4))
}
