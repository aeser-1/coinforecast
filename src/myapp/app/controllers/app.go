package controllers

import (
	"db"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

type Coin struct {
	CoinID int     `json:"coinID"`
	Name   string  ` json:"name" xml:"foo" `
	Symbol string  ` json:"sym" xml:"foo" `
	Price  float64 ` json:"price" xml:"bar" `
}

type Decision struct {
	Symbol   string ` json:"sym" xml:"foo" `
	Decision string ` json:"dec" xml:"foo" `
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Test() revel.Result {
	// greeting := "Aloha World"
	current := db.DBRead(1, "BTC")
	coin := Coin{CoinID: current.Rowid[0], Name: "Bitcoin", Symbol: current.Symbol, Price: current.Price[0]}
	current2 := db.DBRead(1, "XRP")
	coin2 := Coin{CoinID: current2.Rowid[0], Name: "Ripple", Symbol: current2.Symbol, Price: current2.Price[0]}
	current3 := db.DBRead(1, "BCH")
	coin3 := Coin{CoinID: current3.Rowid[0], Name: "BitCoin Cash", Symbol: current3.Symbol, Price: current3.Price[0]}
	current4 := db.DBRead(1, "LTC")
	coin4 := Coin{CoinID: current4.Rowid[0], Name: "Lite Coin", Symbol: current4.Symbol, Price: current4.Price[0]}
	current5 := db.DBRead(1, "ETH")
	coin5 := Coin{CoinID: current5.Rowid[0], Name: "Etherium", Symbol: current5.Symbol, Price: current5.Price[0]}
	data := [5]Coin{coin, coin2, coin3, coin4, coin5}
	//data = append(data, coin2)

	return c.RenderJSON(data)
	// return c.RenderJSON(greeting)
}

func (c App) Decision() revel.Result {
	// greeting := "Aloha World"
	curr := db.DBReadDecision("BTC")
	dec := Decision{Symbol: curr.Symbol, Decision: curr.Decision}
	curr2 := db.DBReadDecision("XRP")
	dec2 := Decision{Symbol: curr2.Symbol, Decision: curr2.Decision}
	curr3 := db.DBReadDecision("BCH")
	dec3 := Decision{Symbol: curr3.Symbol, Decision: curr3.Decision}
	curr4 := db.DBReadDecision("LTC")
	dec4 := Decision{Symbol: curr4.Symbol, Decision: curr4.Decision}
	curr5 := db.DBReadDecision("ETH")
	dec5 := Decision{Symbol: curr5.Symbol, Decision: curr5.Decision}
	dcdata := [5]Decision{dec, dec2, dec3, dec4, dec5}
	//data = append(data, coin2)

	return c.RenderJSON(dcdata)
	// return c.RenderJSON(greeting)
}
