package picpay_test

import (
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/rafaeltokyo/picpay-sdk-go"
)

func TestPaymentCreate(t *testing.T) {
	godotenv.Load()
	client := picpay.New(os.Getenv("PICPAY_TOKEN"), os.Getenv("ENV"))
	response, errAPI, err := client.Payment().Create(&picpay.PaymentRequest{
		ReferenceID: "1010011",
		CallbackURL: "https://google.com",
		ReturnURL:   "https://google.com",
		Value:       10.10,
		ExpiresAt:   time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
		Buyer: &picpay.PaymentBuyer{
			FirstName: "João",
			LastName:  "Da Silva",
			Document:  "123.456.789-10",
			Email:     "teste@picpay.com",
			Phone:     "+55 27 12345-6789",
		},
	})
	if err != nil {
		t.Errorf("err : %#v", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if response == nil {
		t.Error("response is null")
		return
	}
	t.Errorf("ref %v", response)
}

func TestPaymentGetStatusCreated(t *testing.T) {
	godotenv.Load()
	client := picpay.New(os.Getenv("PICPAY_TOKEN"), os.Getenv("ENV"))
	ReferenceID := os.Getenv("REFERENCEID_CREATED")
	responseStatus, errAPI, err := client.Payment().Status(ReferenceID)
	if err != nil {
		t.Errorf("err : %#v", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if responseStatus == nil {
		t.Error("responseStatus is null")
		return
	}
	if responseStatus.Status != "created" {
		t.Error("status is invalid")

	}

}

func TestPaymentGetStatusPaid(t *testing.T) {
	godotenv.Load()
	client := picpay.New(os.Getenv("PICPAY_TOKEN"), os.Getenv("ENV"))
	ReferenceID := os.Getenv("REFERENCEID_PAID")
	responseStatus, errAPI, err := client.Payment().Status(ReferenceID)
	if err != nil {
		t.Errorf("err : %#v", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if responseStatus == nil {
		t.Error("responseStatus is null")
		return
	}
	if responseStatus.Status != "paid" {
		t.Error("status is invalid")

	}

}
