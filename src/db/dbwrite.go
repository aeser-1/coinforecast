package db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DBWrite(cryp CryptoData) { //Write into db
	currentTime := time.Now()
	db := DBconn()
	defer db.Close()

	insert1, err := db.Query("INSERT INTO BTC (Name,Symbol,Price,DateTime) VALUES(?,?,?,?)", cryp.Data.BTC.Name, cryp.Data.BTC.Symbol, cryp.Data.BTC.Quote.USD.Price, currentTime)
	if err != nil {
		panic(err.Error())
	}
	defer insert1.Close()

	insert2, err := db.Query("INSERT INTO BCH (Name,Symbol,Price,DateTime) VALUES(?,?,?,?)", cryp.Data.BCH.Name, cryp.Data.BCH.Symbol, cryp.Data.BCH.Quote.USD.Price, currentTime)
	if err != nil {
		panic(err.Error())
	}
	defer insert2.Close()

	insert3, err := db.Query("INSERT INTO ETH (Name,Symbol,Price,DateTime) VALUES(?,?,?,?)", cryp.Data.ETH.Name, cryp.Data.ETH.Symbol, cryp.Data.ETH.Quote.USD.Price, currentTime)
	if err != nil {
		panic(err.Error())
	}
	defer insert3.Close()

	insert4, err := db.Query("INSERT INTO LTC (Name,Symbol,Price,DateTime) VALUES(?,?,?,?)", cryp.Data.LTC.Name, cryp.Data.LTC.Symbol, cryp.Data.LTC.Quote.USD.Price, currentTime)
	if err != nil {
		panic(err.Error())
	}
	defer insert4.Close()

	insert5, err := db.Query("INSERT INTO XRP (Name,Symbol,Price,DateTime) VALUES(?,?,?,?)", cryp.Data.XRP.Name, cryp.Data.XRP.Symbol, cryp.Data.XRP.Quote.USD.Price, currentTime)
	if err != nil {
		panic(err.Error())
	}
	defer insert5.Close()

}

func DBWriteHiLo(symbol []string) {
	var min, max, close float64
	var id int
	var coin CoinData

	db := DBconn()
	defer db.Close()

	for i := 0; i < len(symbol); i++ {
		coin = DBRead(13, symbol[i])
		id = coin.Rowid[0]
		close = coin.Price[0]
		price := coin.Price

		min = price[0]
		for _, price := range price {
			if price < min {
				min = price
			}
		}

		max = price[0]
		for _, price := range price {
			if price > max {
				max = price
			}
		}

		insert1, err1 := db.Query("INSERT INTO HighLow (RowID,Symbol,High,Low,Close) VALUES(?,?,?,?,?)", id, symbol[i], max, min, close)
		if err1 != nil {
			panic(err1.Error())
		}

		defer insert1.Close()

	}

}
