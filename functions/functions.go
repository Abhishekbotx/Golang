package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func getLang() (string, string, string) {
	return "Go", "Golang", "Go Programming Language"
}

// processIt accepts a function (int -> int) 
// func processIt(fn func(a int) int) {
// 	// expects only functions which takes int as parameter and returns int
// 	fn(5);
// }

//you cant store the  function in a variable and then call it


//below we are not accepting anything but  returning a function from a function whose function signature is func(a int) int
func processIt() func(a int) int {
	return func(a int) int {
		return a * a
	}
}
func main() {

	// fnc := func(a int) int {
	// 	return a * a
	// }
	// processIt(fnc)
	
	fnc:=processIt()
	//Imp🍁: Here what is happening we are storing function returned by processIt in fnc variable and then calling it
	//So fnc is now a function which takes int as parameter and returns int
	//So we can call fnc like this

	fmt.Println("fnc result:", fnc(6))
	// fmt.Println("hofFunc result:", fn(6))


	// fmt.Println("hofFunc result:", processIt(fnc))
	lang1, lang2, lang3 := getLang()
	println(lang1, lang2, lang3)
	println(add(3, 4))
}