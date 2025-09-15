package main

import "fmt"

//
// 🖇️ Pointers in Go
// 👉 A pointer is a variable that stores the memory address of another variable.
//
// 📌 Operators:
//   &  → Address-of operator (get the address of a variable)
//   *  → Dereference operator (access/change the value at an address)
//
// 📌 Important:
//   *T  → Type: pointer to T (e.g., *int = pointer to int)
//   *p  → Dereference: access the value at the pointer p
//

// changeNum updates the value of num through a pointer
// 🌟 *int in the function parameter
// This is part of the type declaration.
// It means: "num is a pointer to an int".
func changeNum(num *int) {
	// *num = 8 means:
	//  - Look at the address stored in num (e.g., 0xc00000a088)
	//  - Go to that memory location
	//  - Change the value stored there to 8
	*num = 8
}

func main() {
	// Normal variable
	num := 5

	fmt.Println("before changeNum:", num)
	fmt.Println("before changeNum Address:", &num)

	// Passing the address of num into the function
	changeNum(&num)

	fmt.Println("after changeNum:", num)
	fmt.Println("after changeNum Address:", &num)

	//
	// 🍁 Important:
	// - We passed the address of num to changeNum
	// - Inside changeNum, *num dereferenced that pointer
	//   and changed the value at that address
	// - So the original variable in main got updated
	//
}
