package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absesnsi = 80

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusAbsesnsi = absesnsi > 80

	var lulus = lulusNilaiAkhir && lulusAbsesnsi

	fmt.Println(lulus)
}
