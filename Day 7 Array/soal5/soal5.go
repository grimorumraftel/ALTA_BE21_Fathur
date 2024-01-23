package main

import "fmt"

func main() {
	var arr [10]int

	fmt.Println("Masukan 10 Elemen Array:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}

	var index1, index2 int
	fmt.Println("Masukan 2 Indeks Array Yang Ingin Dijumlah:")
	fmt.Scan(&index1, &index2)

	sum := arr[index1] + arr[index2]
	fmt.Printf("Sum of the two elements: %d\n", sum)
}
