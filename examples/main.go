package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Bhinneka/kredivo-go"
)

var paymentRequestData = []byte(`{
    "server_key":"8tLHIx8V0N6KtnSpS9Nbd6zROFFJH7",
    "amount":2400000,
    "items": [
        {
            "id":"BB12345678",
            "name":"iPhone 5S",
            "price":2400000,
            "url":"http://merchant.com/cellphones/iphone5s_64g",
            "type":"Smartphone",
            "quantity":1
        }
    ]
  }`)

var cancelRequestData = []byte(`{
    "server_key":"MEJ4FLRc74UU64cxCF8Z3HYSpPctD7",
    "order_id":"8192K383",
    "transaction_id":"f6e33997-9ea2-426c-991b-9421165b9343",
    "cancellation_reason":"Out of stock",
    "cancelled_by":"Althea",
    "cancellation_date":"1501844808"

}`)

var transactionStatusRequestData = []byte(`{
	"server_key":"MEJ4FLRc74UU64cxCF8Z3HYSpPctD7",
	"order_id":"8192K383"
}`)

func checkoutOrder(kred kredivo.KredivoService) {

	var order kredivo.Order

	order.PaymentType = "30_days"
	order.TransactionDetails = kredivo.TransactionDetails{
		OrderID: "782930303",
		Amount:  500000,
		Items: []kredivo.Item{kredivo.Item{ID: "AZ14565678",
			Name:       "Hailee Sneakers Blink Silver",
			Price:      500000,
			URL:        "http://merchant.com/fashion/shoes/sneakers-blink-shoes",
			Type:       "Sneakers",
			Quantity:   1,
			ParentType: "SELLER",
			ParentID:   "SELLER456",
		},
		},
	}
	order.CustomerDetails = &kredivo.CustomerDetails{
		FirstName: "Wuriyanto",
		LastName:  "Musobar",
		Email:     "wuriyanto48@yahoo.co.id",
		Phone:     "02188888",
	}

	order.BillingAddress = &kredivo.Address{
		FirstName:   "Wuriyanto",
		LastName:    "Musobar",
		Address:     "Jalan Teknologi Indonesia No. 25",
		City:        "Jakarta",
		PostalCode:  "12960",
		Phone:       "081513114262",
		CountryCode: "IDN",
	}

	order.ShippingAddress = kredivo.Address{
		FirstName:   "Wuriyanto",
		LastName:    "Musobar",
		Address:     "Jalan Teknologi Indonesia No. 25",
		City:        "Jakarta",
		PostalCode:  "12960",
		Phone:       "081513114262",
		CountryCode: "IDN",
	}

	order.Sellers = []kredivo.Seller{kredivo.Seller{
		ID:    "BH111",
		Name:  "Bhinneka.com",
		Email: "bhinneka@bhinneka.com",
		Address: kredivo.Address{
			FirstName:   "Bhinneka",
			LastName:    "Mentaro",
			Address:     "Jalan Gunung Sahari Indonesia No. 25",
			City:        "Jakarta",
			PostalCode:  "12960",
			Phone:       "081513114262",
			CountryCode: "IDN",
		},
	}}

	result := kred.Checkout(&order)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}

	//result.Result is an interface
	//you should assert to specific type

	checkoutResponse, ok := result.Result.(kredivo.CheckoutResponse)

	if !ok {
		fmt.Println("Result is not Checkout Response")
	}

	fmt.Println(checkoutResponse)
}

func getPayments() {
	kred := kredivo.New("8tLHIx8V0N6KtnSpS9Nbd6zROFFJH7", "https://api.merchant.com/push.php", "https://merchant.com", 8*time.Second)

	var paymentRequest kredivo.PaymentRequest

	err := json.Unmarshal(paymentRequestData, &paymentRequest)

	if err != nil {
		fmt.Println("errrrrr " + err.Error())
	}

	result := kred.GetPayments(&paymentRequest)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}

	paymentRes, ok := result.Result.(kredivo.PaymentResponse)

	if !ok {
		fmt.Println("result is not payment response")
	}

	for _, v := range paymentRes.Payments {
		fmt.Println(v.Amount)
	}
}

func confirm() {
	kred := kredivo.New("8tLHIx8V0N6KtnSpS9Nbd6zROFFJH7", "https://api.merchant.com/push.php", "https://merchant.com", 8*time.Second)

	var confRequest kredivo.ConfirmRequest
	confRequest.TransactionID = "f6e33997-9ea2-426c-991b-9421165b9343"
	confRequest.SignatureKey = "YmQBYtAODqlWkmVrkNY%2BRtHclC9yHMDsKwAJ%2BG4n%2BQ1m1DlgqtIfsjjQrUFEIs%2BnWlHhJahCmJuj%2FSOJ8YmPPuX9iKoquvfJl5n0V57Cess%3D"

	result := kred.Confirm(&confRequest)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}

	confirmRes, ok := result.Result.(kredivo.ConfirmationResponse)

	if !ok {
		fmt.Println("result is not confirmation response")
	}

	fmt.Println(confirmRes)
}

func cancel() {
	kred := kredivo.New("8tLHIx8V0N6KtnSpS9Nbd6zROFFJH7", "https://api.merchant.com/push.php", "https://merchant.com", 8*time.Second)

	var cancelRequest kredivo.CancelRequest

	err := json.Unmarshal(cancelRequestData, &cancelRequest)

	if err != nil {
		fmt.Println("errrrrr " + err.Error())
	}

	result := kred.Cancel(&cancelRequest)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}

	cancelRes, ok := result.Result.(kredivo.CancelResponse)

	if !ok {
		fmt.Println("result is not cancel response")
	}

	fmt.Println(cancelRes)

}

func GetTransactionStatus() {
	kred := kredivo.New("8tLHIx8V0N6KtnSpS9Nbd6zROFFJH7", "https://api.merchant.com/push.php", "https://merchant.com", 8*time.Second)

	var transactionStatusRequest kredivo.TransactionStatusRequest

	err := json.Unmarshal(transactionStatusRequestData, &transactionStatusRequest)

	if err != nil {
		fmt.Println("errrrrr " + err.Error())
	}

	result := kred.TransactionStatus(&transactionStatusRequest)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}

	trxStatusRes, ok := result.Result.(kredivo.TransactionStatusResponse)

	if !ok {
		fmt.Println("result is not transaction status response")
	}

	fmt.Println(trxStatusRes)

}

func main() {
	fmt.Println("KREDIVO")

	kred := kredivo.New("8tLHIx8V0N6KtnSpS9Nbd6zROFFJH7", "https://api.bhinneka.com/push_notif", "https://bhinneka.com", 8*time.Second)

	checkoutOrder(kred)
}
