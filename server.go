// $env:NODE_OPTIONS="--openssl-legacy-provider"
// npm start
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/paymentintent"
)

func main(){
	stripe.Key = "" //put your stripe key

	http.HandleFunc("/test",test)
	http.HandleFunc("/CreatePaymentIntent",CreatePaymentIntent)
	err := http.ListenAndServe("localhost:4242",nil)
	if(err != nil){
		log.Fatal(err)
	}
}

func test(w http.ResponseWriter, r *http.Request){
	sendToClient := []byte("you're connect to server")
	_, err := w.Write(sendToClient)
	if(err != nil){
		log.Fatal(err)
	}
	fmt.Println(sendToClient)
	fmt.Println("somebody connet success!")
}
func CreatePaymentIntent(w http.ResponseWriter, r *http.Request){

	
	if r.Method != "POST" {
    http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    return
  }
//data structure pass by client
  var req struct {
    ProductId string `json:"product_id"`
  }
//decode data from client to json
  if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Printf("json.NewDecoder.Decode: %v", err)
    return
  }
//put json file in params
	 // Create a PaymentIntent with amount and currency
	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(CalculatePrice(req.ProductId)),
		Currency: stripe.String(string(stripe.CurrencyJPY)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	};
	//sent to stripe server which will return client_secreat  
		pi, err := paymentintent.New(params)
		log.Printf("pi.New: %v\n", pi.ClientSecret)
		log.Println(pi.Amount)
	
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("pi.New: %v", err)
			return
		}
		//sent client_secret to fronted(use writer)
		var buf bytes.Buffer
  if err := json.NewEncoder(&buf).Encode(pi.ClientSecret); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Printf("json.NewEncoder.Encode: %v", err)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  if _, err := io.Copy(w, &buf); err != nil {
    log.Printf("io.Copy: %v", err)
    return
  }
}

func CalculatePrice(pi string)int64{
	if pi == "Pants"{
		return 530
	}else{
		return 0
	}
	
}
		
