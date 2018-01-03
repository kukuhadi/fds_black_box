package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func connect() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/atm_location")
    if err != nil {
        return nil, err
    }

    return db, nil
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
        QueryRow("select langtitude, longitude from tb_atm_location where terminalid = ?", terminalID).
        Scan(&result.latitudeTerminal, &result.longitudeTerminal)
        result.statusInformation = "good"
    if err != nil {
        fmt.Println(err.Error())
        return result
    }

    return result
}