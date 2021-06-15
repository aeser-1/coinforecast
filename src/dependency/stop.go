package dependency

import (
	"db"

	_ "github.com/go-sql-driver/mysql"
)

func StopLoss(symbol string, percentage float64) {
	var price float64
	var buy, sell bool

	currcoin := db.DBRead(1, symbol)

	dbase := db.DBconn()
	defer dbase.Close()

	select1, err1 := dbase.Query("SELECT price,buy FROM DECISIONS WHERE Symbol=? ORDER BY id DESC LIMIT 1", symbol)
	if err1 != nil {
		panic(err1.Error())
	}

	for select1.Next() {
		err1 := select1.Scan(&price, &buy)
		if err1 != nil {
			panic(err1)
		}
	}

	if buy {
		if percentage <= (((price - currcoin.Price[0]) / price) * 100) {
			sell = true
		}
	}

	if sell {
		insert1, err1 := dbase.Query("INSERT INTO DECISIONS (RowID,Symbol,Price,Buy,Sell,StopLoss,MovingStop) VALUES(?,?,?,?,?,?,?)", currcoin.Rowid[0], currcoin.Symbol, currcoin.Price[0], false, sell, true, false)
		if err1 != nil {
			panic(err1.Error())
		}
		defer insert1.Close()
	}

}

func MovingStop(symbol string, percentage float64) {
	var price, max float64
	var buy, sell bool
	var rowid int

	sqlprice := [...]string{"SELECT max(price) FROM ", " WHERE id>=?"}
	sqlquery := sqlprice[0] + symbol + sqlprice[1]

	currcoin := db.DBRead(1, symbol)

	dbase := db.DBconn()
	defer dbase.Close()

	select1, err1 := dbase.Query("SELECT rowid,price,buy FROM DECISIONS WHERE Symbol=? ORDER BY id DESC LIMIT 1", symbol)
	if err1 != nil {
		panic(err1.Error())
	}

	for select1.Next() {
		err1 := select1.Scan(&rowid, &price, &buy)
		if err1 != nil {
			panic(err1)
		}
	}

	select2, err2 := dbase.Query(sqlquery, rowid)
	if err2 != nil {
		panic(err2.Error())
	}

	if select2 != nil {
		for select2.Next() {
			err2 := select2.Scan(&max)
			if err2 != nil {
				panic(err2)
			}
		}
	}

	currentprofit := currcoin.Price[0] - price

	if max > currcoin.Price[0] && currentprofit > 0 {

		maxprofit := max - price

		if buy {
			if percentage <= ((maxprofit - currentprofit) / maxprofit * 100) {
				sell = true
			}

		}

		if sell {

			insert1, err1 := dbase.Query("INSERT INTO DECISIONS (RowID,Symbol,Price,Buy,Sell,StopLoss,MovingStop) VALUES(?,?,?,?,?,?,?)", currcoin.Rowid[0], currcoin.Symbol, currcoin.Price[0], false, sell, false, true)
			if err1 != nil {
				panic(err1.Error())
			}
			defer insert1.Close()

		}

	}

}
