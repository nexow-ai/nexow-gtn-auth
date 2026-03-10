package gtn

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Client calls GTN auth and back-office APIs.
type Client struct {
	baseURL        string
	appKey         string
	appSecret      string
	instCode       string
	userID         string
	privateKeyDER  []byte
	throttleKey    string
	httpClient     *http.Client
	assertionExpiry int64
}

// NewClient builds a GTN API client.
func NewClient(baseURL, appKey, appSecret, instCode, userID string, privateKeyDER []byte, throttleKey string) *Client {
	return &Client{
		baseURL:         strings.TrimSuffix(baseURL, "/"),
		appKey:           appKey,
		appSecret:        appSecret,
		instCode:         instCode,
		userID:           userID,
		privateKeyDER:    privateKeyDER,
		throttleKey:      throttleKey,
		httpClient:       &http.Client{},
		assertionExpiry:  3600,
	}
}

// TokenResponse is the response from GET token and refresh endpoints.
type TokenResponse struct {
	Status               string `json:"status"`
	Reason               string `json:"reason"`
	RejectCode           int    `json:"rejectCode"`
	AccessToken          string `json:"accessToken"`
	RefreshToken         string `json:"refreshToken"`
	AccessTokenExpiresAt int64  `json:"accessTokenExpiresAt"`
	RefreshTokenExpiresAt int64 `json:"refreshTokenExpiresAt"`
	TokenType            string `json:"tokenType"`
}

// GetServerToken obtains a new server access token from GTN.
func (c *Client) GetServerToken() (*TokenResponse, error) {
	assertion, err := BuildAssertion(c.privateKeyDER, c.appKey, c.instCode, c.userID, c.assertionExpiry)
	if err != nil {
		return nil, fmt.Errorf("build assertion: %w", err)
	}
	body := map[string]string{"assertion": assertion}
	bodyBytes, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPost, c.baseURL+"/trade/auth/token", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Throttle-Key", c.throttleKey)
	req.SetBasicAuth(c.appKey, c.appSecret)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return decodeTokenResponse(resp)
}

// RefreshServerToken refreshes the server token using refreshToken.
func (c *Client) RefreshServerToken(refreshToken string) (*TokenResponse, error) {
	body := map[string]string{"refreshToken": refreshToken}
	bodyBytes, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPost, c.baseURL+"/trade/auth/token/refresh", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Throttle-Key", c.throttleKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return decodeTokenResponse(resp)
}

// GetCustomerToken returns customer access token using server token and customer number.
func (c *Client) GetCustomerToken(serverAccessToken, customerNumber string) (*TokenResponse, error) {
	body := map[string]string{
		"customerNumber": customerNumber,
		"accessToken":    serverAccessToken,
	}
	bodyBytes, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPost, c.baseURL+"/trade/auth/customer/token", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Throttle-Key", c.throttleKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return decodeTokenResponse(resp)
}

// RefreshCustomerToken refreshes customer token.
func (c *Client) RefreshCustomerToken(refreshToken string) (*TokenResponse, error) {
	body := map[string]string{"refreshToken": refreshToken}
	bodyBytes, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPost, c.baseURL+"/trade/auth/customer/token/refresh", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Throttle-Key", c.throttleKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return decodeTokenResponse(resp)
}

func decodeTokenResponse(resp *http.Response) (*TokenResponse, error) {
	body, _ := io.ReadAll(resp.Body)
	var out TokenResponse
	_ = json.Unmarshal(body, &out)
	if resp.StatusCode != http.StatusOK {
		errMsg := out.Reason
		if errMsg == "" {
			errMsg = resp.Status
			if len(body) > 0 && len(body) < 500 {
				errMsg = string(body)
			}
		}
		msg := fmt.Sprintf("GTN %d: %s (rejectCode=%d)", resp.StatusCode, errMsg, out.RejectCode)
		if resp.StatusCode == 403 {
			msg += " — check IP allowlist, institution API access, and base URL"
		}
		return nil, fmt.Errorf("%s", msg)
	}
	if out.Status == "FAILED" {
		return nil, fmt.Errorf("GTN: %s (rejectCode=%d)", out.Reason, out.RejectCode)
	}
	return &out, nil
}

// CreateCustomerRequest mirrors GTN Create Customer body.
type CreateCustomerRequest struct {
	ReferenceNumber      string `json:"referenceNumber"`
	InstitutionCode      string `json:"institutionCode"`
	FirstName            string `json:"firstName,omitempty"`
	LastName             string `json:"lastName,omitempty"`
	PassportNumber       string `json:"passportNumber,omitempty"`
	NIN                  string `json:"nin,omitempty"`
	DrivingLicense       string `json:"drivingLicense,omitempty"`
	HomeTel              string `json:"homeTel,omitempty"`
	OfficeTel            string `json:"officeTel,omitempty"`
	Mobile               string `json:"mobile,omitempty"`
	Email                string `json:"email,omitempty"`
	Profession           string `json:"profession,omitempty"`
	Address1             string `json:"address1,omitempty"`
	Address2             string `json:"address2,omitempty"`
	City                 string `json:"city,omitempty"`
	CountryCode          string `json:"countryCode,omitempty"`
	Gender               string `json:"gender,omitempty"`
	BirthDate            string `json:"birthDate,omitempty"`
	Nationality          string `json:"nationality,omitempty"`
	ProfileID            string `json:"profileId,omitempty"`
	MasterAccountNumber  string `json:"masterAccountNumber,omitempty"`
	PreferredLanguage    string `json:"preferredLanguage,omitempty"`
}

// CreateCustomerResponse is the GTN create customer response.
type CreateCustomerResponse struct {
	Status               string   `json:"status"`
	Reason               string   `json:"reason,omitempty"`
	RejectCode           int      `json:"rejectCode,omitempty"`
	CustomerNumber       string   `json:"customerNumber,omitempty"`
	CashAccountNumbers   []string `json:"cashAccountNumbers,omitempty"`
	AccountNumbers       []string `json:"accountNumbers,omitempty"`
	ExchangeAccountIds   []string `json:"exchangeAccountIds,omitempty"`
}

// CreateCustomer calls GTN POST /trade/bo/v1.2/customer/account with server token.
func (c *Client) CreateCustomer(serverAccessToken string, req *CreateCustomerRequest) (*CreateCustomerResponse, error) {
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, c.baseURL+"/trade/bo/v1.2/customer/account", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "*/*")
	httpReq.Header.Set("Throttle-Key", c.throttleKey)
	httpReq.Header.Set("Authorization", "Bearer "+serverAccessToken)

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	outBody, _ := io.ReadAll(resp.Body)
	var out CreateCustomerResponse
	if err := json.Unmarshal(outBody, &out); err != nil {
		return nil, fmt.Errorf("decode create customer response: %w", err)
	}
	if resp.StatusCode != http.StatusOK || out.Status == "FAILED" {
		return &out, fmt.Errorf("create customer failed: status=%s reason=%s rejectCode=%d", out.Status, out.Reason, out.RejectCode)
	}
	return &out, nil
}

// BasicAuthHeader returns "Basic base64(appKey:appSecret)" for debugging/docs.
func BasicAuthHeader(appKey, appSecret string) string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(appKey+":"+appSecret))
}
