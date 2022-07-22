package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type student struct{
	Name string `json:"name"`
	Age uint	`json:"age"`
	Branch string `json:"branch"`
	GPA float64 `json:"gpa"`
}
var students =  []student{
	{Name:"Prateek",Age:22,Branch:"CSE",GPA:10},
	{Name:"Rohit",Age:22,Branch:"ECE",GPA:9},
	{Name:"Nidhi",Age:22,Branch:"EEE",GPA:9.5},
}
func Home(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode("Home Page")
	fmt.Println("Home Page")
}
func GetStudents(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(students)
	fmt.Println("/students endpoint hit")
}
func main(){
	r := mux.NewRouter()
	r.Path("/").Methods(http.MethodGet).HandlerFunc(Home)
	r.Path("/students").Methods(http.MethodGet).HandlerFunc(GetStudents)
	log.Fatal(http.ListenAndServe(":9000",r))
}
