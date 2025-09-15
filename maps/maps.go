package main

import (
	"fmt"
	
)

func main(){
	
	//maps are like dictionary in python, object in javascript
	//maps are like struct but struct have fixed fields but maps can have dynamic fields
	//maps are like hash table in other programming languages
	//maps are like json object in javascript
	//maps are unordered collection of key value pairs that does not allow duplicates.
	//maps are mutable
	//maps are reference type

	//syntax to declare a map

	m:=make(map[string]int) //map[keyType]valueType

	//setting values in map
	m["name"]=12
	m["age"]=21
	m["height"]=5

	//getting values from map
	fmt.Println(m)
	fmt.Println(m["name"])
	fmt.Println(m["age"])
	fmt.Println(m["height"])

	//deleting a key from map
	delete(m,"height")
	fmt.Println(m)

	//IMPORTANT:: if key is not present in map then it will return the zero value of the value type

	fmt.Println(m["phone"]) //0
	fmt.Println(len(m)) //0

	clear(m) //clear the map
	fmt.Println(m)


	mp:=map[string]int{"price":40,"phones":30}
//ok is a boolean value which tells whether the key is present in map or not and val is the value of the key
	val, ok:=mp["price"] 
	if ok{
		fmt.Println("price is present in map and value is ",val)
	}else{
		fmt.Println("price is not present in map")
	}
	

	// m1:=map[string]int{"a":1,"b":2,"c":3}
	// m2:=map[string]int{"a":1,"b":2,"c":3}
	// fmt.Println(maps.Equal(m1,m2)) // true
}	