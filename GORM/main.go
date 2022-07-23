package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type student struct{
	//convert struct to GORM
	gorm.Model
	Name string `json:"name"`
	Age uint	`json:"age"`
	Branch string `json:"branch"`
	GPA float64 `json:"gpa"`
}

const DSN = "root:admin@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

//intialise the db
func IntitialMigration(){
	//db details
	DB, err := gorm.Open(mysql.Open(DSN),&gorm.Config{})
	if err!= nil{
		fmt.Println(err.Error())
	}
	DB.AutoMigrate(&student{})


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
func PostStudents(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var newStudent student
	if err := json.NewDecoder(r.Body).Decode(&newStudent);err!=nil{
		http.Error(w,"Error decoding response object",http.StatusBadRequest)
		return
	}
	students = append(students,newStudent)

	response, err := json.Marshal(&students)
	if err!=nil{
		http.Error(w,"Error decoding response object",http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
func main(){
	r := mux.NewRouter()
	r.Path("/").Methods(http.MethodGet).HandlerFunc(Home)
	r.Path("/students").Methods(http.MethodGet).HandlerFunc(GetStudents)
	r.Path("/students").Methods(http.MethodPost).HandlerFunc(PostStudents)
	log.Fatal(http.ListenAndServe(":9000",r))

	IntitialMigration()
}
