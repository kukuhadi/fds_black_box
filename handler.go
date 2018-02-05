package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
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
          messageText := message.Nokartu + "|" + message.Amounttransaction + "|" + message.Terminaltype + "|" + message.Terminalid + "|" + message.Terminalinformation + "|" + message.Statustransaction + "|" + message.Transdate + "|" + message.Transtime + "|" + message.Descriptiontransaction + "|" + message.Longitude
          messageText = messageText + "|" + message.Latitude + "|" + message.Prevterminalid + "|" + message.Prevlongitude + "|" + message.Prevlatitude + "|" + message.Prevtransdate + "|" + message.Prevtranstime + "|" + message.Prevstatus + "|" + message.Labelfraud_location + "|" + message.Scorefraud_location
          messageText = messageText + "|" + message.Labelfraud_sequence + "|" + message.Scorefraud_sequence + "|" + message.Remark + "|" + message.Kodematauang + "\r\n"
          writeToFile(messageText)
          //ProduceKafka(message)
        }
        c.JSON(http.StatusOK, gin.H{"Status Scoring Service": http.StatusOK, /*"data": json, */"Data Scoring": message})
    } else {
    	c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
    }
}

func writeToFile(message string){
  filename := "/Users/apple/work/programming/github.com/Kukuh/project/fds_black_box/file_result/fileresult.csv"
  f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
  if err != nil {
      panic(err)
  }

  defer f.Close()

  if _, err = f.WriteString(message); err != nil {
      panic(err)
  }
}
