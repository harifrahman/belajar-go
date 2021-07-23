package main

import "fmt"

func main() {
	// declare array, max value can hold 3, it always start from 0
	// so catName[3] thats mean, that array only can hold index 0, 1, 2
	var catName[3] string
	catName[0] = "betmen"
	catName[1] = "planet"
	catName[2] = "tikos"

	// doing overwrite for index 1
	catName[1] = "oleen"
	fmt.Println(catName)

	// declare array include values
	var nilaiSiswa = [2] int {
		98,
		34,
	}

	fmt.Println(nilaiSiswa)
}