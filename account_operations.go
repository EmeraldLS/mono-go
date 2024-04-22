package monogo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EmeraldLS/monogo/types"
	"github.com/EmeraldLS/monogo/utils"
)

// This resource represents the account details with the financial institution.
func (m *MonoClient) GetDetails(accountID string) (*types.AccountDataResponse, error) {
	path := fmt.Sprintf("%s/%s", accountBasePath, accountID)
	m.req.Header.Set("x-realtime", "true")

	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var accountResponse types.AccountDataResponse
	err = json.NewDecoder(res.Body).Decode(&accountResponse)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("errror occured: %v", accountResponse.Message)
	}

	return &accountResponse, nil
}

// This resource provides a mini customer identity information
func (m *MonoClient) GetIdentity(accountID string) (*types.IdentityResponse, error) {
	path := fmt.Sprintf("%s/%s/identity", accountBasePath, accountID)

	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var identityResponse types.IdentityResponse
	err = json.NewDecoder(res.Body).Decode(&identityResponse)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("errror occured: %v", identityResponse.Message)
	}

	return &identityResponse, nil
}

// This resource provides the account balance of the user
func (m *MonoClient) GetAccountBalance(accountID string) (*types.AccountBalanceResponse, *types.AccountBalanceErrorResponse) {
	path := fmt.Sprintf("%s/%s/balance", accountBasePath, accountID)

	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, &types.AccountBalanceErrorResponse{
			Error:            "failed",
			ErrorDescription: utils.RequestError(err),
		}
	}

	var accBalErrResponse types.AccountBalanceErrorResponse
	if res.StatusCode == http.StatusBadRequest {
		if err = json.NewDecoder(res.Body).Decode(&accBalErrResponse); err != nil {
			return nil, &types.AccountBalanceErrorResponse{
				Error:            "failed",
				ErrorDescription: utils.DecodeError(err),
			}
		}

	}

	var accBalResponse types.AccountBalanceResponse
	if res.StatusCode == http.StatusOK {
		if err = json.NewDecoder(res.Body).Decode(&accBalResponse); err != nil {
			return nil, &types.AccountBalanceErrorResponse{
				Error:            "failed",
				ErrorDescription: utils.DecodeError(err),
			}
		}
	}

	return &accBalResponse, &accBalErrResponse
}

// This enables you to provide your customers with the option to unlink their financial account(s)
func (m *MonoClient) UnlinkAccount(accountID string) (*types.UnlinkAccountResponse, error) {
	path := fmt.Sprintf("%s/%s/unlink", accountBasePath, accountID)

	res, err := m.withPath(path).withMethod("POST").client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var unlinkRes types.UnlinkAccountResponse
	if err := json.NewDecoder(res.Body).Decode(&unlinkRes); err != nil {
		return nil, utils.DecodeError(err)
	}

	return &unlinkRes, nil
}

// This resource provides all accounts linked to your business.
func (m *MonoClient) GetCustomerAccounts() (*types.CustomerAccountsResponse, *types.CustomerAccountsErrorResponse) {
	res, err := m.withPath(accountBasePath).client.Do(m.req)
	if err != nil {
		return nil, &types.CustomerAccountsErrorResponse{
			Error:            "failed",
			ErrorDescription: utils.RequestError(err),
		}
	}
	var accountsErrResponse types.CustomerAccountsErrorResponse
	if res.StatusCode == http.StatusBadRequest {
		if err := json.NewDecoder(res.Body).Decode(&accountsErrResponse); err != nil {
			return nil, &types.CustomerAccountsErrorResponse{
				Error:            "failed",
				ErrorDescription: utils.DecodeError(err),
			}
		}
	}

	var accountsResponse types.CustomerAccountsResponse
	if res.StatusCode == http.StatusOK {
		if err := json.NewDecoder(res.Body).Decode(&accountsErrResponse); err != nil {
			return nil, &types.CustomerAccountsErrorResponse{
				Error:            "failed",
				ErrorDescription: utils.DecodeError(err),
			}
		}
	}

	return &accountsResponse, &accountsErrResponse
}
