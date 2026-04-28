package main

import "fmt"

func main() {
	type NoKTP string

	var ktpBalpe NoKTP = "123"

	var contoh string = "222"

	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpBalpe)
	fmt.Println(contoh)
	fmt.Println(contohKtp)
}
