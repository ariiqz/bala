package main

import "fmt"

func main() {
	var S string
	var A, B, operasi int64

	fmt.Scanln(&S, &A, &B)
	operasi = A + B
	fmt.Println("kata =", S)
	fmt.Println("jumlah =", operasi)

}
