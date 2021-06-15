package main

import (
	"dependency"
	"os"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"gopkg.in/robfig/cron.v2"
)

func init() { //log initiation
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	cr := cron.New()

	_, err1 := cr.AddFunc("50 */5 * * * *", func() { //5 minute job
		dependency.StopLoss("BTC", 0.5)
		dependency.StopLoss("BCH", 2)
		dependency.StopLoss("LTC", 2)
		dependency.StopLoss("ETH", 2)
		dependency.StopLoss("XRP", 2)

		dependency.MovingStop("BTC", 20)
		dependency.MovingStop("BCH", 2)
		dependency.MovingStop("LTC", 2)
		dependency.MovingStop("ETH", 2)
		dependency.MovingStop("XRP", 2)
		log.Info("[Job 5]StopLoss & MovingStop")
	})
	if err1 != nil {
		panic(err1)
	}

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
