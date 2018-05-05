package coinpayments

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithdrawalService_CreateWithdrawal(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api.php", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"error":"ok","result":{"id":"CPBB0FWI7BW9NUN2REJ3MS56KL","status":0,"amount":0.1}}`)
	})

	client := NewClient("", "", httpClient)
	r, _, err := client.Withdrawal.CreateWithdrawal(&CreateWithdrawalParams{
		Amount:      0.1,
		Currency:    "BTC",
		Currency2:   "BTC",
		Address:     "3BHLa87gpzKSemapUtY6gHhJMEE6UWqPuf",
		Pbntag:      "",
		DestTag:     "",
		IPNUrl:      "",
		AutoConfirm: false,
		Note:        "",
	})

	expected := CreateWithdrawalResponse{
		Error: "ok",
		Result: &createWithdrawalResponseWithdrawal{
			ID:     "CPBB0FWI7BW9NUN2REJ3MS56KL",
			Status: 0,
			Amount: 0.1,
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, r)
}
