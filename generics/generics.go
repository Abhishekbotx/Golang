package main

import "fmt"

func printSliice (items []int){
	for _,item:=range items{
		fmt.Println(item)
	}
}
func printStringSlice (items []string){
	for _,item:=range items{
		fmt.Println(item)
	}
}

// func printGenericSlice[T any](items []T){ // T is a type parameter, any means it can be of any type
// 	for _,item:=range items{
// 		fmt.Println(item)
// 	}
// }
func printGenericSlice[T int | string | bool](items []T){ // T is a type parameter, any means it can be of any type
	for _,item:=range items{
		fmt.Println(item)
	}
}
//🌟Note:
//Comparable is a constraint that permits any type that supports the comparison operators == and !=. 
// All boolean, numeric, string, pointer, channel, and interface types are comparable.
func printable [T comparable](items []T){ // T is a type parameter, comparable means it can be of any type that can be compared
	for _,item:=range items{
		fmt.Println(item)
	}
}

type stack[T string | int] struct{ // its a generic struct
	elements []T // here T can be of type string or int
}

func main(){
	nums:=[]int{1,2,3,4,5}
	names:=[]string{"abhishek","raj","is","awesome"}
	// printStringSlice(names) // here we are passing slice to the function
	// printSliice(nums) // here we are passing slice to the function
	printGenericSlice(names) // here we are passing slice to the function
	printGenericSlice(nums) // here we are passing slice to the function

	myStack:=stack[int]{
		elements:[]int{1,2,3,4,5},
	}

	fmt.Println("mystack:",myStack)
}
