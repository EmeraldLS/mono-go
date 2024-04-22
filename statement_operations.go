package monogo

import (
	"encoding/json"
	"fmt"

	"github.com/EmeraldLS/monogo/types"
	"github.com/EmeraldLS/monogo/utils"
)

// This resource represents the bank statement of the connected financial account. You can query 1-12 months bank statement in one single call.
// Period: In months( 1-12) e.g last12months
// Output: You can set the output as pdf if you want to receive pdf instead of Json
// Format: This expects format i.e. v2
func (m *MonoClient) GetAccountStatements(accountID, period, output, format string) (interface{}, error) {
	path := fmt.Sprintf("%s/%s/statement?period=%s&output=%sformat=%s", accountBasePath, accountID, period, output, format)

	dataSync := m.req.Header.Get("x-realtime")
	if dataSync == "" {
		m.req.Header.Set("x-realtime", "true")
	}

	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	switch output {
	case "pdf":
		var pdfResp types.AccountStatementPDFResponse
		err = json.NewDecoder(res.Body).Decode(&pdfResp)
		if err != nil {
			return nil, utils.DecodeError(err)
		}

		return pdfResp, nil
	case "json":
		var jsonResp types.AccountStatementsResponse
		err = json.NewDecoder(res.Body).Decode(&jsonResp)
		if err != nil {
			return nil, utils.DecodeError(err)
		}

		return jsonResp, nil
	default:
		var jsonResp types.AccountStatementsResponse
		err = json.NewDecoder(res.Body).Decode(&jsonResp)
		if err != nil {
			return nil, utils.DecodeError(err)
		}

		return jsonResp, nil
	}
}

// If you set the output as PDF, you can use this endpoint to poll the status
// Account ID returned from token exchange
// ID returned from statements API
func (m *MonoClient) GetPDFStatus(accountID, jobID string) (*types.AccountStatementPDFResponse, error) {
	// https://api.withmono.com/v2/accounts/{id}/statement/jobs/{jobId}
	path := fmt.Sprintf("%s/%s/statement/jobs/%s", accountBasePath, accountID, jobID)

	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}
	var pdfResp *types.AccountStatementPDFResponse
	if err := json.NewDecoder(res.Body).Decode(&pdfResp); err != nil {
		return nil, utils.DecodeError(err)
	}

	return pdfResp, nil
}
