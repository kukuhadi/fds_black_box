package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    // "fmt"
    //e "github.com/eirka/eirka-libs/errors"
    // "strconv"
)

type Teskonten struct {
	Konten	string	`json:"konten"`
}

type Person struct {
    ID        string   `json:"id"`
    Firstname string   `json:"firstname"`
    Lastname  string   `json:"lastname"`
    Address   *Address `json:"address"` //Jika di dalam variable memiliki detail variable, bisa di deklarasikan detail kemudian. 
}

type Address struct {
    City  string `json:"city"`
    State string `json:"state"`
}

func MessageStream(c *gin.Context){
	var json Prosw
	var message Completemessage
	
	if err := c.ShouldBindJSON(&json); err == nil {
        // Get All Information from previous transaction and location from each terminal.
        message = json.CompleteInformation()
        // fmt.Println(jsonMessage.noKartu)
        //message.Nokartu = json.Nokartu
        c.JSON(http.StatusOK, gin.H{"status hello": http.StatusOK, "data": json, "data 2": message})
    } else {
    	c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
    }
    
    /*
    c.Bind(json)
    fmt.Println(json)
    c.JSON(http.StatusOK, c)
    */
}

func TesStream(c *gin.Context){
	var json Teskonten
	var json2 Teskonten
	if err := c.ShouldBindJSON(&json); err == nil {
		json2.Konten = json.Konten
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": json2})
	}else {
    	c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
    }
}

func CreatePerson(c *gin.Context) {
    var person Person
    var json Person
    if err := c.ShouldBindJSON(&json); err == nil {
        person.ID = json.ID
        person.Firstname = json.Firstname
        person.Lastname = json.Lastname
    } else {
        // param_id = strconv.Atoi(c.PostForm("id"))
        param_id := c.Param("id")
        param_firstname := c.DefaultQuery("firstname", "Guest")
        // param_lastname := c.DefaultQuery("lastname", "")
        param_lastname := c.Query("lastname")
        person.ID = param_id
        person.Firstname = param_firstname
        person.Lastname = param_lastname
    }
    //people = append(people, person)
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": person})
}