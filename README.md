## Unofficial KREDIVO SDK for Go Programming Language

### KREDIVO API DOCS (https://doc.kredivo.com)

[![GoDoc](https://godoc.org/github.com/Bhinneka/kredivo-go?status.svg)](https://godoc.org/github.com/Bhinneka/kredivo-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Bhinneka/kredivo-go)](https://goreportcard.com/report/github.com/Bhinneka/kredivo-go)
[![Build Status](https://travis-ci.org/Bhinneka/kredivo-go.svg?branch=master)](https://travis-ci.org/Bhinneka/kredivo-go)
[![codecov](https://codecov.io/gh/Bhinneka/kredivo-go/branch/master/graph/badge.svg)](https://codecov.io/gh/Bhinneka/kredivo-go)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/Bhinneka/kredivo-go/blob/master/LICENSE)

## Bhinneka :blue_heart: Golang

### Bopher (Bhinneka Gopher) ![alt text](./files/bhoper.png)

### Install
  ```shell
  go get github.com/Bhinneka/kredivo-go
  ```

### Configuration

  By Default KREDIVO Go using Sandbox Environment,
  you can change this config after calling KREDIVO Constructor

  ```go
  kred := kredivo.New("8tLHIx8V0N6KtnSpS9Nbd6zROFFJH7", "https://api.bhinneka.com/push_notif", "https://bhinneka.com", 8*time.Second)
  kred.Env = kredivo.Production.String()
  ```

### Simple Usage

  - Checkout Order Example

    ```go
    package main

    import (
    	"fmt"
    	"time"

    	"github.com/Bhinneka/kredivo-go"
    )

    //Checkout Order
    //in abstraction style
    //you should use KredivoService, instead of invoke directly to kredivo's Constructor
    //this function now easy to test
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

    //main function
    func main() {
    	fmt.Println("KREDIVO")

      //KREDIVO Constructor
      //Required Parameters:
      //1. Merchant Server Key
      //2. URI of merchant push-notification API (HTTP POST)
      //3. URI of your store page. Used on the settlement page. Kredivo’s server will pass some params to this uri for merchant’s server acknowledgement: order_id: Order Id given by Merchant; tr_id: Transaction Id given by Kredivo;tr_status: Transaction status of a transaction; sign_key: Signature key to validate if the notification is originated from Kredivo. Please contact us how to parse this signature_key by using your client key.
      //4. Http Request Timeout
    	kred := kredivo.New("8tLHIx8V0N6KtnSpS9Nbd6zROFFJH7", "https://api.bhinneka.com/push_notif", "https://bhinneka.com", 8*time.Second)

    	checkoutOrder(kred)
    }

    ```

## Test and Coverage

  - Unit Test
    ```shell
    make test
    ```

  - Coverage (Running Test and Display Coverage result)
    ```shell
    make cover
    ```
    
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/for-you.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/fuck-it-ship-it.svg)](https://forthebadge.com)

## Authors
  - Lone Wolf (https://github.com/wuriyanto48)
  - Dhiaz Aditya (https://github.com/dhiaz06)
