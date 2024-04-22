package monogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EmeraldLS/monogo/types"
	"github.com/EmeraldLS/monogo/utils"
)

// This resource is to initiate linking an account.
func (m *MonoClient) InitiateAccountLinking(data *types.AccountLinkRequest) (*types.AccountLinkResponse, error) {

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(jsonBytes)

	res, err := m.withPath("/v2/accounts/initiate").withBody(buf).withMethod("POST").client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var accountLinkResponse types.AccountLinkResponse

	if err := json.NewDecoder(res.Body).Decode(&accountLinkResponse); err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected error occured: %v", accountLinkResponse.Message)
	}

	return &accountLinkResponse, nil
}

// Use this endpoint to reauthorise a previously linked account
func (m *MonoClient) Reauthorization(data *types.ReauthLinkRequest) (*types.AccountLinkResponse, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(jsonBytes)

	res, err := m.withBody(buf).withPath("/v2/accounts/initiate").withMethod("POST").client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var accountLinkResponse types.AccountLinkResponse

	err = json.NewDecoder(res.Body).Decode(&accountLinkResponse)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected error occured: %v", err)
	}

	return &accountLinkResponse, nil

}

// Use this endpoint to request an account ID (that identifies the authenticated account) after successful enrolment on the Mono Connect widget.
func (m *MonoClient) ExchangeToken(data *types.AccountExchangeTokenRequest) (*types.AccountExchangeTokenResponse, error) {
	path := "v2/accounts/auth"
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(jsonBytes)

	res, err := m.withBody(buf).withPath(path).withMethod("POST").client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var exResponse types.AccountExchangeTokenResponse
	if err = json.NewDecoder(res.Body).Decode(&exResponse); err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error has occured: %v", exResponse.Message)
	}

	return &exResponse, nil

}
