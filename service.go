package kredivo

//ServiceResult
type ServiceResult struct {
	Result interface{}
	Error  error
}

//KredivoService
//Generic abstraction for Kredivo Service
type KredivoService interface {
	//Checkout Method
	Checkout(*Order) ServiceResult

	//GetPayments Method, for get available list payment for specific order
	GetPayments(*PaymentRequest) ServiceResult

	//Confirm Method, for confirm order status
	Confirm(*ConfirmRequest) ServiceResult

	//Cancel Method, for cancel transaction
	Cancel(*CancelRequest) ServiceResult

	//PartialCancel Method, for cancel partial transaction
	PartialCancel(*PartialCancelRequest) ServiceResult

	//TransactionStatus Method, for get transaction status
	TransactionStatus(*TransactionStatusRequest) ServiceResult
}
