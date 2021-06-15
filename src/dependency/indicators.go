package dependency

import (
	"db"
	"math"

	_ "github.com/go-sql-driver/mysql"
)

func SMA(amount int, symbol string, dbwrite bool, index int) float64 {
	var sma, sum float64
	price, rowid, symbol := db.DBReadHiLo(amount, symbol)

	db := db.DBconn()
	defer db.Close()

	for i := 0; i < len(price); i++ {
		sum += price[2][i]
	}
	sma = sum / float64(len(price))

	if dbwrite {
		if index == 0 {
			select1, err1 := db.Query("SELECT id FROM METRICS WHERE RowID=? AND Symbol=?", rowid, symbol)
			if err1 != nil {
				panic(err1.Error())
			}
			var id int
			for select1.Next() {
				err2 := select1.Scan(&id)
				if err2 != nil {
					panic(err2.Error())
				}
			}

			if id == 0 {
				insert, err := db.Query("INSERT INTO METRICS (RowID,Symbol,SMA) VALUES (?,?,?)", rowid, symbol, sma)
				if err != nil {
					panic(err.Error())
				}
				defer insert.Close()
			} else {
				update, err := db.Query("UPDATE METRICS SET SMA=? WHERE id=?", sma, id)
				if err != nil {
					panic(err.Error())
				}
				defer update.Close()
			}
		}
		if index == 1 {
			select1, err1 := db.Query("SELECT id FROM METRICS WHERE RowID=? AND Symbol=?", rowid, symbol)
			if err1 != nil {
				panic(err1.Error())
			}
			var id int
			for select1.Next() {
				err2 := select1.Scan(&id)
				if err2 != nil {
					panic(err2.Error())
				}
			}

			if id == 0 {
				insert, err := db.Query("INSERT INTO METRICS (RowID,Symbol,SMA2) VALUES (?,?,?)", rowid, symbol, sma)
				if err != nil {
					panic(err.Error())
				}
				defer insert.Close()
			} else {
				update, err := db.Query("UPDATE METRICS SET SMA2=? WHERE id=?", sma, id)
				if err != nil {
					panic(err.Error())
				}
				defer update.Close()
			}
		}
	}
	return sma

}

func SMAP(price []float64) float64 {
	var sma, sum float64

	for i := 0; i < len(price); i++ {
		sum += price[i]
	}
	sma = sum / float64(len(price))
	return sma

}

func StdDev(amount int, symbol string) float64 {
	var stddev, sum float64
	price, _, _ := db.DBReadHiLo(amount, symbol)
	pr := make([]float64, len(price[2]))

	db := db.DBconn()
	defer db.Close()

	for i := 0; i < len(price); i++ {
		pr[i] = price[2][i] - SMAP(price[2])
		sum += pr[i] * pr[i]
	}
	sum = sum / float64(len(price))
	stddev = math.Sqrt(sum)

	return stddev
}

func StdDevP(price []float64) float64 {
	var stddev, sum float64
	pr := make([]float64, len(price))

	for i := 0; i < len(price); i++ {
		pr[i] = price[i] - SMAP(price)
		sum += pr[i] * pr[i]
	}
	sum = sum / float64(len(price))
	stddev = math.Sqrt(sum)

	return stddev
}

