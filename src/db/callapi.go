package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mail"
	"net/http"
	"net/url"
	"time"
)

//CryptoData...
type CryptoData struct { // Cryptocurrency JSON type
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
	} `json:"status"`
	Data struct {
		BCH struct {
			ID                int         `json:"id"`
			Name              string      `json:"name"`
			Symbol            string      `json:"symbol"`
			Slug              string      `json:"slug"`
			NumMarketPairs    int         `json:"num_market_pairs"`
			DateAdded         time.Time   `json:"date_added"`
			Tags              []string    `json:"tags"`
			MaxSupply         int         `json:"max_supply"`
			CirculatingSupply float64     `json:"circulating_supply"`
			TotalSupply       float64     `json:"total_supply"`
			IsActive          int         `json:"is_active"`
			Platform          interface{} `json:"platform"`
			CmcRank           int         `json:"cmc_rank"`
			IsFiat            int         `json:"is_fiat"`
			LastUpdated       time.Time   `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"BCH"`
		BTC struct {
			ID                int         `json:"id"`
			Name              string      `json:"name"`
			Symbol            string      `json:"symbol"`
			Slug              string      `json:"slug"`
			NumMarketPairs    int         `json:"num_market_pairs"`
			DateAdded         time.Time   `json:"date_added"`
			Tags              []string    `json:"tags"`
			MaxSupply         int         `json:"max_supply"`
			CirculatingSupply int         `json:"circulating_supply"`
			TotalSupply       int         `json:"total_supply"`
			IsActive          int         `json:"is_active"`
			Platform          interface{} `json:"platform"`
			CmcRank           int         `json:"cmc_rank"`
			IsFiat            int         `json:"is_fiat"`
			LastUpdated       time.Time   `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"BTC"`
		ETH struct {
			ID                int         `json:"id"`
			Name              string      `json:"name"`
			Symbol            string      `json:"symbol"`
			Slug              string      `json:"slug"`
			NumMarketPairs    int         `json:"num_market_pairs"`
			DateAdded         time.Time   `json:"date_added"`
			Tags              []string    `json:"tags"`
			MaxSupply         interface{} `json:"max_supply"`
			CirculatingSupply float64     `json:"circulating_supply"`
			TotalSupply       float64     `json:"total_supply"`
			IsActive          int         `json:"is_active"`
			Platform          interface{} `json:"platform"`
			CmcRank           int         `json:"cmc_rank"`
			IsFiat            int         `json:"is_fiat"`
			LastUpdated       time.Time   `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"ETH"`
		LTC struct {
			ID                int         `json:"id"`
			Name              string      `json:"name"`
			Symbol            string      `json:"symbol"`
			Slug              string      `json:"slug"`
			NumMarketPairs    int         `json:"num_market_pairs"`
			DateAdded         time.Time   `json:"date_added"`
			Tags              []string    `json:"tags"`
			MaxSupply         int         `json:"max_supply"`
			CirculatingSupply float64     `json:"circulating_supply"`
			TotalSupply       float64     `json:"total_supply"`
			IsActive          int         `json:"is_active"`
			Platform          interface{} `json:"platform"`
			CmcRank           int         `json:"cmc_rank"`
			IsFiat            int         `json:"is_fiat"`
			LastUpdated       time.Time   `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"LTC"`
		XRP struct {
			ID                int           `json:"id"`
			Name              string        `json:"name"`
			Symbol            string        `json:"symbol"`
			Slug              string        `json:"slug"`
			NumMarketPairs    int           `json:"num_market_pairs"`
			DateAdded         time.Time     `json:"date_added"`
			Tags              []interface{} `json:"tags"`
			MaxSupply         int64         `json:"max_supply"`
			CirculatingSupply int64         `json:"circulating_supply"`
			TotalSupply       int64         `json:"total_supply"`
			IsActive          int           `json:"is_active"`
			Platform          interface{}   `json:"platform"`
			CmcRank           int           `json:"cmc_rank"`
			IsFiat            int           `json:"is_fiat"`
			LastUpdated       time.Time     `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"XRP"`
	} `json:"data"`
}

var Cryp CryptoData
var Resp *http.Response

func Call() CryptoData { //Calling API

	err1 := retry(5, 3*time.Second, func() (err error) {
		err1 := HandleRequest()
		return err1
	})
	if err1 != nil {
		log.Println(err1)
	}

	err2 := retry(5, 3*time.Second, func() (err error) {
		err1 := HandleResponse()
		if err1 != nil {
			err := HandleRequest()
			if err != nil {
				log.Println(err.Error())
			}
		}
		return err1
	})
	if err2 != nil {
		log.Println(err2)
	}

	return Cryp
}

func HandleRequest() error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Println(err)
		mail.ErrorMail(err.Error())
	}

	q := url.Values{}
	q.Add("symbol", "BTC,ETH,XRP,BCH,LTC")
	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "//***API Key***//")
	req.URL.RawQuery = q.Encode()

	Resp, err = client.Do(req)
	if err != nil {
		log.Println("Error sending request to server")
		mail.ErrorMail(err.Error())
	}

	return err
}

func HandleResponse() error {
	var err error
	respBody, _ := ioutil.ReadAll(Resp.Body)
	if err = json.Unmarshal(respBody, &Cryp); err != nil {
		log.Println("Error in response body")
		mail.ErrorMail(err.Error())
	}

	return err

}

func retry(attempts int, sleep time.Duration, f func() error) (err error) {
	for i := 0; ; i++ {
		err = f()
		if err == nil {
			return
		}

		if i >= (attempts - 1) {
			break
		}

		time.Sleep(sleep)

		log.Println("retrying after error:", err)
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
