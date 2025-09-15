package main

import (
	"fmt"
)

func main(){
	 
	nums:=[...]int{2,3,4,5,6}
	sum:=0
	for _,num:=range nums{ // _ is used to ignore the index, its basically for index
		sum+=num
	}
	fmt.Println("sum is ",sum)

	
	nums2:=[...]int{2,3,4,5,6}
	//IMPORTANT:
	// for key:=range nums2{ this will only return keys, for value we can use for _,val=range nums2{} like this
	for key,val:=range nums2{
		fmt.Printf("index is %v and value is %v\n",key,val)
	}
}