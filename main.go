package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Hello, World!")

	const (
		A int = 1
		B     = 3.14
		C     = "Hi!"
	)

	i:= 42

	// %v is for value  and %T is for type
	fmt.Printf("the value of  i is %v and type is %T",i)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)



	var arr1 = [...]int{1,2,3}
  	arr2 := [...]int{4,5,6,7,8}

	var arr3=[...]string{"a","b","c"}

	fmt.Println(arr1)
    fmt.Println(arr2)
    fmt.Println(arr3)

	for j := 0; j < len(arr2); j++ {
		fmt.Println(arr2[j])
	}

	for k := range 5 {
		fmt.Println("heyyy",)
		fmt.Println(k)
	}

	// The range keyword is used to more easily iterate through the elements of an array, slice or map. It returns both the index and the value.
	// The range keyword is used like this:

	// Syntax
	// for index, value := range array|slice|map {
	// // code to be executed for each iteration
	// }
	 fruits := [3]string{"apple", "orange", "banana"}
	 
	 vowels:=[...]string{"a","e","i","0","u"}
	 
	 for idx,value:=range vowels{
		 fmt.Printf("%v\t%v\n",idx,value)
		}
	 for _, val := range fruits {
			fmt.Printf("%v\n", val)
		}


	// Slices 
	//They are most used construct they are like arrays and you you dont need to declare length
	var exnums []int
	//this willl initialy nil if you print it
	fmt.Println(len(exnums))
	var nums = []int{1,2,3}
	fmt.Println(nums)


	// so there are other ways to declare slices , make is one of them
	myslice:=make([]int,5,10) // here 5 is length and 10 is initial capacity
	fmt.Println(myslice)
	// fmt.Println(len(myslice))	
	myslice=append(myslice, 2,3) // append is used to add elements to slice and can add multiple elements
	fmt.Println(myslice)	

	//if the capacity is full it will double the capacity

	myslice[0]=100
	fmt.Println(myslice)

	myslice2:=make(([]int),len(myslice))

	//copy function

	//syntax is copy(destination,source)
	copy(myslice2,myslice)

	fmt.Println("my slice 2",myslice2)

	// if the length is zero then it will not copy anything

	//slice operator

	sarr:=[]int{1,2,3,4,5,6,7,8,9}
	sarr2:=[]int{1,2,3,4,5,6,7,8,9}

	fmt.Println(sarr[2:5]) // it will print from index 2 to index 4 and 4 is excluded
	fmt.Println(sarr[5:]) // it will print from index 5 to end

	fmt.Println(slices.Equal(sarr,sarr2)) // it will check if two slices are equal or not) 		


	//%d is a verb (placeholder) in formatting.
	//%s is for strings.
	//%T is for type.
	//%v is for value.
	//%f is for float.

	fmt.Printf("the value of  i is %d and type is %T",i,i)

}
