package main

func initializeRoutes(){
	//router.POST("/fds", MessageStream)
	router.POST("/fdsComplete", MessageStreamComplete)
}
