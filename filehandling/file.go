package main

import (
	// "bufio"
	// "fmt"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")

	// if err != nil { //mean erro is there
	// 	// loggin the error
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil { //mean erro is there
	// 	// loggin the error
	// 	panic(err)
	// }

	// fmt.Println("filename", fileInfo.Name())
	// fmt.Println("file size", fileInfo.Size())
	// fmt.Println("file lopermisssiong", fileInfo.Mode()) //to check file permission
	// fmt.Println("file modified at", fileInfo.ModTime()) //to check file permission

	//🍁read file
	//Method 1:
	// f,err:=os.Open("example.txt")
	// if err !=nil{
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()

	// defer f.Close() // when main funcn will be exited this will be closed

	// buf:=make([]byte , fileInfo.Size())

	// d,err:=f.Read(buf) // Read return length

	// if err !=nil{
	// 	panic(err)//eof file error
	// }

	// for i:=0; i<len(buf);i++{

	// 	// fmt.Println("data",d,buf)
	// 	fmt.Println("data",d,string(buf[i]))
	// }

	// Method 2:

	// data,err:=os.ReadFile("example.txt")

	// if err !=nil{
	// 	panic(err)
	// }

	// fmt.Println("data here:",string(data))

	// 🍁Read folders

	// dir,err:=os.Open("../")

	// if err !=nil{
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo,err:=dir.ReadDir(-1)// number will decide how much file you want to show // if you give <0 it will return all

	// for _,fi:= range fileInfo{
	// 	fmt.Println("dir:",fi.Name(),fi.IsDir())
	// }

	// 🍁Create file and write into it
	// f,err:=os.Create("example2.txt")

	// if err !=nil{
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("baby tu jaa, haa bhai bande leke pahuch ")
	// f.WriteString("hi golang prithvi pe apka swagat hai ")
	// f.WriteString("hmare mummy papa ko pta hai jo krna hai kr lo")

	// bytes :=[]byte("hello Doland Trump")//Another way of putting data
	// f.Write(bytes)

	// fmt.Print("heyy")

	//🍁readd and write to another file (streaming fashion

	// sourceFile, err := os.Open(("example2.txt"))

	// defer sourceFile.Close()

	// destFile, err := os.Create("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile) // reader is source file
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(err)

	// 	}
	// }

	// writer.Flush()

	// fmt.Println("return to new line successfully")

	//🌟Check for coffee function copy all directly without buffer

	//delete a file

	// sourceFile, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	err:=os.Remove("example.txt")

	if err !=nil{
		panic(err)
	}

	fmt.Println("file deleted successfully")

}
