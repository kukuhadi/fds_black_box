package main

type Proswmin struct{
	Nokartu									string		`json:"nokartu"`
	Amounttransaction				string		`json:"amounttransaction"`
	Terminaltype						string		`json:"terminaltype"`
	Terminalid							string		`json:"terminalid"`
	Terminalinformation			string		`json:"terminalinformation"`
	Statustransaction				string		`json:"statustransaction"`
	Transdate								string		`json:"transdate"`
	Transtime								string		`json:"transtime"`
	Descriptiontransaction	string 		`json:"descriptiontransaction"`
}

type Prosw struct{
	Source_file_prosw	 		 	 string	 	`json:"source_file_prosw"`
	Isocode_prosw	 		 	     string	 	`json:"isocode_prosw"`
	Reqtype_prosw	 		 	     string	 	`json:"reqtype_prosw"`
	Datepost_prosw	 		 	   string	 	`json:"datepost_prosw"`
	Sendfile_prosw	 		 	   string	 	`json:"sendfile_prosw"`
	Txtime_prosw	 		 	     string	 	`json:"txtime_prosw"`
	Subtype_prosw	 		 	     string	 	`json:"subtype_prosw"`
	Txcode_prosw	 		 	     string	 	`json:"txcode_prosw"`
	Brnchcode_prosw	 		 	   string	 	`json:"brnchcode_prosw"`
	Authotel_prosw	 		 	   string	 	`json:"authotel_prosw"`
	Tellid_prosw	 		 	     string	 	`json:"tellid_prosw"`
	Txseqnum_prosw	 		 	   string	 	`json:"txseqnum_prosw"`
	Telseqnum_prosw	 		 	   string	 	`json:"telseqnum_prosw"`
	Remoteaccno_prosw	 		 	 string	 	`json:"remoteaccno_prosw"`
	Toaccno_prosw	 		 	     string	 	`json:"toaccno_prosw"`
	Sendbranch_prosw	 		 	 string	 	`json:"sendbranch_prosw"`
	Ccycode_prosw	 		 	     string	 	`json:"ccycode_prosw"`
	Actname_prosw	 		 	     string	 	`json:"actname_prosw"`
	Pbbalnc_prosw	 		 	     string	 	`json:"pbbalnc_prosw"`
	Avlblnc_prosw	 		 	     string	 	`json:"avlblnc_prosw"`
	Txamount_prosw	 		 	   string	 	`json:"txamount_prosw"`
	Chqnumber_prosw	 		 	   string	 	`json:"chqnumber_prosw"`
	Chqdate_prosw	 		 	     string	 	`json:"chqdate_prosw"`
	Narrative_prosw	 		 	   string	 	`json:"narrative_prosw"`
	Linepb_prosw	 		 	     string	 	`json:"linepb_prosw"`
	Statproc_prosw	 		 	   string	 	`json:"statproc_prosw"`
	Proccode_prosw	 		     string	 	`json:"proccode_prosw"`
	Withpassbook_prosw	 		 string	 	`json:"withpassbook_prosw"`
	Responcode_prosw	 		 	 string	 	`json:"responcode_prosw"`
	Moreprint_prosw	 		 	   string	 	`json:"moreprint_prosw"`
	Savdate1_prosw	 		 	   string	 	`json:"savdate1_prosw"`
	Savdate2_prosw	 		 	   string	 	`json:"savdate2_prosw"`
	Savdate3_prosw	 		 	   string	 	`json:"savdate3_prosw"`
	Savdate4_prosw	 		 	   string	 	`json:"savdate4_prosw"`
	Savdate5_prosw	 		 	   string	 	`json:"savdate5_prosw"`
	Savcode1_prosw	 		 	   string	 	`json:"savcode1_prosw"`
	Savcode2_prosw	 		 	   string	 	`json:"savcode2_prosw"`
	Savcode3_prosw	 		 	   string	 	`json:"savcode3_prosw"`
	Savcode4_prosw	 		 	   string	 	`json:"savcode4_prosw"`
	Savcode5_prosw	 		 	   string	 	`json:"savcode5_prosw"`
	Savtxtype1_prosw	 		 	 string	 	`json:"savtxtype1_prosw"`
	Savtxtype2_prosw	 		 	 string	 	`json:"savtxtype2_prosw"`
	Savtxtype3_prosw	 		 	 string	 	`json:"savtxtype3_prosw"`
	Savtxtype4_prosw	 		 	 string	 	`json:"savtxtype4_prosw"`
	Savtxtype5_prosw	 		 	 string	 	`json:"savtxtype5_prosw"`
	Savamount1_prosw	 		 	 string	 	`json:"savamount1_prosw"`
	Savamount2_prosw	 		 	 string	 	`json:"savamount2_prosw"`
	Savamount3_prosw	 		 	 string	 	`json:"savamount3_prosw"`
	Savamount4_prosw	 		 	 string	 	`json:"savamount4_prosw"`
	Savamount5_prosw	 		 	 string	 	`json:"savamount5_prosw"`
	Savtlrid1_prosw	 		 	   string	 	`json:"savtlrid1_prosw"`
	Savtlrid2_prosw	 		 	   string	 	`json:"savtlrid2_prosw"`
	Savtlrid3_prosw	 		 	   string	 	`json:"savtlrid3_prosw"`
	Savtlrid4_prosw	 		 	   string	 	`json:"savtlrid4_prosw"`
	Savtlrid5_prosw	 		 	   string	 	`json:"savtlrid5_prosw"`
	Savlinepb1_prosw	 		 	 string	 	`json:"savlinepb1_prosw"`
	Savlinepb2_prosw	 		 	 string	 	`json:"savlinepb2_prosw"`
	Savlinepb3_prosw	 		 	 string	 	`json:"savlinepb3_prosw"`
	Savlinepb4_prosw	 		 	 string	 	`json:"savlinepb4_prosw"`
	Savlinepb5_prosw	 		 	 string	 	`json:"savlinepb5_prosw"`
	Savpbbal1_prosw	 		 	   string	 	`json:"savpbbal1_prosw"`
	Savpbbal2_prosw	 		 	   string	 	`json:"savpbbal2_prosw"`
	Savpbbal3_prosw	 		 	   string	 	`json:"savpbbal3_prosw"`
	Savpbbal4_prosw	 		 	   string	 	`json:"savpbbal4_prosw"`
	Savpbbal5_prosw	 		 	   string	 	`json:"savpbbal5_prosw"`
	Firstdata1_prosw	 		 	 string	 	`json:"firstdata1_prosw"`
	Firstdata2_prosw	 		 	 string	 	`json:"firstdata2_prosw"`
	Firstdata3_prosw	 		 	 string	 	`json:"firstdata3_prosw"`
	Firstdata4_prosw	 		 	 string	 	`json:"firstdata4_prosw"`
	Firstdata5_prosw	 		 	 string	 	`json:"firstdata5_prosw"`
	Firstdata6_prosw	 		 	 string	 	`json:"firstdata6_prosw"`
	Firstdata7_prosw	 		 	 string	 	`json:"firstdata7_prosw"`
	Firstdata8_prosw	 		 	 string	 	`json:"firstdata8_prosw"`
	Firstdata9_prosw	 		 	 string	 	`json:"firstdata9_prosw"`
	Firstdata10_prosw	 		 	 string	 	`json:"firstdata10_prosw"`
	Seconddata1_prosw	 		 	 string	 	`json:"seconddata1_prosw"`
	Seconddata2_prosw	 		 	 string	 	`json:"seconddata2_prosw"`
	Seconddata3_prosw	 		 	 string	 	`json:"seconddata3_prosw"`
	Seconddata4_prosw	 		 	 string	 	`json:"seconddata4_prosw"`
	Seconddata5_prosw	 		 	 string	 	`json:"seconddata5_prosw"`
	Seconddata6_prosw	 		 	 string	 	`json:"seconddata6_prosw"`
	Seconddata7_prosw	 		 	 string	 	`json:"seconddata7_prosw"`
	Seconddata8_prosw	 		 	 string	 	`json:"seconddata8_prosw"`
	Seconddata9_prosw	 		 	 string	 	`json:"seconddata9_prosw"`
	Seconddata10_prosw	 		 string	 	`json:"seconddata10_prosw"`
	Seconddata11_prosw	 		 string	 	`json:"seconddata11_prosw"`
	Seconddata12_prosw	 		 string	 	`json:"seconddata12_prosw"`
	Seconddata13_prosw	 		 string	 	`json:"seconddata13_prosw"`
	Seconddata14_prosw	 		 string	 	`json:"seconddata14_prosw"`
	Seconddata15_prosw	 		 string	 	`json:"seconddata15_prosw"`
	Seconddata16_prosw	 		 string	 	`json:"seconddata16_prosw"`
	Seconddata17_prosw	 		 string	 	`json:"seconddata17_prosw"`
	Seconddata18_prosw	 		 string	 	`json:"seconddata18_prosw"`
	Seconddata19_prosw	 		 string	 	`json:"seconddata19_prosw"`
	Seconddata20_prosw	 		 string	 	`json:"seconddata20_prosw"`
	Transdate_prosw	 		 	   string	 	`json:"transdate_prosw"`
	Actamount_prosw	 		 	   string	 	`json:"actamount_prosw"`
	Transtime_prosw	 		 	   string	 	`json:"transtime_prosw"`
	Termid_prosw	 		 	     string	 	`json:"termid_prosw"`
	Prodtype_prosw	 		 	   string	 	`json:"prodtype_prosw"`
	Idatelog_prosw	 		 	   string	 	`json:"idatelog_prosw"`
	Itimelog_prosw	 		 	   string	 	`json:"itimelog_prosw"`
	Itobranch_prosw	 		     string	 	`json:"itobranch_prosw"`
	Istatus_prosw	 		 	     string	 	`json:"istatus_prosw"`
	Istime_prosw	 		 	     string	 	`json:"istime_prosw"`
	Icounter_prosw	 		 	   string	 	`json:"icounter_prosw"`
	Itgvaluta_prosw	 		     string	 	`json:"itgvaluta_prosw"`
	Linkid_prosw	 		 	     string	 	`json:"linkid_prosw"`
	Prodid_prosw	 		 	     string	 	`json:"prodid_prosw"`
	Cardno_prosw	 		 	     string	 	`json:"cardno_prosw"`
	Termtellid_prosw	 		 	 string	 	`json:"termtellid_prosw"`
	Txseqnum2_prosw	 		 	   string	 	`json:"txseqnum2_prosw"`
	Channelid_prosw	 		     string	 	`json:"channelid_prosw"`
	Flagresend_prosw	 		 	 string	 	`json:"flagresend_prosw"`
	Inumber_prosw	 		 	     string	 	`json:"inumber_prosw"`
	Productid_prosw	 		     string	 	`json:"productid_prosw"`
	Terminaltype_prosw	 		 string	 	`json:"TerminalType_prosw"`
	Isreversal_prosw	 		 	 string	 	`json:"IsReversal_prosw"`
	Issuccess_prosw					 string 	`json:"IsSuccess_prosw"`
	termid2_prosw	 					 string	 	`json:"termid2_prosw"`
	Brnchcode2_prosw	 		 	 string	 	`json:"brnchcode2_prosw"`
	Remoteaccno2_prosw	 		 string	 	`json:"remoteaccno2_prosw"`
	Transdate_origin_prosw	 string	 	`json:"transdate_origin_prosw"`
	Transtime_origin_prosw	 string	 	`json:"transtime_origin_prosw"`
	Kartu_prosw	 		 	       string	 	`json:"kartu_prosw"`
	Acctno_prosw	 		 	     string	 	`json:"acctno_prosw"`
	Branch_acctno_prosw	 		 string	 	`json:"branch_acctno_prosw"`
	Branch_terminal_prosw	 	 string	 	`json:"branch_terminal_prosw"`
	Id_fee_prosw	 		 	     string	 	`json:"id_fee_prosw"`
	Fee_prosw	 		 	         string	 	`json:"fee_prosw"`
	Rekening_ia_prosw	 		 	 string	 	`json:"rekening_ia_prosw"`
}

