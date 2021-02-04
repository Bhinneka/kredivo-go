package kredivo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"
)

//kredivo entry struct for kredivo
type kredivo struct {
	serverKey      string
	userCancelURI  string
	pushURI        string
	backToStoreURI string
	client         *kredivoHttpClient
	Env            string
	*logger
}

/*New function, create kredivo pointer
	Required parameter :
	1. your Kredivo Merchant Key
  	2. pushURI parameter, your Notification URL (POST METHOD) that used by Kredivo to Send Payment Notification
  	3. backToStoreURI parameter,
	 	your redirect page when user finished their transaction,
	 	(Your Thank You Page's Website)
  	4. HTTP Request Timeout
*/
func New(serverKey, pushURI, backToStoreURI string, timeout time.Duration) *kredivo {
	httpRequest := newRequest(timeout)
	return &kredivo{
		serverKey: serverKey,

		pushURI: pushURI,

		backToStoreURI: backToStoreURI,

		//httpClient
		client: httpRequest,

		//set default env to SandBox,
		//latter just simply change with kredivo.Env = Production
		Env:    SandBox.String(),
		logger: newLogger(),
	}
}

//call function
func (r *kredivo) call(method, path string, body io.Reader, v interface{}, headers map[string]string) error {
	r.info().Println("Starting http call..")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = fmt.Sprintf("%s%s", r.Env, path)
	return r.client.exec(method, path, body, v, headers)
}

//Checkout Method
func (r *kredivo) Checkout(order *Order) ServiceResult {
	r.info().Println("Starting checkout..")

	//create map headers
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	//set ServerKey, PushURI, and BackToStoreURI to order request
	order.ServerKey = r.serverKey
	order.PushURI = r.pushURI
	order.BackToStoreURI = r.backToStoreURI

	//Marshal Order
	payload, err := json.Marshal(order)

	if err != nil {
		r.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	var checkoutResponse CheckoutResponse

	//call checkout endpoint
	err = r.call("POST", "checkout_url", bytes.NewBuffer(payload), &checkoutResponse, headers)

	if err != nil {
		r.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	return ServiceResult{Result: checkoutResponse}
}

//GetPayments Method, for get available list payment for specific order
func (r *kredivo) GetPayments(paymentRequest *PaymentRequest) ServiceResult {
	r.info().Println("Starting get available payments..")

	//create map headers
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	//set ServerKey, to payment request
	paymentRequest.ServerKey = r.serverKey

	//Marshal Order
	payload, err := json.Marshal(paymentRequest)

	if err != nil {
		r.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	var paymentResponse PaymentResponse

	//call checkout endpoint
	err = r.call("POST", "payments", bytes.NewBuffer(payload), &paymentResponse, headers)

	if err != nil {
		r.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	return ServiceResult{Result: paymentResponse}
}

//Confirm Method, for confirm order status
func (r *kredivo) Confirm(confirmRequest *ConfirmRequest) ServiceResult {
	r.info().Println("Starting Confirm payment..")

	//create map headers
	headers := make(map[string]string)

	var confirmationResponse ConfirmationResponse

	//set query params, transaction_id and signature_key
	queryParams := url.Values{}
	queryParams.Set("transaction_id", confirmRequest.TransactionID)
	queryParams.Set("signature_key", confirmRequest.SignatureKey)

	path := fmt.Sprintf("update?%s", queryParams.Encode())

	//call confirm endpoint
	err := r.call("GET", path, nil, &confirmationResponse, headers)

	if err != nil {
		r.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	return ServiceResult{Result: confirmationResponse}
}

//Cancel Method, for cancel transaction
func (r *kredivo) Cancel(cancelRequest *CancelRequest) ServiceResult {
	r.info().Println("Starting cancel..")

	//create map headers
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	//set ServerKey, to cancel request
	cancelRequest.ServerKey = r.serverKey

	//Marshal Cancel Request
	payload, err := json.Marshal(cancelRequest)

	if err != nil {
		r.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	var cancelResponse CancelResponse

	//call cancel transaction endpoint
	err = r.call("POST", "cancel_transaction", bytes.NewBuffer(payload), &cancelResponse, headers)

	if err != nil {
		r.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	return ServiceResult{Result: cancelResponse}
}

//PartialCancel Method, for cancel partial transaction
func (r *kredivo) PartialCancel(partialRequest *PartialCancelRequest) ServiceResult {
	r.info().Println("Starting partial cancel..")

	//create map headers
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	//set ServerKey, to partial cancel request
	partialRequest.ServerKey = r.serverKey

	//Marshal Partial Cancel Request
	payload, err := json.Marshal(partialRequest)

	if err != nil {
		r.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	var cancelResponse CancelResponse

	//call partial cancel transaction endpoint
	err = r.call("POST", "partial_cancel_transaction", bytes.NewBuffer(payload), &cancelResponse, headers)

	if err != nil {
		r.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	return ServiceResult{Result: cancelResponse}
}

//PartialCancel Method, for cancel partial transaction
func (r *kredivo) TransactionStatus(tansactionStatusRequest *TransactionStatusRequest) ServiceResult {
	r.info().Println("Starting get transaction status..")

	//create map headers
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	//set ServerKey, to partial cancel request
	tansactionStatusRequest.ServerKey = r.serverKey

	//Marshal Partial Cancel Request
	payload, err := json.Marshal(tansactionStatusRequest)

	if err != nil {
		r.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	var transactionStatusReponse TransactionStatusResponse

	//call partial cancel transaction endpoint
	err = r.call("POST", "transaction/status", bytes.NewBuffer(payload), &transactionStatusReponse, headers)

	if err != nil {
		r.error().Println(err.Error())
		return ServiceResult{Error: err}
	}

	return ServiceResult{Result: transactionStatusReponse}
}

func GenerateServiceResult(data interface{}, err error) ServiceResult {
	var output ServiceResult
	output = ServiceResult{Result: data, Error: err}
	return output
}
