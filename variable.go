package main

import "fmt"

func main() {
	// how to declare [var nameOfVar typeData]
	var fisrtName string

	// no need write type data if we intiate like this
	var lastName = "Hakim"
	fisrtName = "Arif"

	// no need to write "var" if declare like this
	middleName := "Rahman"

	// for multiple variable declaration
	var (
		type1 = "aku"
		type2 = "kamu"
	)

	fmt.Println(fisrtName)
	fmt.Println(lastName)
	fmt.Println(middleName)

	fmt.Println(type1)
	fmt.Println(type2)


	// constanta same like var, but it cant change, and value must be initiate
	const fullName = "Arif Rahman Hakim"

	// multiple const ?
	const (
		isActive = true
		isAdmin = false
	)

	// constanta tidak harus digunakan, krn ada kemungkinan akan di gunakan kedepannya dari file lain. jadi tidak error 
	// constanta must be prefix with "const"
}