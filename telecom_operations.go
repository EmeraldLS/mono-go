package monogo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EmeraldLS/monogo/types"
	"github.com/EmeraldLS/monogo/utils"
)

// This resource is to get telco account information such as the name of the telco provider, connected phone number etc via the telco account id.
func (m *MonoClient) GetTelecomDetails(telcoAccID string) (*types.TelcoAccountDetailsResponse, error) {
	path := fmt.Sprintf("/%s/%s", telcoAccBasePath, telcoAccID)

	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var detailsRes types.TelcoAccountDetailsResponse
	err = json.NewDecoder(res.Body).Decode(&detailsRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	return &detailsRes, nil
}

// This resource is to pull the customer's transaction balance data, which includes main balance, airtime bundle, data bonuses etc.
func (m *MonoClient) GetTelcoBalances(telcoAccID string) ([]types.TelcoAccountBalanceResponse, error) {
	path := fmt.Sprintf("/%s/%s/balances", telcoAccBasePath, telcoAccID)

	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var balances []types.TelcoAccountBalanceResponse
	err = json.NewDecoder(res.Body).Decode(&balances)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	return balances, nil
}

// This resource is to get all telco-related account transactions such as the transaction id, type etc.
func (m *MonoClient) GetTelcoTransactions(telcoAccID string, params *types.AccountTransactionParams) (*types.TelcoTransactionsResponse, *types.TelcoTransactionsErrorResponse, error) {
	path := fmt.Sprintf("%s/%s/transactions?start=%s&end=%s&narration=%s&type=%s&paginate=%s&limits=%d", telcoAccBasePath, telcoAccID, params.Start, params.End, params.Narration, params.Type, strconv.FormatBool(params.Paginate), params.Limits)

	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, nil, utils.RequestError(err)
	}

	if res.StatusCode != http.StatusOK {
		var errRes types.TelcoTransactionsErrorResponse
		err = json.NewDecoder(res.Body).Decode(&errRes)
		if err != nil {
			return nil, nil, utils.DecodeError(err)
		}
		return nil, &errRes, errors.New("server error occured")
	}

	var telcoTransactions types.TelcoTransactionsResponse
	err = json.NewDecoder(res.Body).Decode(&telcoTransactions)
	if err != nil {
		return nil, nil, utils.DecodeError(err)
	}

	return &telcoTransactions, nil, nil
}

// This resource is to get the identity information from a connected telco account.
func (m *MonoClient) GetTelcoIdentity(telcoAccID string) (*types.TelcoIdentityResponse, *types.TelcoIdentityErrorResponse, error) {
	path := fmt.Sprintf("/%s/%s/identity", telcoAccBasePath, telcoAccID)

	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, nil, utils.RequestError(err)
	}

	if res.StatusCode != http.StatusOK {
		var errRes types.TelcoIdentityErrorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, nil, utils.DecodeError(err)
		}

		return nil, &errRes, errors.New("no document found")
	}

	var identityRes types.TelcoIdentityResponse
	err = json.NewDecoder(res.Body).Decode(&identityRes)
	if err != nil {
		return nil, nil, utils.DecodeError(err)
	}

	return &identityRes, nil, nil
}
