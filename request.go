package kredivo

//ConfirmRequest
type ConfirmRequest struct {
	TransactionID string
	SignatureKey  string
}

//PaymentRequest request, for getting available list payment
type PaymentRequest struct {
	ServerKey string `json:"server_key"`
	Amount    int    `json:"amount"`
	Items     []Item `json:"items"`
}

//CancelRequest request, for cancel transaction
type CancelRequest struct {
	ServerKey          string `json:"server_key"`
	OrderID            string `json:"order_id"`
	TransactionID      string `json:"transaction_id"`
	CancellationReason string `json:"cancellation_reason"`
	CancelledBy        string `json:"cancelled_by"`
	CancellationDate   string `json:"cancellation_date"`
}

//PartialCancelRequest request, for cancel transaction
type PartialCancelRequest struct {
	ServerKey          string             `json:"server_key"`
	TransactionDetails TransactionDetails `json:"transaction_details"`
	CancelledBy        string             `json:"cancelled_by"`
	CancellationDate   string             `json:"cancellation_date"`
}

//Order request
type Order struct {
	ServerKey          string             `json:"server_key"`
	PaymentType        string             `json:"payment_type"`
	TransactionDetails TransactionDetails `json:"transaction_details"`
	Sellers            []Seller           `json:"sellers,omitempty"`
	CustomerDetails    *CustomerDetails   `json:"customer_details,omitempty"`
	BillingAddress     *Address           `json:"billing_address,omitempty"`
	ShippingAddress    Address            `json:"shipping_address"`
	UserCancelURI      string             `json:"user_cancel_uri,omitempty"`
	PushURI            string             `json:"push_uri"`
	BackToStoreURI     string             `json:"back_to_store_uri"`
}

//TransactionDetails request
type TransactionDetails struct {
	Amount  float64 `json:"amount"`
	OrderID string  `json:"order_id"`
	Items   []Item  `json:"items"`
}

//Item request
type Item struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	Price              float64 `json:"price"`
	Type               string  `json:"type,omitempty"`
	URL                string  `json:"url,omitempty"`
	Quantity           int     `json:"quantity"`
	ParentType         string  `json:"parent_type,omitempty"`
	ParentID           string  `json:"parent_id,omitempty"`
	CancellationReason string  `json:"cancellation_reason,omitempty"`
}

//CustomerDetails request
type CustomerDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

//Address request
type Address struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country_code"`
}

//Seller request
type Seller struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

//TransactionStatusRequest request, for getting transaction status
type TransactionStatusRequest struct {
	ServerKey string `json:"server_key"`
	OrderID   string `json:"order_id"`
}
