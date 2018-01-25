package main

import (
	"time"
	"strconv"
	"math"
	"github.com/keltia/leftpad"
	//"fmt"
)

type Parsechecktransaction struct{
	productID_parse, seconData_parse, branchCode_parse string
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

func (t Parsechecktransaction) isCheckTransaction() bool {
	// Bikin Role filternya disini
	if (t.productID_parse == "ATM" || t.seconData_parse == "ATM") && (t.branchCode_parse == "960" || t.branchCode_parse == "7168" || t.branchCode_parse == "7169" || t.branchCode_parse == "8192"){
		return true
	} else if (t.productID_parse == "SHA" || t.seconData_parse == "SHA") && (t.branchCode_parse == "960" || t.branchCode_parse == "7168" || t.branchCode_parse == "8192"){
		return true
	} else if (t.productID_parse == "HMB" || t.seconData_parse == "HMB") && (t.branchCode_parse == "960" || t.branchCode_parse == "7168" || t.branchCode_parse == "8192"){
		return true
	} else if (t.productID_parse == "BCA" || t.seconData_parse == "BCA") && (t.branchCode_parse == "960" || t.branchCode_parse == "7168" || t.branchCode_parse == "8192"){
		return true
	} else if (t.productID_parse == "MC" || t.seconData_parse == "MC") && (t.branchCode_parse == "960" || t.branchCode_parse == "7168" || t.branchCode_parse == "8192"){
		return true
	} else if (t.productID_parse == "MEPS" || t.seconData_parse == "MEPS") && (t.branchCode_parse == "960" || t.branchCode_parse == "7168" || t.branchCode_parse == "8192"){
		return true
	} else if (t.productID_parse == "CI" || t.seconData_parse == "CI") && (t.branchCode_parse == "960" || t.branchCode_parse == "7168" || t.branchCode_parse == "8192"){
		return true
	} else if (t.productID_parse == "MS" || t.seconData_parse == "MS") && (t.branchCode_parse == "960" || t.branchCode_parse == "7168" || t.branchCode_parse == "8192"){
		return true
	} else if (t.productID_parse == "EDCBRI" || t.seconData_parse == "EDCBRI") && (t.branchCode_parse == "960" || t.branchCode_parse == "7168" || t.branchCode_parse == "8192" || t.branchCode_parse == "7169"){
		return true
	} else {
		return false
	}
}

func secondsToTimeStamp(input int) string {
    seconds := input % (60 * 60 * 24)
    hours := math.Floor(float64(seconds) / 60 / 60)
    seconds = input % (60 * 60)
    minutes := math.Floor(float64(seconds) / 60)
    seconds = input % 60

    shours, _ :=  leftpad.PadChar(strconv.Itoa(int(hours)),2,'0')
    sMinutes, _ :=  leftpad.PadChar(strconv.Itoa(int(minutes)),2,'0')
    sSeconds, _ :=  leftpad.PadChar(strconv.Itoa(int(seconds)),2,'0')

    result := shours +":" + sMinutes + ":" + sSeconds

    return result
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
	var previousTransaction Previoustransaction
	var returnCalculateScore Returncalculatescore
	var parseScore Parsecalculatescore
	var parserSetRedisData Parsesetredisdata
  var tempSequence Tempsequence

	i.Nokartu = j.Cardno_prosw
	i.Amounttransaction = j.Txamount_prosw
	i.Terminaltype	= j.Terminaltype_prosw
	i.Terminalid = j.Termid_prosw
	x := j.Termid_prosw
	i.Terminalinformation = j.Narrative_prosw
	i.Statustransaction = j.Responcode_prosw
	i.Transdate = j.Transdate_prosw
  iTranstimeConvert, _ := strconv.Atoi(j.Transtime_prosw)
	i.Transtime = secondsToTimeStamp(iTranstimeConvert)
	i.Descriptiontransaction = j.Narrative_prosw

	locationAnalyst := true
	//sequenceAnalyst := false
	prevTransactionData := true

	// Lookup to Database for ATM BRI Location
	if i.Terminaltype == "ATM_BRI" || i.Terminaltype == "ATM MERAH PUTIH" {

		if i.Terminaltype == "ATM_BRI" {
			sX,_ := leftpad.PadChar(x, 7, '0')
			x = "BRI"+sX
		}
		locationString = findLocation(x)

		if locationString.statusInformation != "FOUND" {
			i.Longitude = locationString.longitudeTerminal
			i.Latitude = locationString.latitudeTerminal
		} else {
			locationAnalyst = false
		}
	}

	if locationAnalyst == false {
		locationString = callGoogleMaps(i.Terminalinformation)
		if locationString.longitudeTerminal != "" {
			i.Longitude = locationString.longitudeTerminal
			i.Latitude = locationString.latitudeTerminal
		} else {
			locationAnalyst = false
		}
	}

	previousTransaction = getRedisData(i.Nokartu)

	if previousTransaction.statusInformation != "response is nil" {
		i.Prevlatitude = previousTransaction.latitudeTerminal
		i.Prevlongitude = previousTransaction.longitudeTerminal
		i.Prevtransdate = previousTransaction.dateTransaction
		i.Prevtranstime = previousTransaction.timeTransaction
		i.Prevstatus = previousTransaction.statusTransaction
	} else
	{
		prevTransactionData = false
	}

		// ** START SEQUENCE ANALYST ** //
		// Get List of Transaction on DB.
		// Loop data lakukan perbandingan antara waktu.
		// Cek kode mata uang, IDR = 360.
		// Ada pembedaan antara perbandingan list transaksi IDR dengan non IDR.
		// Data yang dibandingkan adalah Data transaksi h-48 jam dari sebelumnya.
		resultsSequence := findListSequence(i.Nokartu)
		n := 0
		score := 0
		scorePhase1 := 0
		recordIDTemp := ""
		//lastTransactionSeqDB_waktuTemp := ""
		//lastTransactionSeqDB_responseTemp := ""
		constFormatDateTime := "2006-01-02 15:04"
		DateTimeString := ""
		PastDateTimeString := ""

		for resultsSequence.Next() {
			var tempSequence Tempsequence

			err := resultsSequence.Scan(&tempSequence.kartuTemp, &tempSequence.responseTemp, &tempSequence.matauangTemp, &tempSequence.waktuTemp, &tempSequence.scoringTemp, &tempSequence.idRecordTemp)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}

			// Selisih waktu
			if n == 0 { // Melakukan perbandingan antara transaksi actual dengan history terakhir.
				recordIDTemp = tempSequence.idRecordTemp
				DateTimeString = i.Transdate + " " + i.Transtime
				PastDateTimeString = tempSequence.waktuTemp

				parserStartDate, errParse1 := time.Parse(constFormatDateTime, DateTimeString)
				parserEndDate, errParse2 := time.Parse(constFormatDateTime, PastDateTimeString)

				if errParse1 != nil {
          i.Remark = "Error : "+errParse1.Error()
          return i
        }
				if errParse2 != nil {
          i.Remark = "Error : "+errParse2.Error()
          return i
        }

		    delta := parserStartDate.Sub(parserEndDate)
		    SelisihWaktu := delta.Minutes() // Selisih dalam satuan jam

				if SelisihWaktu < 60 { // Menggunakan satuan menit melakukan perbandingan kode status transaksi.
					if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
							score += 0
					} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
							score += 80
					} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
							score += 90
					} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
							score += 100
					}
				} else if SelisihWaktu > 60 && SelisihWaktu < 300 {
					if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
							score += 0
					} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
							score += 60
					} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
							score += 70
					} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
							score += 80
					}
				} else if SelisihWaktu > 300 && SelisihWaktu < 720 {
					if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
							score += 0
					} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
							score += 40
					} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
							score += 50
					} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
							score += 60
					}
				} else if SelisihWaktu > 720 && SelisihWaktu < 1440 {
					if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
							score += 0
					} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
							score += 20
					} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
							score += 30
					} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
							score += 40
					}
				} else {
					if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
							score += 0
					} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
							score += 10
					} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
							score += 20
					} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
							score += 30
					}
				}

				if SelisihWaktu < 120 { // Menggunakan satuan menit melakukan perbandingan  mata uang
					if j.Ccycode_prosw == "360" && tempSequence.matauangTemp == "360" {
						score += 0 // Justifikasinya adalah karena BRI bank di Indonesia, dan kemungkinan terbesar menggunakan IDR dalam bertransaksi.
					} else if j.Ccycode_prosw != "360" && tempSequence.matauangTemp == "360" {
						score += 1000 // Justifikasinya adalah engga mungkin dalam waktu kurang dari 120 menit berubah mata uang transaksinya.
					} else if j.Ccycode_prosw == "360" && tempSequence.matauangTemp != "360" {
						score += 750
					}
				} else {
					if j.Ccycode_prosw == "360" && tempSequence.matauangTemp == "360" {
						score += 0 // Justifikasinya adalah karena BRI bank di Indonesia, dan kemungkinan terbesar menggunakan IDR dalam bertransaksi.
					} else if j.Ccycode_prosw != "360" && tempSequence.matauangTemp == "360" {
						score += 300 // Justifikasinya adalah engga mungkin dalam waktu kurang dari 120 menit berubah mata uang transaksinya.
					} else if j.Ccycode_prosw == "360" && tempSequence.matauangTemp != "360" {
						score += 400
					}
				}
				scorePhase1 = score
			} else { // Melakukan penjumlahan scoring dari 4 transaksi sebelumnya.
        intScoringTemp, _ := strconv.Atoi(tempSequence.scoringTemp)
				score += intScoringTemp
				recordIDTemp = recordIDTemp + "," + tempSequence.idRecordTemp
			}
			n += 1
		}

		// menghapus record diluar 5 aktivitas terakhir dari masing-masing kartu
		deleteOldSequence(recordIDTemp, i.Nokartu)

		// menambahkan record baru
		tempSequence.kartuTemp = i.Nokartu
		tempSequence.responseTemp = i.Statustransaction
		tempSequence.matauangTemp = j.Ccycode_prosw
		tempSequence.scoringTemp = strconv.Itoa(scorePhase1)
		tempSequence.waktuTemp = DateTimeString

		tempSequence.setDataSequence()

		// Convert Scoring Fraud base on Sequence.
		if score < 200 {
			i.Labelfraud_sequence = "Hijau"
		} else if score < 1500 && score > 200 {
    	i.Labelfraud_sequence = "Kuning"
    } else {
    	i.Labelfraud_sequence = "Merah"
    }

		i.Scorefraud_sequence = strconv.Itoa(score)

		// ** END SEQUENCE ANALYST ** //
    tambahanRemark := ""
	if locationAnalyst {
			parseScore.currentLatitude = i.Latitude
			parseScore.currentLongitude = i.Longitude
			parseScore.pastLatitude = i.Prevlatitude
			parseScore.pastLongitude = i.Prevlongitude
			parseScore.currentDate = i.Transdate
			parseScore.currentTime = i.Transtime
			parseScore.pastDate = i.Prevtransdate
			parseScore.pastTime = i.Prevtranstime
			parseScore.currentStatus = i.Statustransaction
			parseScore.pastStatus = i.Prevstatus

			if prevTransactionData {
				returnCalculateScore = parseScore.calculateScore()
				i.Labelfraud_location = returnCalculateScore.labelFraud
				i.Scorefraud_location = returnCalculateScore.scoreFraud
				i.Remark = returnCalculateScore.errorCalculate
			} else {
				i.Labelfraud_location = "Hijau"
				i.Scorefraud_location = "0"
				i.Remark = "No Previous Data"
			}

			// Set Data to REDIS
			parserSetRedisData.latitudeTerminal = i.Latitude
			parserSetRedisData.longitudeTerminal = i.Longitude
			parserSetRedisData.dateTransaction = i.Transdate
			parserSetRedisData.timeTransaction = i.Transtime
			parserSetRedisData.statusTransaction = i.Statustransaction
			parserSetRedisData.nomerKartu = i.Nokartu
			tambahanRemark = parserSetRedisData.setRedisData()
	}

	// Set Data untuk sequenceAnalyst.

	i.Remark = i.Remark + ";" + tambahanRemark
	i.Remark = previousTransaction.statusInformation
	i.Status = prevTransactionData
	return i

}
