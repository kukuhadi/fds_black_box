package main

import (
	"time"
	"strconv"
	"math"
	//"fmt"
)

type Prosw struct{
	Nokartu					string		`json:"nokartu"`
	Amounttransaction		string		`json:"amounttransaction"`
	Terminaltype			string		`json:"terminaltype"`
	Terminalid				string		`json:"terminalid"`
	Terminalinformation		string		`json:"terminalinformation"`
	Statustransaction		string		`json:"statustransaction"`
	Transdate				string		`json:"transdate"`
	Transtime				string		`json:"transtime"`
	Descriptiontransaction	string 		`json:"descriptiontransaction"`
}

type Completemessage struct{
	Nokartu					string		`json:"nokartu"`
	Amounttransaction		string		`json:"amounttransaction"`
	Terminaltype			string		`json:"terminaltype"`
	Terminalid				string		`json:"terminalid"`
	Terminalinformation		string		`json:"terminalinformation"`
	Statustransaction		string		`json:"statustransaction"`
	Transdate				string		`json:"transdate"`
	Transtime				string		`json:"transtime"`
	Descriptiontransaction	string 		`json:"descriptiontransaction"`
	Longitude				string		`json:"longitude"`
	Latitude				string		`json:"latitude"`
	Prevterminalid			string 		`json:"prevterminalid"`
	Prevlongitude			string		`json:"prevlongitude"`
	Prevlatitude			string		`json:"prevlatitude"`
	Prevtransdate			string 		`json:"prevtransdate"`
	Prevtranstime			string 		`json:"prevtranstime"`	
	Prevstatus				string 		`json:"prevstatus"`
	Labelfraud				string 		`json:"labelfraud"`
	Scorefraud				string 		`json:"scorefraud"`
	Remark					string 		`json:"remark"`
}

type Locationterminal struct{
	latitudeTerminal, longitudeTerminal, statusInformation 	string	
}

type Parsecalculatescore struct{
	currentLatitude, currentLongitude, pastLatitude, pastLongitude, currentDate, currentTime, currentStatus, pastDate, pastTime, pastStatus string
}

