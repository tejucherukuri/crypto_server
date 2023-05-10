package main

import (
	"io/ioutil"
	"encoding/json"
    "fmt"
    "net/http"
)

type Crypto struct {
    ID           string `json:"id"`
	Name         string `json:"fullName"`
	Crypto       bool   `json:"crypto"`
	PayinEnabled bool   `json:"payinEnabled"`
	PayoutEnabled bool `json:"payoutEnabled"`
}


func main() {
	fmt.Println("Microservice API that uses the crypto API endpoint and returns the real-time crypto prices" )
	http.HandleFunc("/currency/", getSingleCrypto)
	http.HandleFunc("/currency/all", getAllCryptos)
	http.ListenAndServe(":8080", nil)
}

func getAllCryptos(w http.ResponseWriter, r *http.Request) {
	url := "https://api.hitbtc.com/api/2/public/currency"

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//TODO:Handle the error print and return
		return 
	}

	var response []Crypto
	err = json.Unmarshal(body, &response)
	json.NewEncoder(w).Encode(response)
}

func getSingleCrypto(w http.ResponseWriter, r *http.Request) {
    symbol := r.URL.Path[len("/currency/"):]
    url := fmt.Sprintf("https://api.hitbtc.com/api/2/public/currency/%s", symbol)

    resp, err := http.Get(url)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//TODO:Handle the error print and return
		return 
	}

    var currency Crypto
    err = json.Unmarshal(body, &currency)
	json.NewEncoder(w).Encode(currency)
}