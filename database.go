package main

import "fmt"
import "os"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "github.com/mediocregopher/radix.v2/redis"
import "gopkg.in/resty.v1"

func connect() (*sql.DB, error) {
    connectionString := os.Getenv("usernameDBMysql") + ":" + os.Getenv("passwordDBMysql") + "@tcp("+os.Getenv("hostDatabase")+")/"+os.Getenv("schemaDatabase")
    db, err := sql.Open("mysql", connectionString)
    if err != nil {
        return nil, err
        fmt.Println(err.Error())
    }

    return db, nil
}

func connectSequenceData() (*sql.DB, error) {
    connectionString := os.Getenv("usernameDBMysql") + ":" + os.Getenv("passwordDBMysql") + "@tcp("+os.Getenv("hostDatabase")+")/"+os.Getenv("schemaDatabase_Seq")
    db, err := sql.Open("mysql", connectionString)
    if err != nil {
        return nil, err
        fmt.Println(err.Error())
    }

    return db, nil
}

func getRedisData(noKartu string) Previoustransaction {
  var result = Previoustransaction{}
  result.latitudeTerminal = ""
  result.longitudeTerminal = ""
  result.statusTransaction = ""
  result.dateTransaction = ""
  result.timeTransaction = ""
  result.statusInformation = "Success"
  conn, err := redis.Dial("tcp", os.Getenv("redisServer"))
    if err != nil {
        result.statusInformation = err.Error()+noKartu
        return result
    }

    defer conn.Close()

    status, err := conn.Cmd("HGET", "nokartu:"+noKartu, "status").Str()
    if err != nil {
        result.statusInformation = err.Error()+noKartu+"status"
        return result
    }

    longtitude, err := conn.Cmd("HGET", "nokartu:"+noKartu, "longtitude").Str()
    if err != nil {
        result.statusInformation = err.Error()+noKartu
        return result
    }

    latitude, err := conn.Cmd("HGET", "nokartu:"+noKartu, "latitude").Str()
    if err != nil {
        result.statusInformation = err.Error()+noKartu
        return result
    }

    timeTransaction, err := conn.Cmd("HGET", "nokartu:"+noKartu, "timetransaction").Str()
    if err != nil {
        result.statusInformation = err.Error()
        return result
    }

    dateTransaction, err := conn.Cmd("HGET", "nokartu:"+noKartu, "datetransaction").Str()
    if err != nil {
        result.statusInformation = err.Error()
        return result
    }

    result.latitudeTerminal = latitude
    result.longitudeTerminal = longtitude
    result.statusTransaction = status
    result.dateTransaction = dateTransaction
    result.timeTransaction = timeTransaction

    return result
  }

  func (x Parsesetredisdata) setRedisData() string {
    var hasil string
    // hasil = "Coba Masukin ke Redis "+x.nomerKartu
    conn, err := redis.Dial("tcp", os.Getenv("redisServer"))
      if err != nil {
          hasil = "Failed to Connect " + err.Error()
      }

      defer conn.Close()

      resp := conn.Cmd("HMSET", "nokartu:"+x.nomerKartu, "latitude", x.latitudeTerminal, "longtitude", x.longitudeTerminal, "status", x.statusTransaction, "datetransaction", x.dateTransaction, "timetransaction", x.timeTransaction)
      // Check the Err field of the *Resp object for any errors.
      if resp.Err != nil {
          hasil = "Failed to Set Data " + resp.Err.Error()
      }
      return hasil
    }

func findLocation(terminalID string) Locationterminal {
	var result = Locationterminal{}
    var db, err = connect()
    if err != nil {
        fmt.Println(err.Error())
        return result
    }
    defer db.Close()

    err = db.
        QueryRow("select latitude, longitude from Tbl_atmlocation where id_atm = '" + terminalID + "' or id_atm like '%"+terminalID+"%'").
        Scan(&result.latitudeTerminal, &result.longitudeTerminal)
        result.statusInformation = "FOUND"
        fmt.Println(terminalID)
    if err != nil {
        fmt.Println(err.Error())
        return result
    }

    return result
}

func (sequenceData Tempsequence) setDataSequence() string {
    hasil := "Success"
    var db, err = connectSequenceData()
    if err != nil {
        fmt.Println(err.Error())
        hasil = "Failed : " + err.Error()
        return hasil
    }
    defer db.Close()

    // perform a db.Query insert
    insert, err := db.Query("INSERT INTO Tbl_sequence (kartu, responcode, ccycode, waktu, score) VALUES ( '"+sequenceData.kartuTemp+"', '"+sequenceData.responseTemp+"', '"+sequenceData.matauangTemp+"', '"+sequenceData.waktuTemp+"', '"+sequenceData.scoringTemp+"' )")

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
        hasil = "Failed : " + err.Error()
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()
    return hasil
}

func findListSequence(nomerKartu string) *sql.Rows {
    //hasil := "Success"
    var db, err = connectSequenceData()
    if err != nil {
        fmt.Println(err.Error())
        //hasil = "Failed : " + err.Error()
    }
    defer db.Close()

    // Execute the query
  	results, err := db.Query("SELECT kartu, responcode, ccycode, waktu, score, idrecord FROM Tbl_sequence where kartu = '"+nomerKartu+"' order by waktu desc limit 5")
  	if err != nil {
  		panic(err.Error()) // proper error handling instead of panic in your app
  	}

    return results
}

func deleteOldSequence(idRecord string, nomerKartu string) string {
  var db, err = connectSequenceData()
  if err != nil {
      fmt.Println(err.Error())
      return "FAILED"
  }
  defer db.Close()
  fmt.Println(idRecord)
  // Execute the query
  _, err = db.Query("DELETE from Tbl_sequence where idrecord not in ("+idRecord+") and kartu = '"+nomerKartu+"'")
  if err != nil {
    panic(err.Error()) // proper error handling instead of panic in your app
    return "FAILED"
  }

  return  "SUCCESS"
}

// Functions Declaration
func InitKafkaConsumer() {
    //resty.R().SetBody("{\"topic\":\"crs-unscored\",\"group\":\"test-group\"}").Post("http://35.186.144.202:8020/subscribe/topic/add")
    //resty.R().SetBody("{\"data\":[{\"topic\":\"crs-unscored\",\"url\":[\"http://0.0.0.0:8000/crs\"]}]}").Post("http://localhost:8020/subscribe/url/add")
}

func ProduceKafka(data interface{}) {
    resty.R().SetBody(data).Post("http://35.186.144.202:8020/publish/fds-result")
}
