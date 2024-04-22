package monogo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/EmeraldLS/monogo/types"
	"github.com/EmeraldLS/monogo/utils"
)

/*
Manual Data Sync has a rate limit of one API call per connected account every two minutes.

If Manual Data Sync returns INCOMPLETE_STATEMENT_ERROR, you can set the query param:allow_incomplete_statement to true, to return the latest data which is less than the recorded transactions of the user when they initially linked their account, or when Data sync was last initiated.
*/
func (m *MonoClient) ManualDataSync(accountID string) (*types.DataSyncResponse, *types.DataSyncErrorResponse, error) {
	path := fmt.Sprintf("/%s/%s/sync", accountBasePath, accountID)

	res, err := m.withPath(path).withMethod("POST").client.Do(m.req)
	if err != nil {
		return nil, nil, utils.RequestError(err)
	}

	var errRes types.DataSyncErrorResponse
	if res.StatusCode != http.StatusOK {
		err = json.NewDecoder(res.Body).Decode(&errRes)
		if err != nil {
			return nil, nil, utils.DecodeError(err)
		}
		return nil, &errRes, errors.New("error occured data syncing")
	}

	var dataSyncRes types.DataSyncResponse
	err = json.NewDecoder(res.Body).Decode(&dataSyncRes)
	if err != nil {
		return nil, nil, utils.DecodeError(err)
	}

	return &dataSyncRes, nil, nil
}

/*
Reauth code is a mono generated code for the account you want to re-authenticate, which must be requested by your server and sent to your frontend where you can pass it to mono connect widget.

Mono connect widget will ask for the required information and re-authenticate the user's account and notify your server.
*/
func (o *MonoClient) ReauthCode(accountID string) (*types.ReauthCodeResponse, error) {
	path := fmt.Sprintf("/%s/%s/reauthorize", accountBasePath, accountID)

	res, err := o.withPath(path).withMethod("POST").client.Do(o.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var reauthCodeRes types.ReauthCodeResponse
	err = json.NewDecoder(res.Body).Decode(&reauthCodeRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	return &reauthCodeRes, nil
}
