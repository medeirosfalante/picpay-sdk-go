package picpay

import "fmt"

//PaymentBuyer - refer for struct Master for Payment

type Payment struct {
	client *APIClient
}

//PaymentRequest - Object for Create a new Payment
type PaymentRequest struct {
	ReferenceID string        `json:"referenceId"`
	CallbackURL string        `json:"callbackUrl"`
	ReturnURL   string        `json:"returnUrl"`
	Value       float32       `json:"value"`
	ExpiresAt   string        `json:"expiresAt"`
	Buyer       *PaymentBuyer `json:"buyer"`
}

//PaymentBuyer - refer for buyer in payment
type PaymentBuyer struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Document  string `json:"document"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

//PaymentResponse - refer for response in payment
type PaymentResponse struct {
	ReferenceID string         `json:"referenceId"`
	PaymentURL  string         `json:"paymentUrl"`
	ExpiresAt   string         `json:"expiresAt"`
	QrCode      *PaymentQrCode `json:"qrcode"`
}

//PaymentStatusResponse - refer for status payment
type PaymentStatusResponse struct {
	AuthorizationID string `json:"authorizationId"`
	ReferenceID     string `json:"referenceId"`
	Status          string `json:"status"`
}

//PaymentQrCode - Qrcode for payment
type PaymentQrCode struct {
	Content string `json:"content"`
	Base64  string `json:"base64"`
}

//Payment - func for return refer payment in client
func (c *APIClient) Payment() *Payment {
	return &Payment{client: c}
}

//Create - Create a new payment
func (p *Payment) Create(req *PaymentRequest) (*PaymentResponse, *Error, error) {
	response := &PaymentResponse{}

	err, errAPI := p.client.Request("POST", "/ecommerce/public/payments", req, nil, response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

//Status - Get Status Payment
func (p *Payment) Status(ReferenceID string) (*PaymentStatusResponse, *Error, error) {
	response := &PaymentStatusResponse{}
	err, errAPI := p.client.Request("GET", fmt.Sprintf("/ecommerce/public/payments/%s/status", ReferenceID), nil, nil, response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
