package kredivo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestKredivo(t *testing.T) {

	//Order request payload
	orderData := []byte(`{
      "server_key":"8tLHIx8V0N6KtnSpS9Nbd6zROFFJH7",
      "payment_type":"30_days",
      "transaction_details": {
          "amount":9505000,
          "order_id":"736373K383",
          "items": [
              {
                  "id":"BB12345678",
                  "name":"iPhone 5S",
                  "price":6000000,
                  "type":"Smartphone",
                  "url":"http://merchant.com/cellphones/iphone5s_64g",
                  "quantity":1
              },
              {
                  "id":"AZ14565678",
                  "name":"Hailee Sneakers Blink Silver",
                  "price":250000,
                  "url":"http://merchant.com/fashion/shoes/sneakers-blink-shoes",
                  "type":"Sneakers",
                  "quantity":2,
                  "parent_type":"SELLER",
                  "parent_id":"SELLER456"
              },
              {
                  "id":"taxfee",
                  "name":"Tax Fee",
                  "price":1000,
                  "quantity":1
              },
              {
                  "id":"shippingfee",
                  "name":"Shipping Fee",
                  "price":90000,
                  "quantity":1,
                  "parent_type":"SELLER",
                  "parent_id":"SELLER456"
              },
              {
                  "id":"discount",
                  "name":"Discount",
                  "price":50000,
                  "quantity":1
              }
          ]
      },
      "sellers":[
          {
              "id":"SELLER123",
              "name":"Sunrise",
              "email": "sunrise@gmail.com",
              "address" : {
                  "first_name":"Irfan",
                  "last_name":"Sutandro",
                  "address":"Jalan Tentara Pelajar no 49",
                  "city":"Jakarta Utara",
                  "postal_code":"12960",
                  "phone":"08123456789",
                  "country_code":"IDN"
              }
          },
          {
              "id":"SELLER456",
              "name":"Toko Bagus",
              "email": "tokobagus@gmail.com",
              "address" : {
                  "first_name":"Toto",
                  "last_name":"Wahyuni",
                  "address":"Jalan Krici raya IX",
                  "city":"Jakarta Selatan",
                  "postal_code":"12960",
                  "phone":"08123456789",
                  "country_code":"IDN"
              }
          }
      ],
      "customer_details":{
          "first_name":"Oemang",
          "last_name":"Tandra",
          "email":"alie@satuduatiga.com",
          "phone":"081513114262"
      },
      "billing_address": {
          "first_name":"Oemang",
          "last_name":"Tandra",
          "address":"Jalan Teknologi Indonesia No. 25",
          "city":"Jakarta",
          "postal_code":"12960",
          "phone":"081513114262",
          "country_code":"IDN"
      },
      "shipping_address": {
          "first_name":"Oemang",
          "last_name":"Tandra",
          "address":"Jalan Teknologi Indonesia No. 25",
          "city":"Jakarta",
          "postal_code":"12960",
          "phone":"081513114262",
          "country_code":"IDN"
      },
      "push_uri":"https://api.merchant.com/push.php",
      "back_to_store_uri":"https://merchant.com"
  }`)

	//Payment Request Payload
	paymentRequestData := []byte(`{
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

	//Cancel Request Payload
	cancelRequestData := []byte(`{
	      "server_key":"MEJ4FLRc74UU64cxCF8Z3HYSpPctD7",
	      "order_id":"8192K383",
	      "transaction_id":"f6e33997-9ea2-426c-991b-9421165b9343",
	      "cancellation_reason":"Out of stock",
	      "cancelled_by":"Althea",
	      "cancellation_date":"1501844808"

	  }`)

	//partialCancelData
	partialCancelData := []byte(`{
      "server_key":"8tLHIx8V0N6KtnSpS9Nbd6zROFFJH7",
      "transaction_details": {
       "order_id":"KD14721",
       "amount":1500100,
       "items": [
          {
           "id":"11",
           "name":"Mesin Cuci",
           "cancellation_reason":"Out of stock"
          }
        ]
       },
      "cancelled_by":"Althea",
      "cancellation_date":"1501846306"
      }`)

	t.Run("Test Kredivo Creation", func(t *testing.T) {

		//construct kredivo
		//and add required parameters
		kred := New("123456", "bhinneka.com/kredivo/notif", "bhinneka.com", 8*time.Second)

		if kred == nil {
			t.Error("Cannot Call Kredivo Constructor")
		}
	})

	t.Run("Test Kredivo Checkout", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method != "POST" {
				t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
			}

			if r.URL.EscapedPath() != "/checkout_url" {
				t.Errorf("Expected request to ‘/checkout_url, got ‘%s’", r.URL.EscapedPath())
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
        "status":"OK",
        "message":"Message if any",
        "redirect_url":"https://sandbox.kredivo.com/kredivo/v2/signin?tk=XXX"
      }`))

		}))

		//close server
		defer ts.Close()

		//construct kredivo
		kred := New("123456", "bhinneka.com/kredivo/notif", "bhinneka.com", 8*time.Second)

		var order Order

		err := json.Unmarshal(orderData, &order)

		if err != nil {
			t.Error("Cannot Unmarshal Order JSON data")
		}

		kred.Env = ts.URL
		checkoutResult := kred.Checkout(&order)

		if checkoutResult.Error != nil {
			t.Errorf("Checkout() returned an error: %s", checkoutResult.Error)
		}

	})

	t.Run("Test Kredivo Checkout Error Empty Result", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method != "POST" {
				t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
			}

			if r.URL.EscapedPath() != "/checkout_url" {
				t.Errorf("Expected request to ‘/checkout_url, got ‘%s’", r.URL.EscapedPath())
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`error`))

		}))

		//close server
		defer ts.Close()

		//construct kredivo
		kred := New("123456", "bhinneka.com/kredivo/notif", "bhinneka.com", 8*time.Second)

		var order Order

		err := json.Unmarshal([]byte("order error"), &order)

		if err == nil {
			t.Error("Should Error Unmarshal Order JSON data")
		}

		kred.Env = ts.URL
		checkoutResult := kred.Checkout(&order)

		if checkoutResult.Error == nil {
			t.Errorf("Checkout() Should returned an error: %s", checkoutResult.Error)
		}

	})

	t.Run("Test Kredivo Get Payments", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			if r.Method != "POST" {
				t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
			}

			if r.URL.EscapedPath() != "/payments" {
				t.Errorf("Expected request to ‘/payments, got ‘%s’", r.URL.EscapedPath())
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{
          "status": "OK",
          "message": "Available payment types are listed.",
          "payments": [
              {
                  "down_payment": 0,
                  "name": "Bayar dalam 30 hari",
                  "amount": 2400000,
                  "installment_amount": 2400000,
                  "rate": 0,
                  "monthly_installment": 2400000,
                  "discounted_monthly_installment": 0,
                  "tenure": 1,
                  "id": "30_days"
              }
              ]}`))
		}))

		// close server
		defer ts.Close()

		//construct kredivo
		kred := New("123456", "bhinneka.com/kredivo/notif", "bhinneka.com", 8*time.Second)

		var paymentRequest PaymentRequest

		err := json.Unmarshal(paymentRequestData, &paymentRequest)

		if err != nil {
			t.Error("Cannot Unmarshal Payment Request JSON data")
		}

		//hit the get payment endpoint
		kred.Env = ts.URL
		paymentResult := kred.GetPayments(&paymentRequest)

		if paymentResult.Error != nil {
			t.Errorf("GetPayments() returned an error: %s", paymentResult.Error)
		}

	})

	t.Run("Test Kredivo Get Payments Error Empty Result", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			if r.Method != "POST" {
				t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
			}

			if r.URL.EscapedPath() != "/payments" {
				t.Errorf("Expected request to ‘/payments, got ‘%s’", r.URL.EscapedPath())
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`"error"`))
		}))

		// close server
		defer ts.Close()

		//construct kredivo
		kred := New("123456", "bhinneka.com/kredivo/notif", "bhinneka.com", 8*time.Second)

		var paymentRequest PaymentRequest

		err := json.Unmarshal([]byte(""), &paymentRequest)

		if err == nil {
			t.Error("Should be Error Unmarshal Payment Request JSON data")
		}

		//hit the get payment endpoint
		kred.Env = ts.URL
		paymentResult := kred.GetPayments(&paymentRequest)

		if paymentResult.Error == nil {
			t.Errorf("GetPayments() Should returned an error: %s", paymentResult.Error)
		}

	})

	t.Run("Test Kredivo Confirm Payment", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			if r.Method != "GET" {
				t.Errorf("Expected GET request, got ‘%s’", r.Method)
			}

			if r.URL.EscapedPath() != "/update" {
				t.Errorf("Expected request to ‘/update, got ‘%s’", r.URL.EscapedPath())
			}

			if r.URL.Query().Get("transaction_id") == "" && r.URL.Query().Get("signature_key") == "" {
				t.Error("Expected request to ‘/update, got transaction id and signature key")
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{
            "status": "OK",
            "legal_name": "TANDRA",
            "fraud_status": "accept",
            "order_id": "KD125262",
            "transaction_time": 1501846094,
            "amount": "1500100.00",
            "payment_type": "30_days",
            "transaction_status": "settlement",
            "message": "Confirmed order status. Valid!",
            "transaction_id": "fadee4e5-99a2-48d6-952d-007f3fa508e8"
        }`))
		}))

		// close server
		defer ts.Close()

		//construct kredivo
		kred := New("123456", "bhinneka.com/kredivo/notif", "bhinneka.com", 8*time.Second)

		var confirmRequest ConfirmRequest
		confirmRequest.TransactionID = "f6e33997-9ea2-426c-991b-9421165b9343"
		confirmRequest.SignatureKey = "YmQBYtAODqlWkmVrkNY%2BRtHclC9yHMDsKwAJ%2BG4n%2BQ1m1DlgqtIfsjjQrUFEIs%2BnWlHhJahCmJuj%2FSOJ8YmPPuX9iKoquvfJl5n0V57Cess%3D"

		//hit the confirm endpoint
		kred.Env = ts.URL
		confirmResult := kred.Confirm(&confirmRequest)

		if confirmResult.Error != nil {
			t.Errorf("Confirm() returned an error: %s", confirmResult.Error)
		}

	})

	t.Run("Test Kredivo Cancel Transaction", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			if r.Method != "POST" {
				t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
			}

			if r.URL.EscapedPath() != "/cancel_transaction" {
				t.Errorf("Expected request to ‘/cancel_transaction, got ‘%s’", r.URL.EscapedPath())
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{
            "status": "OK",
            "fraud_status": "accept",
            "order_id": "KD14721",
            "transaction_time": 1501842660,
            "amount": "1515521.00",
            "payment_type": "30_days",
            "transaction_status": "cancel",
            "message": "Cancelled the transaction!",
            "transaction_id": "6febc2b2-ac4f-462c-9e7e-56fc5da05d91"
        }`))
		}))

		// close server
		defer ts.Close()

		//construct kredivo
		kred := New("123456", "bhinneka.com/kredivo/notif", "bhinneka.com", 8*time.Second)

		var cancelRequest CancelRequest

		err := json.Unmarshal(cancelRequestData, &cancelRequest)

		if err != nil {
			t.Error("Cannot Unmarshal Cancel Request JSON data")
		}

		//hit the cancel transaction endpoint
		kred.Env = ts.URL
		cancelResult := kred.Cancel(&cancelRequest)

		if cancelResult.Error != nil {
			t.Errorf("Cancel() returned an error: %s", cancelResult.Error)
		}

	})

	t.Run("Test Kredivo Cancel Transaction Error Empty Result", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			if r.Method != "POST" {
				t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
			}

			if r.URL.EscapedPath() != "/cancel_transaction" {
				t.Errorf("Expected request to ‘/cancel_transaction, got ‘%s’", r.URL.EscapedPath())
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`error`))
		}))

		// close server
		defer ts.Close()

		//construct kredivo
		kred := New("123456", "bhinneka.com/kredivo/notif", "bhinneka.com", 8*time.Second)

		var cancelRequest CancelRequest

		err := json.Unmarshal([]byte(""), &cancelRequest)

		if err == nil {
			t.Error("Should Error Unmarshal Cancel Request JSON data")
		}

		//hit the cancel transaction endpoint
		kred.Env = ts.URL
		cancelResult := kred.Cancel(&cancelRequest)

		if cancelResult.Error == nil {
			t.Errorf("Cancel() Should returned an error: %s", cancelResult.Error)
		}

	})

	t.Run("Test Kredivo Partial Cancel Transaction", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			if r.Method != "POST" {
				t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
			}

			if r.URL.EscapedPath() != "/partial_cancel_transaction" {
				t.Errorf("Expected request to ‘/partial_cancel_transaction, got ‘%s’", r.URL.EscapedPath())
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{
          "status": "OK",
          "message": "Processed Partial Cancellation successfully!",
          "fraud_status": "accept"
      }`))
		}))

		// close server
		defer ts.Close()

		//construct kredivo
		kred := New("123456", "bhinneka.com/kredivo/notif", "bhinneka.com", 8*time.Second)

		var partialRequest PartialCancelRequest

		err := json.Unmarshal(partialCancelData, &partialRequest)

		if err != nil {
			t.Error("Cannot Unmarshal Cancel Request JSON data")
		}

		//hit the partial cancel transaction endpoint
		kred.Env = ts.URL
		cancelResult := kred.PartialCancel(&partialRequest)

		if cancelResult.Error != nil {
			t.Errorf("PartialCancel() returned an error: %s", cancelResult.Error)
		}

	})

	t.Run("Test Kredivo Partial Cancel Transaction Error Empty Result", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			if r.Method != "POST" {
				t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
			}

			if r.URL.EscapedPath() != "/partial_cancel_transaction" {
				t.Errorf("Expected request to ‘/partial_cancel_transaction, got ‘%s’", r.URL.EscapedPath())
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`error`))
		}))

		// close server
		defer ts.Close()

		//construct kredivo
		kred := New("123456", "bhinneka.com/kredivo/notif", "bhinneka.com", 8*time.Second)

		var partialRequest PartialCancelRequest

		err := json.Unmarshal([]byte(`error`), &partialRequest)

		if err == nil {
			t.Error("Should Return Error Unmarshal Cancel Request JSON data")
		}

		//hit the partial cancel transaction endpoint
		kred.Env = ts.URL
		cancelResult := kred.PartialCancel(&partialRequest)

		if cancelResult.Error == nil {
			t.Errorf("PartialCancel() Should returned an error: %s", cancelResult.Error)
		}

	})
}
