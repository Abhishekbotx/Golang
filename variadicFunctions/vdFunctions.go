package main

import "fmt"

//variadic function can take variable number of arguments
//syntax: func functionName(param1 type, param2 type, param3 ...type) returnType
//...type means it can take variable number of arguments of type type
//eg: Println, Printf, Sprintf are variadic functions

func sum(nums ...int) int{ //...int means it can take variable number of arguments of type int
	total:=0// and they are in left  side of the type
	for _,num:=range nums{
		total+=num
	}

	return total
}
func main() {
 add:=sum(1,6,9,6)
 fmt.Println("sum is ",add)

 nums:=[]int{1,2,3,4,5}
 add2:=sum(nums...) //... is used to unpack the slice similar to spread operator in javascript
 fmt.Println("sum is ",add2)
}

///last time 2:26:11