func Bollinger(amount int, symbol string, stdevmultip float64, dbwrite bool) [3]float64 {
	var upperband, middleband, lowerband float64
	var bollingervar [3]float64
	price, rowid, symbol := db.DBReadHiLo(amount, symbol)
	tp := make([]float64, amount)

	db := db.DBconn()
	defer db.Close()

	for i := 0; i < amount; i++ {
		tp[i] = (price[0][i] + price[1][i] + price[2][i]) / 3
	}

	middleband = SMAP(tp)
	upperband = SMAP(tp) + (StdDevP(tp) * stdevmultip)
	lowerband = SMAP(tp) - (StdDevP(tp) * stdevmultip)
	bollingervar[0] = upperband
	bollingervar[1] = middleband
	bollingervar[2] = lowerband

	if dbwrite {
		select1, err1 := db.Query("SELECT id FROM METRICS WHERE RowID=? AND Symbol=?", rowid, symbol)
		if err1 != nil {
			panic(err1.Error())
		}
		var id int
		for select1.Next() {
			err2 := select1.Scan(&id)
			if err2 != nil {
				panic(err2.Error())
			}
		}
		if id == 0 {
			insert, err := db.Query("INSERT INTO METRICS (RowID,Symbol,BollingerUp,BollingerMid,BollingerLow) VALUES (?,?,?,?,?)", rowid, symbol, bollingervar[0], bollingervar[1], bollingervar[2])
			if err != nil {
				panic(err.Error())
			}
			defer insert.Close()
		} else {
			update, err := db.Query("UPDATE METRICS SET BollingerUp=?,BollingerMid=?,BollingerLow=? WHERE id=?", bollingervar[0], bollingervar[1], bollingervar[2], id)
			if err != nil {
				panic(err.Error())
			}
			defer update.Close()
		}
	}
	return bollingervar
}

func STOC(amount int, slow int, sbool bool, symbol string, dbwrite bool) (float64, float64) {
	var min, max, stoc, stocs float64
	price, rowid, symbol := db.DBReadHiLo(amount, symbol)

	db := db.DBconn()
	defer db.Close()

	min = price[1][0]
	for _, price := range price[1] {
		if price < min {
			min = price
		}
	}

	max = price[0][0]
	for _, price := range price[0] {
		if price > max {
			max = price
		}
	}

	stoc = ((price[2][0] - min) / (max - min)) * 100

	if dbwrite {
		select1, err1 := db.Query("SELECT id FROM METRICS WHERE RowID=? AND Symbol=?", rowid, symbol)
		if err1 != nil {
			panic(err1.Error())
		}
		var id int
		for select1.Next() {
			err2 := select1.Scan(&id)
			if err2 != nil {
				panic(err2.Error())
			}
		}
		if id == 0 {
			insert, err := db.Query("INSERT INTO METRICS (RowID,Symbol,STOC) VALUES (?,?,?)", rowid, symbol, stoc)
			if err != nil {
				panic(err.Error())
			}
			defer insert.Close()
		} else {
			update, err := db.Query("UPDATE METRICS SET STOC=? WHERE id=?", stoc, id)
			if err != nil {
				panic(err.Error())
			}
			defer update.Close()
		}
	}
	if sbool {
		stocs = STOCS(slow, symbol, dbwrite)
	}
	return stoc, stocs
}

func STOCS(amount int, symbol string, dbwrite bool) float64 {
	var stocs float64
	stoc := make([]float64, amount)
	id := make([]int, amount)

	db := db.DBconn()
	defer db.Close()

	select1, err1 := db.Query("SELECT id,STOC FROM METRICS WHERE Symbol=? ORDER BY id DESC LIMIT ? ", symbol, amount)
	if err1 != nil {
		panic(err1.Error())
	}

	counter := 0
	for select1.Next() {
		err2 := select1.Scan(&id[counter], &stoc[counter])
		if err2 != nil {
			panic(err2.Error())
		}
		counter++
	}

	for i := 0; i < len(stoc); i++ {
		stocs += stoc[i]
	}

	stocs = stocs / float64(amount)

	if dbwrite {

		update, err := db.Query("UPDATE METRICS SET STOCS=? WHERE id=?", stocs, id[0])
		if err != nil {
			panic(err.Error())
		}
		defer update.Close()

	}

	return stocs
}

func STOCP(price [][]float64) float64 {
	var min, max, stoc float64

	min = price[1][0]
	for _, price := range price[1] {
		if price < min {
			min = price
		}
	}

	max = price[0][0]
	for _, price := range price[0] {
		if price > max {
			max = price
		}
	}

	stoc = ((price[2][0] - min) / (max - min)) * 100

	return stoc
}

