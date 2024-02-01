package main

import "fmt"

func isPrima(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func generatePrimeAfter(start int) int { // 1
	var result = -1    // saya anggap tidak ada
	for result == -1 { // -1 == -1 V | 2 == -1 X
		start++             // star+1 -> 1+1
		if isPrima(start) { //isPrima(2) V
			result = start // result=2
		}
	}
	return result
}

func primaSegiEmpat(wide, high, start int) string { //start 13
	var hasilJumlah = 0
	var output = ""
	for i := 0; i < high; i++ { // 0 / 1
		for j := 0; j < wide; j++ { // 0 | 1 / 0
			// current adalah bilangan prima setelah nilai start
			current := generatePrimeAfter(start) // current = 17 | current = 19 / current = 23
			output += fmt.Sprint(current, " ")   // tampilkan(17) | tampilkan(19) | tampilkan(23)
			hasilJumlah += current               // 0+=17 | 17+=19 / 36+=23
			start = current                      // 13=17 | 17=19 / 19=23
		}
		output += fmt.Sprintln() // enter
	}
	output += fmt.Sprintln(hasilJumlah)
	return output
}

func main() {
	fmt.Println(primaSegiEmpat(2, 3, 13))
	/*
	   17 19
	   23 29
	   31 37
	   156
	*/
	fmt.Println(primaSegiEmpat(5, 2, 1))
	/*
	   2  3  5  7 11
	   13 17 19 23 29
	   129
	*/
}
