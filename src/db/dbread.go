package db

import (
	_ "github.com/go-sql-driver/mysql"
)

type CoinData struct {
	Symbol string
	Rowid  []int
	Price  []float64
}

type Decision struct {
	Symbol   string
	Decision string
}

func DBRead(amount int, symbol string) CoinData { //read db
	var current CoinData
	db := DBconn()
	defer db.Close()
	current.Symbol = symbol
	current.Rowid = make([]int, amount)
	current.Price = make([]float64, amount)
	sqlprice := [...]string{"SELECT price FROM ", " ORDER BY id DESC LIMIT ?"}
	sqlid := [...]string{"SELECT id FROM ", " ORDER BY id DESC LIMIT ?"}
	sqlquery1 := sqlprice[0] + symbol + sqlprice[1]
	sqlquery2 := sqlid[0] + symbol + sqlid[1]

	select1, err1 := db.Query(sqlquery1, amount)
	if err1 != nil {
		panic(err1.Error())
	}

	select2, err2 := db.Query(sqlquery2, amount)
	if err2 != nil {
		panic(err2.Error())
	}

	var counter int
	for select1.Next() {
		err1 := select1.Scan(&current.Price[counter])
		if err1 != nil {
			panic(err1)
		}
		counter++
	}

	counter = 0
	for select2.Next() {
		err2 := select2.Scan(&current.Rowid[counter])
		if err2 != nil {
			panic(err2)
		}
		counter++
	}

	return current
}

func DBReadHiLo(amount int, symbol string) ([][]float64, int, string) {
	var rowid int

	high := make([]float64, amount)
	low := make([]float64, amount)
	close := make([]float64, amount)
	prices := make([][]float64, 3, amount)

	db := DBconn()
	defer db.Close()

	select1, err1 := db.Query("SELECT RowID FROM HighLow WHERE Symbol=? ORDER BY id DESC LIMIT 1", symbol)
	if err1 != nil {
		panic(err1.Error())
	}

	for select1.Next() {
		err1 := select1.Scan(&rowid)
		if err1 != nil {
			panic(err1)
		}
	}

	select2, err2 := db.Query("SELECT high,low,close FROM HighLow WHERE Symbol=? ORDER BY id DESC LIMIT ?", symbol, amount)
	if err2 != nil {
		panic(err2.Error())
	}
	counter := 0
	for select2.Next() {
		err2 := select2.Scan(&high[counter], &low[counter], &close[counter])
		if err2 != nil {
			panic(err2)
		}
		counter++
	}

	prices[0] = high
	prices[1] = low
	prices[2] = close

	return prices, rowid, symbol
}

func DBReadDecision(symbol string) Decision {
	var dec Decision
	var decision string
	var buy, sell bool
	db := DBconn()
	defer db.Close()

	select1, err1 := db.Query("SELECT Buy,Sell FROM DECISIONS WHERE Symbol=? ORDER BY id DESC LIMIT 1", symbol)
	if err1 != nil {
		panic(err1.Error())
	}

	counter := 0
	for select1.Next() {
		err1 := select1.Scan(&buy, &sell)
		if err1 != nil {
			panic(err1)
		}
		counter++
	}

	if buy {
		decision = "buy"
	}

	if sell {
		decision = "sell"
	}

	dec.Symbol = symbol
	dec.Decision = decision

	return dec
}