func STOCD(amount int, slow int, sbool bool, symbol string, dbwrite bool) (float64, float64) {
	var stocd, stocds float64
	price, rowid, symbol := db.DBReadHiLo(amount*3, symbol)
	stoc := make([]float64, len(price))

	db := db.DBconn()
	defer db.Close()

	price1 := make([][]float64, len(price), amount)
	price2 := make([][]float64, len(price), amount)
	price3 := make([][]float64, len(price), amount)

	for j := 0; j < len(price); j++ {
		price1[j] = price[j][0:amount]
		price2[j] = price[j][amount:(2 * amount)]
		price3[j] = price[j][(2 * amount):(3 * amount)]
	}

	stoc[0] = STOCP(price1)
	stoc[1] = STOCP(price2)
	stoc[2] = STOCP(price3)
	stocd = SMAP(stoc)

	if dbwrite {
		select1, err1 := db.Query("SELECT id FROM METRICS WHERE RowID=? AND Symbol=?", rowid, symbol)
		if err1 != nil {
			panic(err1.Error())
		}
		var id int
		for select1.Next() {
			err2 := select1.Scan(&id)
			if err2 != nil {
				panic(err2.Error())
			}
		}
		if id == 0 {
			insert, err := db.Query("INSERT INTO METRICS (RowID,Symbol,STOCD) VALUES (?,?,?)", rowid, symbol, stocd)
			if err != nil {
				panic(err.Error())
			}
			defer insert.Close()
		} else {
			update, err := db.Query("UPDATE METRICS SET STOCD=? WHERE id=?", stocd, id)
			if err != nil {
				panic(err.Error())
			}
			defer update.Close()
		}
	}
	if sbool {
		stocds = STOCDS(slow, symbol, dbwrite)
	}
	return stocd, stocds
}

func STOCDS(amount int, symbol string, dbwrite bool) float64 {
	var stocds float64
	stocd := make([]float64, amount)
	id := make([]int, amount)

	db := db.DBconn()
	defer db.Close()

	select1, err1 := db.Query("SELECT id,STOCD FROM METRICS WHERE Symbol=? ORDER BY id DESC LIMIT ? ", symbol, amount)
	if err1 != nil {
		panic(err1.Error())
	}

	counter := 0
	for select1.Next() {
		err2 := select1.Scan(&id[counter], &stocd[counter])
		if err2 != nil {
			panic(err2.Error())
		}
		counter++
	}

	for i := 0; i < len(stocd); i++ {
		stocds += stocd[i]
	}

	stocds = stocds / float64(amount)

	if dbwrite {

		update, err := db.Query("UPDATE METRICS SET STOCDS=? WHERE id=?", stocds, id[0])
		if err != nil {
			panic(err.Error())
		}
		defer update.Close()

	}

	return stocds
}

func EMA(amount int, symbol string, dbwrite bool, index int) float64 {
	ema := make([]float64, amount)
	price, rowid, symbol := db.DBReadHiLo(amount*2, symbol)
	pricecorr := make([]float64, amount*2)
	multiplier := float64(2 / (float64(amount) + 1))

	db := db.DBconn()
	defer db.Close()

	for i := 0; i < len(price[2]); i++ {
		pricecorr[i] = price[2][len(price[2])-i-1]
	}

	temp := make([]float64, amount)
	for i := 0; i < len(temp); i++ {
		temp[i] = pricecorr[i]
	}

	ema[0] = SMAP(temp)
	for i := 1; i < len(ema); i++ {
		ema[i] = (pricecorr[amount+i]-ema[i-1])*multiplier + ema[i-1]
	}

	if dbwrite {
		if index == 0 {
			select1, err1 := db.Query("SELECT id FROM METRICS WHERE RowID=? AND Symbol=?", rowid, symbol)
			if err1 != nil {
				panic(err1.Error())
			}
			var id int
			for select1.Next() {
				err2 := select1.Scan(&id)
				if err2 != nil {
					panic(err2.Error())
				}
			}
			if id == 0 {
				insert, err := db.Query("INSERT INTO METRICS (RowID,Symbol,EMA) VALUES (?,?,?)", rowid, symbol, ema[len(ema)-1])
				if err != nil {
					panic(err.Error())
				}
				defer insert.Close()
			} else {
				update, err := db.Query("UPDATE METRICS SET EMA=? WHERE id=?", ema[len(ema)-1], id)
				if err != nil {
					panic(err.Error())
				}
				defer update.Close()
			}
		}
		if index == 1 {
			select1, err1 := db.Query("SELECT id FROM METRICS WHERE RowID=? AND Symbol=?", rowid, symbol)
			if err1 != nil {
				panic(err1.Error())
			}
			var id int
			for select1.Next() {
				err2 := select1.Scan(&id)
				if err2 != nil {
					panic(err2.Error())
				}
			}
			if id == 0 {
				insert, err := db.Query("INSERT INTO METRICS (RowID,Symbol,EMA2) VALUES (?,?,?)", rowid, symbol, ema[len(ema)-1])
				if err != nil {
					panic(err.Error())
				}
				defer insert.Close()
			} else {
				update, err := db.Query("UPDATE METRICS SET EMA2=? WHERE id=?", ema[len(ema)-1], id)
				if err != nil {
					panic(err.Error())
				}
				defer update.Close()
			}
		}
	}
	return ema[len(ema)-1]
}

