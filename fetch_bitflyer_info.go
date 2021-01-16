package main

import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
)

type jsonData1 struct {
        Ticker string `json`
}

type jsonData struct {
        ProductCode string  `json:"product_code"`
        Timestamp   string  `json:"timestamp"`
        TickId      float64 `json:"tick_id"`
        Ask         float64 `json:"best_ask"`
        Bid         float64 `json:"best_bid"`
        AskSize     float64 `json:"best_ask_size"`
        BidSize     float64 `json:"best_bid_size"`
        AskDepth    float64 `json:"total_ask_depth"`
        BidDepth    float64 `json:"total_bid_depth"`
        MAskSize    float64 `json:"market_ask_size"`
        MBidSize    float64 `json:"market_bid_size"`
        Ltp         float64 `json:"ltp"`
        Volume      float64 `json:"volume"`
        VByProduct  float64 `json:"volume_by_product"`
}

func main() {

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
                "product_code: %v\ntimestamp: %s\ntick_id: %0.1f\nask: %0.1f\nbid: %0.1f\nask_size: %0.6f\nbid_size: %0.6f\nmarket_ask_size: %0.1f\nmarket_bid_size: %0.1f\ntotal_ask_depth: %0.6f\ntotal_bid_depth: %0.6f\nltp: %0.1f\nvolume: %0.6f\nvolume_by_product: %0.6f\n",
                d.ProductCode,
                d.Timestamp,
                d.TickId,
                d.Ask,
                d.Bid,
                d.AskSize,
                d.BidSize,
                d.AskDepth,
                d.BidDepth,
                d.MAskSize,
                d.MBidSize,
                d.Ltp,
                d.Volume,
                d.VByProduct,
        )

}
