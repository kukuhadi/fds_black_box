package main

import "github.com/gin-gonic/gin"

var router *gin.Engine

func main(){
	// Definisikan router dengan menggunakan Gin Framework
	router = gin.Default()

	// Load Configurasi route 
	initializeRoutes()

	//jalankan web service dengan menggunakan port yang sudah ditentukan
	router.Run("127.0.0.1:8001")
}
