package main

import (
	"os"
	"tradealgo"

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
	
	//You should create a cron job for any strategy in tradealgo.
	_, err := cr.AddFunc("25 0 */1 * * *", func() { //5 minute job
		tradealgo.BTCAlgo()
		log.Info("[Job 1]Trade algo BTC")
	})
	if err != nil {
		panic(err.Error())
	}

	_, err2 := cr.AddFunc("35 0 */1 * * *", func() { //5 minute job
		tradealgo.ETHAlgo()
		log.Info("[Job 3]Trade algo ETH")
	})
	if err2 != nil {
		panic(err2)
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
