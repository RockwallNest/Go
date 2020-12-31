package main 

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "log"
)

type jsonData struct {
  ProductCode string  `json:"product_code"`
  Timestamp   string  `json:"timestamp"`
  Ask         float64 `json:"best_ask"`
  Bid         float64 `json:"best_bid"`
  Ltp         float64 `json:"ltp"`
}

func main() {
  /*
  url := "https://api.bitflyer.com/v1/ticker?product_code=ETH_JPY"
  
  resp, err := http.Get(url)
  if err != nil {
    log.Printf("ERROR: %s", err)
  }
  */
 
  url := "https://api.bitflyer.com/v1/ticker"

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Printf("ERROR: %s", err)
  }

  params := req.URL.Query()
  params.Add("product_code", "BTC_JPY")
  req.URL.RawQuery = params.Encode()

  client := http.Client{}

  resp, err := client.Do(req)
  if err != nil {
    log.Printf("ERROR: %s", err)
  }
  
  defer resp.Body.Close()
  
  byteArray, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Printf("ERROR: %s", err)
  }

  var d jsonData
  if err = json.Unmarshal(byteArray, &d); err != nil {
    log.Printf("ERROR: %s", err)
  }

  fmt.Printf(
    "product_code: %v\ntimestamp: %s\nask: %0.1f\nbid: %0.1f\nltp: %0.1f\n",
    d.ProductCode,
    d.Timestamp,
    d.Ask,
    d.Bid,
    d.Ltp)

}
