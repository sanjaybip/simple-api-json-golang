package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)
type data struct {
	Endpoint string  `json:"endpoint"`
	Quotes []map[string]interface{} `json:"quotes"`
	Requested_time string `json:"requested_time"`
	Timestamp int32 `json:"timestamp"`
}
func main() {
	curreinces := "EURUSD,GBPUSD"
	
	envErr := godotenv.Load(".env")
	api_key := os.Getenv("TRADERMADE_API_KEY")

	if envErr != nil {
		log.Fatalf("Error loading .env file")
	}

	url := "https://marketdata.tradermade.com/api/v1/live?currency=" + curreinces + "&api_key=" + api_key
	resp, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Print(string(body))

	data_obj := data{}

	jsonErr := json.Unmarshal(body, &data_obj)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println("endpoint", data_obj.Endpoint, "requested time", data_obj.Requested_time, "timestamp", data_obj.Timestamp)

	for key, value := range data_obj.Quotes {
		fmt.Println(key)
		fmt.Println("symbol", value["base_currency"], value["quote_currency"], "bid", value["bid"], "mid",  value["mid"], "ask", value["ask"])
	}
}