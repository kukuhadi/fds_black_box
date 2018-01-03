package main

func initializeRoutes(){
	router.POST("/stream", MessageStream)
	router.POST("/tes", TesStream)
	router.POST("/tes2", CreatePerson)
}