type Completemessage struct{
	Nokartu									string		`json:"nokartu"`
	Amounttransaction				string		`json:"amounttransaction"`
	Terminaltype						string		`json:"terminaltype"`
	Terminalid							string		`json:"terminalid"`
	Terminalinformation			string		`json:"terminalinformation"`
	Statustransaction				string		`json:"statustransaction"`
	Transdate								string		`json:"transdate"`
	Transtime								string		`json:"transtime"`
	Descriptiontransaction	string 		`json:"descriptiontransaction"`
	Longitude								string		`json:"longitude"`
	Latitude								string		`json:"latitude"`
	Prevterminalid					string 		`json:"prevterminalid"`
	Prevlongitude						string		`json:"prevlongitude"`
	Prevlatitude						string		`json:"prevlatitude"`
	Prevtransdate						string 		`json:"prevtransdate"`
	Prevtranstime						string 		`json:"prevtranstime"`
	Prevstatus							string 		`json:"prevstatus"`
	Labelfraud_location			string 		`json:"labelfraud_location"`
	Scorefraud_location			string 		`json:"scorefraud_location"`
	Labelfraud_sequence			string 		`json:"labelfraud_sequence"`
	Scorefraud_sequence			string 		`json:"scorefraud_sequence"`
	Remark									string 		`json:"remark"`
	Status									bool			`json:"status"`
	Kodematauang						string		`json:"kodematauang"`
}

type Locationterminal struct{
	latitudeTerminal, longitudeTerminal, statusInformation 	string
}

type Previoustransaction struct{
	latitudeTerminal, longitudeTerminal, statusInformation, statusTransaction, dateTransaction, timeTransaction 	string
}

type Parsecalculatescore struct{
	currentLatitude, currentLongitude, pastLatitude, pastLongitude, currentDate, currentTime, currentStatus, pastDate, pastTime, pastStatus string
}

type Parsesetredisdata struct{
	latitudeTerminal, longitudeTerminal, statusTransaction, dateTransaction, timeTransaction, nomerKartu string
}

type Returncalculatescore struct {
	labelFraud, scoreFraud, selisihWaktu, selisihJarak, errorCalculate string
}

type ParseSequenceanalyst struct{
	currenctDateTime_seq, previousDateTime_seq, currentStatus_seq, previousStatus_seq, currentCurrency_seq, previousCurrency_seq string
}

type Tempsequence struct {
	kartuTemp, responseTemp, matauangTemp, waktuTemp, scoringTemp, idRecordTemp string
}
