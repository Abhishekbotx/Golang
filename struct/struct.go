package main

import (
	"fmt"
	// "go/types"
	"time"
)

// struct is a collection of fields
// struct is a user defined type
// struct is a composite type
// struct is a value type
// struct is used to group related data together
// struct is used to create complex data types

//struct is a blueprint and concrete instance of struct is called struct variable or struct literal
//eg: Order struct is a blueprint and order1, order2 are struct variables or struct literals

//using struct we can we can do object oriented programming in go

//syntax to declare a struct
type Customer struct{
	name string
	age int
}

type Order struct{
	id string
	amount float32
	status string
	createdAt time.Time //nanosecond precision
	newCustomer Customer //embedding another struct, this is called anonymous field
}


//reciever type //IMP:🌟🌟
//🍁Case 1:
// func (o Order) changeStatus(status string){// we have to use pointer receiver to change the value of struct field
// 	// o is a copy of the struct
// 	// so changing the value of o will not change the value of the original struct
// 	o.status=status
// }

//🍁Case 2: using pointer receiver
func (o *Order) changeStatus(status string){ //syntax: func (receiverName *ReceiverType) methodName(params) returnType{}
	o.status=status //why we didnt use *o.status because go automatically dereference the pointer when we access the field of the struct
}

//constructor function for struct
func newOrder(id string, amount float32, status string) *Order{ //returning pointer to struct
	myOrder:=Order{
		id:id,
		amount:amount,
		status: status,
	}

	return &myOrder //returning the address of the struct
}	

func main(){
	var order1 Order=Order{
		id: "12345",
		amount: 100.50,
		status: "pending",
		createdAt: time.Now(),
	}

	order2:=Order{
		id:"54321",
		amount:200.75,
		createdAt: time.Now(),
	}

	embedOrder:=Order{
		id:"98765",
		amount:300.25,
		status:"processing",
		createdAt: time.Now(),
		newCustomer: Customer{ //setting value of embedded struct
			name: "John Doe",
			age: 30,
		},
	}
	fmt.Println("customer::",embedOrder.newCustomer)
	fmt.Println(order2)
	//🌟Note:
	//if we dont set the value of a field then it will take the zero value of the field type
	//eg: for string it will be "" for int it will be 0 for float it will be 0.0 for bool it will be false for time.Time it will be 0001-01-01 00:00:00 +0000 UTC

	fmt.Println(order1.amount)

	order1.changeStatus("shipped")
	// 🍁.Case 1// fmt.Println(order1.status) // still pending because struct is value type, so a copy of the struct is passed to the method.
	fmt.Println(order1.status) // shipped because we used pointer receiver

	myOrder:=newOrder("67890",300.25,"delivered")
	fmt.Println("using constructor",myOrder) // &{67890 300.25 delivered 2023-10-14 12:34:56.789012345 +0000 UTC}


	language:=struct{
		name string
		isGood bool
	}{
		name: "Go",
		isGood: true,
	}

	fmt.Println(language) // struct { name string; isGood bool }


	

}