type Returncalculatescore struct {
	labelFraud, scoreFraud, selisihWaktu, selisihJarak, errorCalculate string
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
  	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func (j Prosw) CompleteInformation() Completemessage { 
	var i Completemessage
	var locationString Locationterminal
	var returnCalculateScore Returncalculatescore
	var parseScore Parsecalculatescore

	i.Nokartu = j.Nokartu
	i.Amounttransaction = j.Amounttransaction
	i.Terminaltype	= j.Terminaltype
	i.Terminalid = j.Terminalid
	x := j.Terminalid
	i.Terminalinformation = j.Terminalinformation
	i.Statustransaction = j.Statustransaction
	i.Transdate = j.Transdate
	i.Transtime = j.Transtime
	i.Descriptiontransaction = j.Descriptiontransaction

	// return i 
	
	local := true

	// Lookup to Database for ATM BRI Location
	if i.Terminaltype == "ATM BRI" {
		locationString = findLocation(x)

		if locationString.longitudeTerminal != "" {
			i.Longitude = locationString.longitudeTerminal
			i.Latitude = locationString.latitudeTerminal	
		} else {
			local = false
		}
	}

	if local == false{
		// Lookup to Google API
	}

	// Lookup to Redis for Previous Transaction
	// getRedisData("No Kartu")
	// Masih bingung cara ngerjainnya :D 
	// Format REDIS --> NoKartu sebagai Key, dengan Informasi Terminal ID, Langitude, Lotitude, Date & Time

	i.Prevlatitude = "-8.610403"
	i.Prevlongitude = "115.173001"
	i.Prevtransdate = "2017-12-01"
	i.Prevtranstime = "18:00"
	i.Prevstatus = "61 - Transaksi melebihi limit transaksi harian"

	parseScore.currentLatitude = i.Latitude
	parseScore.currentLongitude = i.Longitude
	parseScore.pastLatitude = i.Prevlatitude
	parseScore.pastLongitude = i.Prevlongitude
	parseScore.currentDate = i.Transdate
	parseScore.currentTime = i.Transtime
	parseScore.pastDate = i.Prevtransdate
	parseScore.pastTime = i.Prevtranstime
	parseScore.pastStatus = i.Prevstatus

	// Get Label Fraud
	// i.Labelfraud = parseScore.calculateScore()
	returnCalculateScore = parseScore.calculateScore()
	i.Labelfraud = returnCalculateScore.labelFraud
	i.Scorefraud = returnCalculateScore.scoreFraud
	i.Remark = returnCalculateScore.errorCalculate

	return i 
	
}

func (j Parsecalculatescore) calculateScore() Returncalculatescore {
	// Menghitung Score dengan menggunakan 
		// - Selisih Waktu
		// - Status Transaksi Sebelum dengan Saat ini. 
		// - Selisih Jarak antara terminal ID sebelum dengan sesaat ini. 
	indicator := "Kosong"
	var result = Returncalculatescore{}
	result.labelFraud = ""
	result.scoreFraud = ""
	result.selisihJarak = ""
	result.selisihWaktu = ""
	result.errorCalculate = "Success"
	// Selisih Waktu
	constFormatDateTime := "2006-01-02 15:04"
	var Score int
	Score = 0

	DateTimeString := j.currentDate + " " + j.currentTime
	PastDateTimeString := j.pastDate + " " + j.pastTime

	parserStartDate, errParse1 := time.Parse(constFormatDateTime, DateTimeString)
	parserEndDate, errParse2 := time.Parse(constFormatDateTime, PastDateTimeString)

	if errParse1 != nil {
        result.errorCalculate = "Gagal Konversi Tanggal Awal " + errParse1.Error()
        return result
    }

    if errParse2 != nil {
        result.errorCalculate = "Gagal Konversi Tanggal Akhir " + errParse2.Error()
        return result
    }

    delta := parserStartDate.Sub(parserEndDate)
    SelisihWaktu := delta.Hours() // Selisih dalam satuan jam
    
    // Selisih Jarak
    flNewLatitude,errflNewLatitude := strconv.ParseFloat(j.currentLatitude, 64)
    flNewLongitude,errflNewLongitude := strconv.ParseFloat(j.currentLongitude, 64)
    flOldLatitude,errflOldLatitude := strconv.ParseFloat(j.pastLatitude, 64)
    flOldLongitude,errflOldLongitude := strconv.ParseFloat(j.pastLongitude, 64)
    if errflOldLongitude != nil {
	  result.errorCalculate = "Gagal Konversi Longtitude Baru " + errflOldLongitude.Error()
      return result
	}

	if errflNewLatitude != nil {
		result.errorCalculate = "Gagal Konversi Latitude Baru " + errflNewLatitude.Error()
      	return result
	}

	if errflOldLatitude != nil {
		result.errorCalculate = "Gagal Konversi Latitude Lama " + errflOldLatitude.Error()
      	return result
	}

	if errflNewLongitude != nil {
		result.errorCalculate = "Gagal Konversi Longtitude Lama " + errflNewLongitude.Error()
      	return result
	}

    SelisihJarak := Distance(flNewLatitude, flNewLongitude, flOldLatitude, flOldLongitude)

    // Menentukan Score
    // Berdasarkan Status Transaksi Sebelumnya
    if j.pastStatus == "51 - Saldo tidak cukup" { // 51 - Saldo tidak cukup
    	Score += 40
    } else if j.pastStatus == "61 - Transaksi melebihi limit transaksi harian" { //"61 - Transaksi melebihi limit transaksi harian"
    	Score += 50
    } else if j.pastStatus == "00 - Sukses" { // 
    	Score += 10
    } else {
    	Score += 0
    }

    // Berdasarkan Status Transaksi Saat Ini
    if j.currentStatus == "51 - Saldo tidak cukup" { // 51 - Saldo tidak cukup
    	Score += 40
    } else if j.currentStatus == "61 - Transaksi melebihi limit transaksi harian" { //"61 - Transaksi melebihi limit transaksi harian"
    	Score += 50
    } else if j.currentStatus == "00 - Sukses" { // 
    	Score += 10
    } else {
    	Score += 0
    }

    // Berdasarkan kombinasi Selisih Waktu dan Jarak
    // Jika waktu kurang dari 1 jam, menggunakan acuan kendaraan 80-120 km per jam 
    // Jika waktu antara 1 - 5 jam, menggunakan acuan kendaraan 200 km per jam (pesawat dan kereta)
    if SelisihWaktu < 1 && SelisihJarak > 100 {
    	Score += 50
    } else if SelisihWaktu < 5 && SelisihWaktu > 1 && 1000 > SelisihJarak && SelisihJarak > 100 {
    	Score += 40
    } else if SelisihWaktu < 12 && SelisihWaktu > 5 && SelisihJarak > 1000 {
    	Score += 30
    } else if SelisihWaktu > 12 && SelisihJarak > 1000 {
    	Score += 20
    } else {
    	Score += 10
    }

    if Score > 120 {
    	indicator = "Merah"
    } else if Score < 120 && Score > 60 {
    	indicator = "Kuning"
    } else {
    	indicator = "Hijau"
    }

    result.labelFraud = indicator
	result.scoreFraud = strconv.FormatInt(int64(Score), 10)
	result.selisihJarak = strconv.FormatFloat(SelisihJarak, 'f', -1, 32)
	result.selisihWaktu = strconv.FormatFloat(SelisihWaktu, 'f', -1, 32)

    return result
}