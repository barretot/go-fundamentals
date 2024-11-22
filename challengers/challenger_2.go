package challengers

import "fmt"

func ArrayPointer(arr *[3]int) {
	for i := range *arr {
		(*arr)[i] += 1
	}

	fmt.Println("Array alterado dentro da função:", *arr)
}

func SlicePointer(slice *[]int) {
	if len(*slice) > 0 {
		*slice = (*slice)[:len(*slice)-1]
	}

	fmt.Println("Slice alterado dentro da função:", *slice)
}
