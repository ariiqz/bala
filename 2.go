package main

import "fmt"

func main() {
	var r float64
	var luas_lingkaran float64
	fmt.Scanln(&r)
	luas_lingkaran = (float64(22/7) * float64(r))
	fmt.Println("Luas lingkaran dengan jari-jari =", r, "adalah", luas_lingkaran)

}
