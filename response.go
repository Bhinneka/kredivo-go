package kredivo

//CheckoutResponse
type CheckoutResponse struct {
	Status                string                 `json:"status"`
	Message               string                 `json:"message"`
	RedirectURL           string                 `json:"redirect_url"`
	CheckoutErrorResponse *CheckoutErrorResponse `json:"error,omitempty"`
}

//CheckoutErrorResponse
type CheckoutErrorResponse struct {
	Message string `json:"message"`
	Kind    string `json:"kind"`
}

//Payment response
type PaymentResponse struct {
	Status   string    `json:"status"`
	Message  string    `json:"message"`
	Payments []Payment `json:"payments"`
}

//Payment response
type Payment struct {
	DownPayment                  float64 `json:"down_payment"`
	Name                         string  `json:"name"`
	Amount                       float64 `json:"amount"`
	InstallmentAmount            float64 `json:"installment_amount"`
	Rate                         float64 `json:"rate"`
	MonthlyInstallment           float64 `json:"monthly_installment"`
	DiscountedMonthlyInstallment float64 `json:"discounted_monthly_installment"`
	Tenure                       int     `json:"tenure"`
	ID                           string  `json:"id"`
	InterestRateTransitionTerm   float64 `json:"interest_rate_transition_term,omitempty"`
}

//Notification response, for handle payment Notification from Kredivo
type Notification struct {
	Status            string          `json:"status"`
	Amount            string          `json:"amount"`
	PaymentType       string          `json:"payment_type"`
	TransactionStatus string          `json:"transaction_status"`
	OrderID           string          `json:"order_id"`
	Message           string          `json:"message"`
	ShippingAddress   ShippingAddress `json:"shipping_address"`
	TransactionTime   int             `json:"transaction_time"`
	TransactionID     string          `json:"transaction_id"`
	SignatureKey      string          `json:"signature_key"`
}

//ShippingAddress of the customer/shopper
type ShippingAddress struct {
	City            string `json:"city"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Countrycode     string `json:"countrycode"`
	CreationDate    string `json:"creation_date"`
	Phone           string `json:"phone"`
	State           string `json:"state"`
	Transaction     int    `json:"transaction"`
	Postcode        string `json:"postcode"`
	LocationDetails string `json:"location_details"`
}

//StatusMessage, response status send to Kredivo
type StatusMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

//ConfirmationResponse, response status from update transaction
type ConfirmationResponse struct {
	Status            string `json:"status"`
	LegalName         string `json:"legal_name"`
	FraudStatus       string `json:"fraud_status"`
	OrderID           string `json:"order_id"`
	TransactionTime   int    `json:"transaction_time"`
	Amount            string `json:"amount"`
	PaymentType       string `json:"payment_type"`
	TransactionStatus string `json:"transaction_status"`
	Message           string `json:"message"`
	TransactionID     string `json:"transaction_id"`
}

//CancelResponse, response status from cancel transaction
type CancelResponse struct {
	Status            string `json:"status,omitempty"`
	FraudStatus       string `json:"fraud_status,omitempty"`
	OrderID           string `json:"order_id,omitempty"`
	TransactionTime   int    `json:"transaction_time,omitempty"`
	Amount            string `json:"amount,omitempty"`
	PaymentType       string `json:"payment_type,omitempty"`
	TransactionStatus string `json:"transaction_status,omitempty"`
	Message           string `json:"message,omitempty"`
	TransactionID     string `json:"transaction_id,omitempty"`
}
