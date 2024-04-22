package monogo

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/EmeraldLS/monogo/types"
	"github.com/EmeraldLS/monogo/utils"
)

// This resource represents the known transactions on the account.
func (m *MonoClient) GetTransactions(accountID string, params types.AccountTransactionParams) (*types.AccountTransactionResponse, error) {

	path := fmt.Sprintf("%s/%s/transactions?start=%s&end=%s&narration=%s&type=%s&paginate=%s&limits=%d", accountBasePath, accountID, params.Start, params.End, params.Narration, params.Type, strconv.FormatBool(params.Paginate), params.Limits)

	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var transactionResp types.AccountTransactionResponse
	err = json.NewDecoder(res.Body).Decode(&transactionResp)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	return &transactionResp, nil
}

// TODO: work on transaction Metadata
