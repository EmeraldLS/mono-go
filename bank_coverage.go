package monogo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EmeraldLS/monogo/types"
	"github.com/EmeraldLS/monogo/utils"
)

func (m *MonoClient) BankCoverage(params *types.BankCoverageParams) (*types.BankCoverageResponse, error) {
	path := "/v3/institutions"
	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var bankCoverageRes types.BankCoverageResponse
	err = json.NewDecoder(res.Body).Decode(&bankCoverageRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error occured: %s", bankCoverageRes.Message)
	}

	return &bankCoverageRes, nil

}
