package main

import (
	"fmt"
  "time"
)

func hello(){
  time.Sleep(3 * time.Second)
  fmt.Println("Hello Guys! Welcome back to my channel")
}

func main() {
  //Input and Output
	var username string
	var password string
	fmt.Println("Enter Username:")
	fmt.Scan(&username)
	fmt.Println("Enter Password:")
	fmt.Scan(&password)
	fmt.Printf("Username: %v and Password: %v\n",username,password)
	//Arrays and Slices
	
	var names = []string{"Prateek","Nidhi","paks"}
	fmt.Println(names)

	var namesNew [2]string
	namesNew[0]="prateek"
	namesNew[1]="Robo"
	fmt.Println(namesNew)

	//What if we don't know the size of the array?
	//Then slice come into picture
	//slice -> dynamic array

	var bookings []string

	bookings = append(bookings,"Prateek")
	bookings = append(bookings,"Nidhi")
	bookings = append(bookings,"Paksboi")
	bookings = append(bookings,"Paro")

	fmt.Println(bookings)

	//Loops
	//Only for loop is available in GO
	i := 0
	for i<10{
		i=i+1
		fmt.Println(i)
	}
	//we can use range to iterate the slice
	for index, booking := range bookings{
		fmt.Print(index," ")
		fmt.Println(booking)
	}

  //Conditional statements
  //if else
  abs := 1
  if abs == 0{
    fmt.Println("Go to Gym")
  }else{
    fmt.Println("Keep Going to Gym")
  }

  //Go programs are organized in packages
  //A package is a collection of Go files
  //package 1, package 2 and main package
  SayHello()

  //Points to Note
  //To use functions from other packages we need to import
  //To export a function we need to make the first letter capital
  

  //Maps
  //key-value pair
  userData := make(map[string]string) //defines the map
  userData["Name"]="Prateek"
  userData["Age"]="22"
  userData["DOB"]="20 Sep 1999"
  fmt.Println(userData)

  //struct 
  //data with different dataypes
  type student struct{
    Name string
    Age uint
    email string
  }

  studentData := make([]student,2,2)
  
  studentData[0] = student{
    Name: "Prateek",
    Age: 22,
    email: "prateek@google.com",
  }
  studentData[1] = student{
    Name: "Paks",
    Age: 23,
    email: "paks@g.com",
  }
  
  fmt.Println(studentData)
  
  // using struct student as a datatype
  studentDataNew := student{
    Name: "NewBoi",
    Age: 0,
    email: "newboi@gmail.com",
  }
  fmt.Println(studentDataNew)

  //time packege
  //time.sleep(10 * time.second)

  //Go routines - Concurrency in Go
  //Using concurrency, when a task takes time - in go it runs it in a new thread
  go hello()
  fmt.Println("Glad to have you back!")
}
