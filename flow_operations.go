package monogo

import (
	"encoding/json"
	"fmt"

	"github.com/EmeraldLS/monogo/types"
	"github.com/EmeraldLS/monogo/utils"
)

// This endpoint returns all the inflow into an account grouped by months
func (m *MonoClient) GetCredits(accountID string) (*types.FlowResponse, error) {
	// https://api.withmono.com/accounts/{id}/credits
	path := fmt.Sprintf("/%s/%s/credits", accountBasePath, accountID)

	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var inflowRes types.FlowResponse
	err = json.NewDecoder(res.Body).Decode(&inflowRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	return &inflowRes, nil
}

// This endpoint returns all the outflow from an account grouped by months
func (o *MonoClient) GetDebits(accountID string) (*types.FlowResponse, error) {
	// https://api.withmono.com/accounts/{id}/credits
	path := fmt.Sprintf("/%s/%s/debits", accountBasePath, accountID)
	o.withPath(path)

	res, err := o.client.Do(o.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var inflowRes types.FlowResponse
	err = json.NewDecoder(res.Body).Decode(&inflowRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	return &inflowRes, nil
}
