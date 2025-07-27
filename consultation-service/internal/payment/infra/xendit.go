package infra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type xendit struct{}

func NewXendit() *xendit {
	return &xendit{}
}

func (x *xendit) Create(external_id string, amount float64) (interface{}, error) {
	// invoiceReq := dto.XenditInvoiceRequest{
	// 	ExternalID: "booking-" + strconv.FormatUint(uint64(newBooking.ID), 10),
	// 	Amount:     amountToPay,
	// 	Customer: dto.XenditCustomer{
	// 		GivenNames: user.FullName,
	// 		Email:      user.Email,
	// 	},
	// }

	paymentReq := map[string]interface{}{
		"external_id": external_id,
		"amount":      amount,
	}

	reqBody, err := json.Marshal(paymentReq)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", "https://api.xendit.co/v2/invoices", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	xenditKey := os.Getenv("XENDIT_KEY")
	if xenditKey == "" {
		return nil, fmt.Errorf("XENDIT_KEY environment variable is not set")
	}

	fmt.Println(xenditKey)

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(xenditKey, "")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to create invoice: %s", string(body))
	}

	var xenditResponse interface{}
	if err := json.NewDecoder(resp.Body).Decode(&xenditResponse); err != nil {
		return nil, err
	}

	return xenditResponse, nil
}

// }
// req.Header.Set("Content-Type", "application/json")
// req.Header.Set("Authorization", authHeader)

// resp, err := client.Do(req)
// if err != nil {
// 	return dto.BookingResponse{}, echo.NewHTTPError(500, "failed to send invoice request")
// }
// defer resp.Body.Close()

// if resp.StatusCode != 200 && resp.StatusCode != 201 {
// 	body, _ := io.ReadAll(resp.Body)
// 	return dto.BookingResponse{}, echo.NewHTTPError(resp.StatusCode, "failed to create invoice: "+string(body))
// }

// var invoiceResp dto.XenditInvoiceResponse
// if err := json.NewDecoder(resp.Body).Decode(&invoiceResp); err != nil {
// 	return dto.BookingResponse{}, echo.NewHTTPError(500, "failed to decode invoice response")
// }

// response.PaymentURL = invoiceResp.InvoiceURL
