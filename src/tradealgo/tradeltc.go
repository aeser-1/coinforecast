package tradealgo

import (
	"db"
	"dependency"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func LTCAlgo() {
	var buy, sell, buyold, sellold bool
	var emaold, smaold [2]float64
	var high, low float64
	ema := dependency.EMA(4, "LTC", true, 0)
	sma := dependency.SMA(5, "LTC", true, 0)
	_, stoc := dependency.STOC(40, 20, true, "LTC", true)
	_, stoc1 := dependency.STOC(10, 5, true, "LTC", false)
	_, stocd := dependency.STOCD(10, 5, true, "LTC", true)
	boll := dependency.Bollinger(14, "LTC", 2.5, true)
	bollsell := dependency.Bollinger(14, "LTC", 2, true)
	trix := dependency.TRIX(30, "LTC", true)
	ltccurr := db.DBRead(1, "LTC")

	db := db.DBconn()
	defer db.Close()

	select1, err1 := db.Query("SELECT sma,ema FROM METRICS WHERE Symbol=? ORDER BY id DESC LIMIT 2", ltccurr.Symbol)
	if err1 != nil {
		panic(err1.Error())
	}

	var counter int
	for select1.Next() {
		err1 := select1.Scan(&smaold[counter], &emaold[counter])
		if err1 != nil {
			panic(err1)
		}
		counter++
	}

	select2, err2 := db.Query("SELECT buy,sell FROM DECISIONS WHERE Symbol=? ORDER BY id DESC LIMIT 1", ltccurr.Symbol)
	if err2 != nil {
		panic(err2.Error())
	}

	for select2.Next() {
		err2 := select2.Scan(&buyold, &sellold)
		if err1 != nil {
			panic(err2)
		}
	}

	select3, err3 := db.Query("SELECT High,Low FROM HighLow WHERE Symbol=? ORDER BY id DESC LIMIT 1", ltccurr.Symbol)
	if err3 != nil {
		panic(err3.Error())
	}

	for select3.Next() {
		err3 := select3.Scan(&high, &low)
		if err1 != nil {
			panic(err3)
		}
	}

	if ema > sma {
		if stoc < 82 && stoc > 10 && low > (0.995*boll[2]) && stoc1 > stocd && trix > -0.13 && !buyold {
			buy = true
		}
	}

	if !buyold && !sellold {

		sell = false

	} else {
		if (ema < sma || stoc > 88 || high > (1.014*bollsell[0])) && !sellold {

			sell = true

		}
	}

	fmt.Println("New:")
	fmt.Println("EMA 	:", ema)
	fmt.Println("SMA 	:", sma)
	fmt.Println("Old:")
	fmt.Println("EMA 	:", emaold[1])
	fmt.Println("SMA 	:", smaold[1])
	fmt.Println("STOC		:", stoc)
	fmt.Println("STOC1		:", stoc1)
	fmt.Println("STOCD		:", stocd)
	fmt.Println("Bollinger	:", boll)
	fmt.Println("TRIX	:", trix)
	fmt.Println("Price	:", ltccurr.Price[0])
	fmt.Println("BuyLTC	:", buy)
	fmt.Println("SellLTC:", sell)

	if buy || sell {
		insert1, err1 := db.Query("INSERT INTO DECISIONS (RowID,Symbol,Price,Buy,Sell) VALUES(?,?,?,?,?)", ltccurr.Rowid[0], ltccurr.Symbol, ltccurr.Price[0], buy, sell)
		if err1 != nil {
			panic(err1.Error())
		}
		defer insert1.Close()
	}
}
