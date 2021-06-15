package main

import (
	"db"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/robfig/cron.v2"
)

func init() { //log initiation
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	coins := []string{"BTC", "BCH", "ETH", "XRP", "LTC"}
	cr := cron.New()

	cr.AddFunc("0 */5 * * * *", func() { //5 minute job
		db.DBWrite(db.Call())
		log.Info("[Job 1]API CALL & DB Write every 5 minutes")
	})

	cr.AddFunc("20 0 */1 * * *", func() { //Writing Hi low and close prices into HighLow table
		db.DBWriteHiLo(coins)
		log.Info("[Job 2]Write hourly High Low price")
	})
	cr.AddFunc("50 0 */1 * * *", func() {
		db.Profit("BTC")
		db.Profit("BCH")
		db.Profit("ETH")
		db.Profit("LTC")
		db.Profit("XRP")
		log.Info("[Job 3]Calculate Profit-Loss")
	})

	cr.Start()
	printCronEntries(cr.Entries())
	runbackground()
}

func printCronEntries(cronEntries []cron.Entry) { // writing  Cron Info Logs
	log.Infof("Cron Info: %+v\n", cronEntries)
}

func runbackground() { // Spawn all you worker goroutines, and send a message to exit when you're done.
	exit := make(chan string)

	for {
		select {
		case <-exit:
			os.Exit(0)
		}
	}
}
