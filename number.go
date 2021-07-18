package main

import "fmt"

func main() {
	fmt.Println("angka satu", 1)
	fmt.Println("angka dua", 2)
	fmt.Println("angka tiga koma lima", 3.5)

	// boolean
	fmt.Println("Benar = ", true)
	fmt.Println("Salah =", false)

	// string
	fmt.Println(len("Arif"))
	fmt.Println("Arif", "Rahman", "Hakim")
	fmt.Println("Arif Rahman Hakim"[0]) //will output 65, ASCII "A"
}
