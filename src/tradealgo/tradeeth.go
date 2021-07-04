package tradealgo

import (
	"db"
	"dependency"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ETHAlgo() {
	
	//Sample Strategy --> BCH
	ema := dependency.EMA(10, "ETH", true, 0)
	sma := dependency.SMA(10, "ETH", true, 0)
	_, stoc := dependency.STOC(14, 6, true, "ETH", true)
	_, stoc1 := dependency.STOC(14, 6, true, "ETH", false)
	_, stocd := dependency.STOCD(14, 6, true, "ETH", true)
	boll := dependency.Bollinger(12, "ETH", 2, true)
	bollsell := dependency.Bollinger(12, "ETH", 2, true)
	trix := dependency.TRIX(20, "ETH", true)
	bchcurr := db.DBRead(1, "ETH")
	ethcurr := db.DBRead(1, "ETH")

	db := db.DBconn()
	defer db.Close()

	select1, err1 := db.Query("SELECT sma,ema FROM METRICS WHERE Symbol=? ORDER BY id DESC LIMIT 2", ethcurr.Symbol)
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

	select2, err2 := db.Query("SELECT buy,sell FROM DECISIONS WHERE Symbol=? ORDER BY id DESC LIMIT 1", ethcurr.Symbol)
	if err2 != nil {
		panic(err2.Error())
	}

	for select2.Next() {
		err2 := select2.Scan(&buyold, &sellold)
		if err1 != nil {
			panic(err2)
		}
	}

	select3, err3 := db.Query("SELECT High,Low FROM HighLow WHERE Symbol=? ORDER BY id DESC LIMIT 1", ethcurr.Symbol)
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
	fmt.Println("Price	:", ethcurr.Price[0])
	fmt.Println("BuyETH	:", buy)
	fmt.Println("SellETH:", sell)

	if buy || sell {
		insert1, err1 := db.Query("INSERT INTO DECISIONS (RowID,Symbol,Price,Buy,Sell) VALUES(?,?,?,?,?)", ethcurr.Rowid[0], ethcurr.Symbol, ethcurr.Price[0], buy, sell)
		if err1 != nil {
			panic(err1.Error())
		}
		defer insert1.Close()
	}
}
