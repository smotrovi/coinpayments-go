package coinpayments

import (
	"net/http"
	"fmt"
	"github.com/dghubble/sling"
)

type WithdrawalService struct {
	sling        *sling.Sling
	ApiPublicKey string
	Params       APIParams
}

type CreateWithdrawalParams struct {
	Amount      float64 `url:"amount"`
	Currency    string  `url:"currency"`
	Currency2   string  `url:"currency2"`
	Address     string  `url:"address"`
	Pbntag      string  `url:"pbntag"`
	DestTag     string  `url:"dest_tag"`
	IPNUrl      string  `url:"ipn_url"`
	AutoConfirm bool    `url:"auto_confirm"`
	Note        string  `url:"note"`
}

type CreateWithdrawalBodyParams struct {
	APIParams
	CreateWithdrawalParams
}

func newWithdrawalService(sling *sling.Sling, apiPublicKey string) *WithdrawalService {
	s := &WithdrawalService{
		sling:        sling.Path("api.php"),
		ApiPublicKey: apiPublicKey,
	}

	s.Params.Version = "1"
	s.Params.Key = apiPublicKey
	return s
}

func (s *WithdrawalService) getHMAC() string {
	return getHMAC(getPayload(s.Params))
}

type createWithdrawalResponseWithdrawal struct {
	ID     string  `json:"id"`
	Status uint    `json:"status"`
	Amount float64 `json:"amount"`
}

type CreateWithdrawalResponse struct {
	Error  string                              `json:"error"`
	Result *createWithdrawalResponseWithdrawal `json:"result"`
}

func (s *WithdrawalService) CreateWithdrawal(params *CreateWithdrawalParams) (CreateWithdrawalResponse, *http.Response, error) {
	response := new(CreateWithdrawalResponse)
	s.Params.Command = "create_withdrawal"
	bodyParams := CreateWithdrawalBodyParams{
		APIParams:              s.Params,
		CreateWithdrawalParams: *params,
	}

	fmt.Println(getPayload(s.Params))
	fmt.Println(getHMAC(getPayload(s.Params)))
	resp, err := s.sling.New().Set("HMAC", s.getHMAC()).Post(
		"api.php").BodyForm(bodyParams).ReceiveSuccess(response)

	return *response, resp, err
}