func TRIX(amount int, symbol string, dbwrite bool) float64 {
	var price [][]float64
	var trix float64
	var rowid int

	pricecorrhi := make([]float64, amount*3)
	pricecorrlo := make([]float64, amount*3)
	pricecorrcl := make([]float64, amount*3)

	price, rowid, _ = db.DBReadHiLo(amount*3, symbol)

	for i := 0; i < len(price[0]); i++ {
		pricecorrhi[i] = price[0][len(price[0])-i-1]
	}
	for i := 0; i < len(price[1]); i++ {
		pricecorrlo[i] = price[1][len(price[1])-i-1]
	}
	for i := 0; i < len(price[2]); i++ {
		pricecorrcl[i] = price[2][len(price[2])-i-1]
	}

	temp1 := make([]float64, len(pricecorrcl)/3)
	temp2 := make([]float64, len(pricecorrcl)*2/3)
	sma1 := make([]float64, (len(pricecorrcl)*2/3)+1)
	sma2 := make([]float64, (len(pricecorrcl)*2/3)+1)
	sma3 := make([]float64, (len(pricecorrcl)*2/3)+1)

	for i := 0; i < len(temp1); i++ {
		temp1[i] = pricecorrcl[i]
	}
	for i := 0; i < len(temp2); i++ {
		temp2[i] = pricecorrcl[len(temp1)+i]
	}

	sma1[0] = SMAP(temp1)
	for i := 1; i < len(sma1); i++ {
		sma1[i] = (sma1[i-1] + temp2[i-1]) / 2
	}

	sma2[0] = SMAP(sma1)
	for i := 1; i < len(sma2); i++ {
		sma2[i] = (sma2[i-1] + sma1[i]) / 2
	}

	sma3[0] = SMAP(sma2)
	for i := 1; i < len(sma3); i++ {
		sma3[i] = (sma3[i-1] + sma2[i]) / 2
	}

	trix = (sma3[len(sma3)-1] - sma3[len(sma3)-2]) / sma3[len(sma3)-2]
	db := db.DBconn()
	defer db.Close()

	if dbwrite {
		select1, err1 := db.Query("SELECT id FROM METRICS WHERE RowID=? AND Symbol=?", rowid, symbol)
		if err1 != nil {
			panic(err1.Error())
		}
		var id int
		for select1.Next() {
			err2 := select1.Scan(&id)
			if err2 != nil {
				panic(err2.Error())
			}
		}
		if id == 0 {
			insert, err := db.Query("INSERT INTO METRICS (RowID,Symbol,TRIX) VALUES (?,?,?)", rowid, symbol, trix)
			if err != nil {
				panic(err.Error())
			}
			defer insert.Close()
		} else {
			update, err := db.Query("UPDATE METRICS SET TRIX=? WHERE RowId=? AND Symbol=?", trix, rowid, symbol)
			if err != nil {
				panic(err.Error())
			}
			defer update.Close()
		}
	}
	return trix
}
