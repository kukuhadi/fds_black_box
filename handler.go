package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

/*func MessageStream(c *gin.Context){
	var json Proswmin
	var message Completemessage

	if err := c.ShouldBindJSON(&json); err == nil {
        // Get All Information from previous transaction and location from each terminal.
        message = json.CompleteInformation()

        ProduceKafka(message)

        c.JSON(http.StatusOK, gin.H{"status hello": http.StatusOK, "data": json, "data 2": message})
    } else {
    	c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
    }
}*/

func MessageStreamComplete(c *gin.Context){
	var json Prosw
	var message Completemessage
  var parseCheckTransaction Parsechecktransaction

	if err := c.ShouldBindJSON(&json); err == nil {
        // Get All Information from previous transaction and location from each terminal.
        parseCheckTransaction.productID_parse = json.Productid_prosw
        parseCheckTransaction.seconData_parse = json.Seconddata1_prosw
        parseCheckTransaction.branchCode_parse = json.Brnchcode_prosw

        status := parseCheckTransaction.isCheckTransaction()

        if status {
          message = json.CompleteInformation()
          ProduceKafka(message)
        }



        c.JSON(http.StatusOK, gin.H{"status hello": http.StatusOK, "data": json, "data 2": message})
    } else {
    	c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
    }
}
