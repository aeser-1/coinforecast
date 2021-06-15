package db

import (
	_ "github.com/go-sql-driver/mysql"
)

func Profit(symbol string) float64 {
	var profitloss, sumBuy, sumSell float64
	var countBuy, countSell int

	db := DBconn()
	defer db.Close()

	err3 := db.QueryRow("SELECT COUNT(*) FROM DECISIONS WHERE Symbol=? && Buy=?", symbol, true).Scan(&countBuy)

	if err3 != nil {
		panic(err3)
	}

	err4 := db.QueryRow("SELECT COUNT(*) FROM DECISIONS WHERE Symbol=? && Sell=?", symbol, true).Scan(&countSell)

	if err4 != nil {
		panic(err4)
	}

	if countBuy > countSell {
		countBuy = countBuy - (countBuy - countSell)
	}

	priceBuy := make([]float64, countBuy)
	priceSell := make([]float64, countSell)

	if countBuy == countSell && countBuy != 0 && countSell != 0 {
		select1, err1 := db.Query("SELECT price FROM DECISIONS WHERE Symbol=? && Buy=? ORDER BY id ASC LIMIT ?", symbol, true, countBuy)
		if err1 != nil {
			panic(err1.Error())
		}

		var counter int
		for select1.Next() {
			err1 := select1.Scan(&priceBuy[counter])
			if err1 != nil {
				panic(err1)
			}
			counter++
		}

		select2, err2 := db.Query("SELECT price FROM DECISIONS WHERE Symbol=? && Sell=? ORDER BY id ASC LIMIT ?", symbol, true, countSell)
		if err2 != nil {
			panic(err2.Error())
		}

		counter = 0
		for select2.Next() {
			err2 := select2.Scan(&priceSell[counter])
			if err2 != nil {
				panic(err2)
			}
			counter++
		}

		for i := 0; i < len(priceBuy); i++ {
			sumBuy += priceBuy[i]
		}

		for i := 0; i < len(priceSell); i++ {
			sumSell += priceSell[i]
		}
	}
	profitloss = sumSell - sumBuy

	select3, err3 := db.Query("SELECT id FROM ProfitLoss WHERE Symbol=?", symbol)
	if err3 != nil {
		panic(err3.Error())
	}
	var id int
	for select3.Next() {
		err3 := select3.Scan(&id)
		if err3 != nil {
			panic(err3.Error())
		}
	}
	if id == 0 {
		insert, err := db.Query("INSERT INTO ProfitLoss (Symbol,Total) VALUES (?,?)", symbol, profitloss)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
	} else {
		update, err := db.Query("UPDATE ProfitLoss SET Total=? WHERE id=?", profitloss, id)
		if err != nil {
			panic(err.Error())
		}
		defer update.Close()
	}

	return profitloss

}
