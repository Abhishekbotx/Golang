package main

import ("fmt")

type OrderStatus int

const (
	Pending OrderStatus = iota //iota is a special identifier that is reset to 0 whenever the keyword const appears in the source and increments by 1 for each item in the const block
	Shipped
	Delivered
	Canceled
)

func changeStatus(status OrderStatus) {
	fmt.Println("Order status changed to", status)
}

func main(){
	changeStatus(Canceled) // here you will get suggestion for enums
}