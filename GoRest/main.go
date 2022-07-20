package main

import (
	"fmt"
)

func main(){
	//struct tags - specify what field names should be when serialized into JSON
	type album struct{
		ID 	   string  `json:"id"`
		Title  string  `json:"title"`
		Artist string  `json:"artist"`
		Price  float64 `json:"price"`
	} 

	albums := []album{
		{ID:"A1",Title:"Summer",Artist:"Harry",Price:79.99},
		{ID:"A2",Title:"Winter",Artist:"Zayn",Price:199.99},
		{ID:"A3",Title:"Spring",Artist:"Niall",Price:49.99},
		{ID:"A4",Title:"Monsoon",Artist:"Liam",Price:39.99},
	}
	fmt.Println(albums)
}