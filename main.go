package main

import (
	"fmt"
)

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
}
