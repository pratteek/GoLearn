package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//struct tags - specify what field names should be when serialized into JSON
type album struct{
	ID 	   string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
var albums = []album{
	{ID:"A1",Title:"Summer",Artist:"Harry",Price:79.99},
	{ID:"A2",Title:"Winter",Artist:"Zayn",Price:199.99},
	{ID:"A3",Title:"Spring",Artist:"Niall",Price:49.99},
	{ID:"A4",Title:"Monsoon",Artist:"Liam",Price:39.99},
}
//gin.Context is the most important part of Gin. 
//It carries request details, validates and serializes JSON, and more.
func getAlbums(c *gin.Context){
	
	//the below func creates JSON from slice of album structs
	//It serialize the struct into JSON and add it to the response.
	c.IndentedJSON(http.StatusOK,albums)
}
func getAlbumByID(c *gin.Context){
	id := c.Param("id")

	//Look for the id in albums
	for _,album := range albums{
		if album.ID == id{
			c.IndentedJSON(http.StatusOK,album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"album not found"})
}
func createAlbum(c *gin.Context){
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil{
		return
	}

	albums = append(albums,newAlbum)
	c.IndentedJSON(http.StatusCreated,albums)
}

func main(){
	//create router
	router := gin.Default()

	//add API endpoint /albums
	router.GET("/albums",getAlbums)
	router.GET("/albums/:id",getAlbumByID)
	router.POST("/albums",createAlbum)

	router.Run("localhost:5500")
}