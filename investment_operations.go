package monogo

import (
	"encoding/json"
	"fmt"

	"github.com/EmeraldLS/monogo/types"
	"github.com/EmeraldLS/monogo/utils"
)

// This resource represents the assets of your connected customers in their investment accounts.
func (m *MonoClient) GetAssets(accountID string) (*types.AssetsResponse, error) {
	path := fmt.Sprintf("/%s/%s/assets", accountBasePath, accountID)

	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var assetsRes types.AssetsResponse
	if err := json.NewDecoder(res.Body).Decode(&assetsRes); err != nil {
		return nil, utils.DecodeError(err)
	}

	return &assetsRes, nil
}

// This resource represents the earnings of your connected customers in their investment accounts.
func (m *MonoClient) GetEarnings(accountID string) (*types.EarningsResponse, error) {
	path := fmt.Sprintf("/%s/%s/earnings", accountBasePath, accountID)

	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var earningsResp types.EarningsResponse
	err = json.NewDecoder(res.Body).Decode(&earningsResp)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	return &earningsResp, nil
}
