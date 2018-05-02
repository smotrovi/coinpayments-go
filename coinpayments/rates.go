package coinpayments

import (
	"net/http"

	"github.com/dghubble/sling"
)

type RateService struct {
	sling        *sling.Sling
	ApiPublicKey string
	Params       RateBodyParams
}

type RateBodyParams struct {
	APIParams
	rateParamsWithAccepted
}

type rateParamsWithAccepted struct {
	Short    uint8 `url:"short"`
	Accepted uint8 `url:"accepted"`
}

func newRateService(sling *sling.Sling, apiPublicKey string) *RateService {
	rateService := &RateService{
		sling:        sling.Path("api.php"),
		ApiPublicKey: apiPublicKey,
	}
	// In each request the params are the same.
	rateService.getParams()
	return rateService
}

func (s *RateService) getHMAC() string {
	return getHMAC(getPayload(s.Params))
}

type RateInfo struct {
	IsFiat         uint8    `json:"is_fiat"`
	RateBTC        string   `json:"rate_btc"`
	LastUpdate     string   `json:"last_update"`
	TransactionFee string   `json:"tx_fee"`
	Name           string   `json:"name"`
	Confirms       string   `json:"confirms"`
	CanConvert     uint8    `json:"can_convert"`
	Status         string   `json:"status"`
	Capabilities   []string `json:"capabilities"`
}

type RateResponse struct {
	Error  string              `json:"error"`
	Result map[string]RateInfo `json:"result"`
}

type RateParams struct {
	Short uint8 `url:"short"`
}

func (s *RateService) Show(params *RateParams) (RateResponse, *http.Response, error) {
	rateResponse := new(RateResponse)
	s.Params.Short = params.Short
	resp, err := s.sling.New().Set("HMAC", s.getHMAC()).Post(
		"api.php").BodyForm(s.Params).ReceiveSuccess(rateResponse)
	return *rateResponse, resp, err
}

type RateResponseWithAccepted struct {
	Error  string                          `json:"error"`
	Result map[string]RateInfoWithAccepted `json:"result"`
}

type RateInfoWithAccepted struct {
	IsFiat         uint8    `json:"is_fiat"`
	RateBTC        string   `json:"rate_btc"`
	LastUpdate     string   `json:"last_update"`
	TransactionFee string   `json:"tx_fee"`
	Name           string   `json:"name"`
	Confirms       string   `json:"confirms"`
	CanConvert     uint8    `json:"can_convert"`
	Status         string   `json:"status"`
	Accepted       int      `json:"accepted"`
	Capabilities   []string `json:"capabilities"`
}

func (s *RateService) ShowWithAccepted(params *RateParams) (RateResponseWithAccepted, *http.Response, error) {
	rateResponse := new(RateResponseWithAccepted)
	s.Params.Short = params.Short
	s.Params.Accepted = 1
	resp, err := s.sling.New().Set("HMAC", s.getHMAC()).Post(
		"api.php").BodyForm(s.Params).ReceiveSuccess(rateResponse)
	return *rateResponse, resp, err
}

func (s *RateService) getParams() {
	s.Params.Command = "rates"
	s.Params.Key = s.ApiPublicKey
	s.Params.Version = "1"
}
