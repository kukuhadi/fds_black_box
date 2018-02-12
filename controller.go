package main

import (
	"time"
	"strconv"
	"math"
	"github.com/keltia/leftpad"
	"os"
  "fmt"
  "strings"
)

type Parsechecktransaction struct{
	productID_parse, seconData_parse, branchCode_parse, proccode_parse string
}

func (t Parsechecktransaction) isCheckTransaction() bool {
	// Bikin Role filternya disini
	fmt.Println("Product-ID : " + t.productID_parse)
	fmt.Println("Secondata1 : " + t.seconData_parse)
	fmt.Println("Itobranch : " + t.branchCode_parse)
	fmt.Println("Proccode : " + t.proccode_parse)

	if t.proccode_parse == "12,000" || t.proccode_parse == "2,000" {
		if (t.productID_parse == "ATM" || t.seconData_parse == "ATM") && (t.branchCode_parse == "960" || t.branchCode_parse == "7,168" || t.branchCode_parse == "7,169" || t.branchCode_parse == "8,192"){
			return true
		} else if (t.productID_parse == "SHA" || t.seconData_parse == "SHA") && (t.branchCode_parse == "960" || t.branchCode_parse == "7,168" || t.branchCode_parse == "8,192"){
			return true
		} else if (t.productID_parse == "HMB" || t.seconData_parse == "HMB") && (t.branchCode_parse == "960" || t.branchCode_parse == "7,168" || t.branchCode_parse == "8,192"){
			return true
		} else if (t.productID_parse == "BCA" || t.seconData_parse == "BCA") && (t.branchCode_parse == "960" || t.branchCode_parse == "7,168" || t.branchCode_parse == "8,192"){
			return true
		} else if (t.productID_parse == "MC" || t.seconData_parse == "MC") && (t.branchCode_parse == "960" || t.branchCode_parse == "7,168" || t.branchCode_parse == "8,192"){
			return true
		} else if (t.productID_parse == "MEPS" || t.seconData_parse == "MEPS") && (t.branchCode_parse == "960" || t.branchCode_parse == "7,168" || t.branchCode_parse == "8,192"){
			return true
		} else if (t.productID_parse == "CI" || t.seconData_parse == "CI") && (t.branchCode_parse == "960" || t.branchCode_parse == "7,168" || t.branchCode_parse == "8,192"){
			return true
		} else if (t.productID_parse == "MS" || t.seconData_parse == "MS") && (t.branchCode_parse == "960" || t.branchCode_parse == "7,168" || t.branchCode_parse == "8,192"){
			return true
		} else if (t.productID_parse == "EDCBRI" || t.seconData_parse == "EDCBRI") && (t.branchCode_parse == "960" || t.branchCode_parse == "7,168" || t.branchCode_parse == "7,169" || t.branchCode_parse == "8,192"){
			return true
		} else {
			return false
		}
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
	var returnLocationAnalyst Returnlocationanalyst
	var parseLocationAnalyst Parselocationanalyst
	var parserSetRedisData Parsesetredisdata
  var tempSequence Tempsequence

	i.Nokartu = j.Cardno_prosw
	i.Amounttransaction = j.Txamount_prosw
	i.Terminaltype	= j.Terminaltype_prosw
	i.Terminalid = j.Termid_prosw
	x := j.Termid_prosw
  x = x[len(x)-7:]
	i.Terminalinformation = j.Narrative_prosw + "," +j.Firstdata8_prosw
	i.Statustransaction = j.Responcode_prosw
	i.Transdate = j.Transdate_prosw
  i.Transdate = strings.Replace(i.Transdate, " 00:00:00", "", -1)
  iTranstimeConvert, _ := strconv.Atoi(j.Transtime_prosw)
	i.Transtime = secondsToTimeStamp(iTranstimeConvert)
  i.Transtime = i.Transtime[0:5]
  fmt.Println("Transtime : " + i.Transtime)
	fmt.Println("Transdate : " + i.Transdate)
	fmt.Println("Source File : " + j.Source_file_prosw)
	i.Descriptiontransaction = j.Narrative_prosw

	locationAnalyst := true
	//sequenceAnalyst := false
	prevTransactionData := true
	fmt.Println("Terminal Type : " + i.Terminaltype)
	fmt.Println("Terminal ID : " + x)
	fmt.Println("Product ID : " + j.Productid_prosw)
	// Lookup to Database for ATM BRI Location
	if i.Terminaltype == "ATM_BRI" || i.Terminaltype == "ATM MERAH PUTIH" {

		if i.Terminaltype == "ATM_BRI" {
			sX,_ := leftpad.PadChar(x, 7, '0')
			x = "BRI"+sX
			fmt.Println("Terminal ID : " + x)
		}
		locationString = findLocation(x)

		if locationString.statusInformation == "FOUND" {
			i.Longitude = locationString.longitudeTerminal
			i.Latitude = locationString.latitudeTerminal
		} else {
			locationAnalyst = false
		}
	}

	if locationAnalyst == false {
		fmt.Println("Cek Google : CEK GOOGLE")
		fmt.Println("String To Google : " + i.Terminalinformation)
		locationString = callGoogleMaps(i.Terminalinformation)
		if locationString.longitudeTerminal != "" {
			i.Longitude = locationString.longitudeTerminal
			i.Latitude = locationString.latitudeTerminal
		} else {
			fmt.Println("Cek Google : NOT FOUND")
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

		DateTimeString = i.Transdate + " " + i.Transtime
		DateTimeString = DateTimeString[0:16]
		fmt.Println("Date Time String : " + DateTimeString)

		for resultsSequence.Next() {
			var tempSequence Tempsequence

			err := resultsSequence.Scan(&tempSequence.kartuTemp, &tempSequence.responseTemp, &tempSequence.matauangTemp, &tempSequence.waktuTemp, &tempSequence.scoringTemp, &tempSequence.idRecordTemp)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}

			// Selisih waktu
			if n == 0 { // Melakukan perbandingan antara transaksi actual dengan history terakhir.
				recordIDTemp = tempSequence.idRecordTemp
				PastDateTimeString = tempSequence.waktuTemp
        PastDateTimeString = PastDateTimeString[0:16]

				parserStartDate, errParse1 := time.Parse(constFormatDateTime, DateTimeString)
				parserEndDate, errParse2 := time.Parse(constFormatDateTime, PastDateTimeString)

				if errParse1 != nil {
          i.Remark = "Error : "+errParse1.Error()
          fmt.Println(errParse1.Error())
        }
				if errParse2 != nil {
          i.Remark = "Error : "+errParse2.Error()
          fmt.Println(errParse2.Error())
        }

		    delta := parserStartDate.Sub(parserEndDate)
		    SelisihWaktu := delta.Minutes() // Selisih dalam satuan jam

				if j.Ccycode_prosw == "360" && tempSequence.matauangTemp == "360" {
					if SelisihWaktu < 61 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_1"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_2"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_3"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_4"))
						}
					} else if SelisihWaktu > 60 && SelisihWaktu < 301 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_5"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_6"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_7"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_8"))
						}
					} else if SelisihWaktu > 300 && SelisihWaktu < 721 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_9"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_10"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_11"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_12"))
						}
					} else if SelisihWaktu > 720 && SelisihWaktu < 1441 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_13"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_14"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_15"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_16"))
						}
					} else if SelisihWaktu > 1440 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_17"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_18"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_19"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_20"))
						}
					}
				} else if j.Ccycode_prosw == "360" && tempSequence.matauangTemp != "360" {
					if SelisihWaktu < 61 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_21"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_22"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_23"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_24"))
						}
					} else if SelisihWaktu > 60 && SelisihWaktu < 301 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_25"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_26"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_27"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_28"))
						}
					} else if SelisihWaktu > 300 && SelisihWaktu < 721 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_29"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_30"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_31"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_32"))
						}
					} else if SelisihWaktu > 720 && SelisihWaktu < 1441 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_33"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_34"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_35"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_36"))
						}
					} else if SelisihWaktu > 1440 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_37"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_38"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_39"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_40"))
						}
					}
				} else if j.Ccycode_prosw != "360" && tempSequence.matauangTemp == "360" {
					if SelisihWaktu < 61 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_41"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_42"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_43"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_44"))
						}
					} else if SelisihWaktu > 60 && SelisihWaktu < 301 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_45"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_46"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_47"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_48"))
						}
					} else if SelisihWaktu > 300 && SelisihWaktu < 721 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_49"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_50"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_51"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_52"))
						}
					} else if SelisihWaktu > 720 && SelisihWaktu < 1441 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_53"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_54"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_55"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_56"))
						}
					} else if SelisihWaktu > 1440 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_57"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_58"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_59"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_60"))
						}
					}
				} else if j.Ccycode_prosw != "360" && tempSequence.matauangTemp != "360" {
					if SelisihWaktu < 61 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_61"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_62"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_63"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_64"))
						}
					} else if SelisihWaktu > 60 && SelisihWaktu < 301 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_65"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_66"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_67"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_68"))
						}
					} else if SelisihWaktu > 300 && SelisihWaktu < 721 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_69"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_70"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_71"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_72"))
						}
					} else if SelisihWaktu > 720 && SelisihWaktu < 1441 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_73"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_74"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_75"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_76"))
						}
					} else if SelisihWaktu > 1440 {
						if i.Statustransaction == "00" && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_77"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && tempSequence.responseTemp == "00" {
								score += convertStringToInt(os.Getenv("SA_78"))
						} else if i.Statustransaction == "00" && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_79"))
						} else if (i.Statustransaction == "51" || i.Statustransaction == "61") && (tempSequence.responseTemp == "51" || tempSequence.responseTemp == "61") {
								score += convertStringToInt(os.Getenv("SA_80"))
						}
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
    /*if recordIDTemp != "" {
        deleteOldSequence(recordIDTemp, i.Nokartu)
    }
		*/
		// menambahkan record baru
		tempSequence.kartuTemp = j.Cardno_prosw
		tempSequence.responseTemp = i.Statustransaction
		tempSequence.matauangTemp = j.Ccycode_prosw
		tempSequence.scoringTemp = strconv.Itoa(scorePhase1)
		tempSequence.waktuTemp = DateTimeString

    fmt.Println("Kartu : " + j.Cardno_prosw)
    fmt.Println("Response : " + i.Statustransaction)
    fmt.Println("Mata Uang : " + j.Ccycode_prosw)
    fmt.Println("Scoring : " + strconv.Itoa(scorePhase1))
    fmt.Println("Date Time : " + DateTimeString)

		tempSequence.setDataSequence()

		// Convert Scoring Fraud base on Sequence.
		if score <= 100 {
			i.Labelfraud_sequence = "Hijau"
		} else if score <= 300 && score > 100 {
    	i.Labelfraud_sequence = "Kuning"
    } else if score > 300{
    	i.Labelfraud_sequence = "Merah"
    }

		i.Scorefraud_sequence = strconv.Itoa(score)
    i.Kodematauang = j.Ccycode_prosw

		// ** END SEQUENCE ANALYST ** //
    tambahanRemark := ""
    fmt.Println("Location Analyst : " + strconv.FormatBool(locationAnalyst))
    fmt.Println("previousTransaction : " + strconv.FormatBool(prevTransactionData))

    tLocationAnalyst_0 := time.Now()
	if locationAnalyst {
			parseLocationAnalyst.currentLatitude = i.Latitude
			parseLocationAnalyst.currentLongitude = i.Longitude
			parseLocationAnalyst.pastLatitude = i.Prevlatitude
			parseLocationAnalyst.pastLongitude = i.Prevlongitude
			parseLocationAnalyst.currentDate = i.Transdate
			parseLocationAnalyst.currentTime = i.Transtime
			parseLocationAnalyst.pastDate = i.Prevtransdate
			parseLocationAnalyst.pastTime = i.Prevtranstime
			parseLocationAnalyst.currentStatus = i.Statustransaction
			parseLocationAnalyst.pastStatus = i.Prevstatus

			if prevTransactionData {
				returnLocationAnalyst = parseLocationAnalyst.calculateLocationAnalyst()
				i.Labelfraud_location = returnLocationAnalyst.labelFraud
				i.Scorefraud_location = returnLocationAnalyst.scoreFraud
				i.Remark = returnLocationAnalyst.errorCalculate
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
  tLocationAnalyst_1 := time.Now()
  fmt.Println("The call took %v to run.\n", tLocationAnalyst_1.Sub(tLocationAnalyst_0))

	i.Remark = i.Remark + ";" + tambahanRemark
	i.Remark = previousTransaction.statusInformation
	i.Status = prevTransactionData
	return i

}


func (j Parselocationanalyst) calculateLocationAnalyst() Returnlocationanalyst {
	// Menghitung Score dengan menggunakan
		// - Selisih Waktu
		// - Status Transaksi Sebelum dengan Saat ini.
		// - Selisih Jarak antara terminal ID sebelum dengan sesaat ini.
	indicator := "Kosong"
	var result = Returnlocationanalyst{}
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
		/*
    // Menentukan Score
    // Berdasarkan Status Transaksi Sebelumnya
    if j.pastStatus == "51" { // 51 - Saldo tidak cukup
    	Score += 40
    } else if j.pastStatus == "61" { //"61 - Transaksi melebihi limit transaksi harian"
    	Score += 50
    } else if j.pastStatus == "00" { //
    	Score += 10
    } else {
    	Score += 0
    }

    // Berdasarkan Status Transaksi Saat Ini
    if j.currentStatus == "51" { // 51 - Saldo tidak cukup
    	Score += 40
    } else if j.currentStatus == "61" { //"61 - Transaksi melebihi limit transaksi harian"
    	Score += 50
    } else if j.currentStatus == "00" { //
    	Score += 10
    } else {
    	Score += 0
    }
		*/
    // Berdasarkan kombinasi Selisih Waktu dan Jarak
    // Jika waktu kurang dari 1 jam, menggunakan acuan kendaraan 80-120 km per jam
    // Jika waktu antara 1 - 5 jam, menggunakan acuan kendaraan 200 km per jam (pesawat dan kereta)
    if SelisihWaktu < 1 && SelisihJarak < 100 {
    	Score += convertStringToInt(os.Getenv("LA_1"))
    } else if SelisihWaktu < 1 && SelisihJarak > 100 {
    	Score += convertStringToInt(os.Getenv("LA_2"))
    } else if SelisihWaktu < 5 && SelisihWaktu > 1 && 1000 > SelisihJarak && SelisihJarak > 100 {
    	Score += convertStringToInt(os.Getenv("LA_3"))
    } else if SelisihWaktu < 12 && SelisihWaktu > 5 && SelisihJarak > 1000 {
    	Score += convertStringToInt(os.Getenv("LA_4"))
    } else if SelisihWaktu > 12 && SelisihJarak > 1000 {
    	Score += convertStringToInt(os.Getenv("LA_5"))
    } else {
    	Score += convertStringToInt(os.Getenv("LA_6"))
    }

    if Score > 99 {
    	indicator = "Merah"
    } else if Score < 99 && Score > 50 {
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

func convertStringToInt(sString string) int {
  valueInt, _ := strconv.Atoi(sString)

  return valueInt
}
