package monogo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/EmeraldLS/monogo/types"
	"github.com/EmeraldLS/monogo/utils"
)

/*
Use this resource to initiate a BVN Consent request

Scope:
Kindly note that by default, scope is set to IDENTITY. To retrieve all accounts connected to a BVN,BANK_ACCOUNTS must be passed as scope.
*/
func (m *MonoClient) InitiateBVNLookup(data *types.BVNLookupRequest) (*types.BVNLookupResponse, error) {
	path := fmt.Sprintf("/%s/bvn/initiate", lookupBasePath)
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(jsonBytes)

	res, err := m.withPath(path).withMethod("POST").withBody(buf).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var bvnLookupRes types.BVNLookupResponse
	err = json.NewDecoder(res.Body).Decode(&bvnLookupRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	return &bvnLookupRes, nil

}

// Use this resource to verify the BVN request via OTP
func (m *MonoClient) BVNVerifyOTP(body *types.BVNVerifyOTPRequest, sessionID string) (*types.BVNVerifyOTPResponse, error) {
	path := fmt.Sprintf("/%s/bvn/verify", lookupBasePath)
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(jsonBytes)

	m.req.Header.Set("x-session-id", sessionID)
	res, err := m.withPath(path).withMethod("POST").withBody(buf).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var bvnVerifyRes types.BVNVerifyOTPResponse
	err = json.NewDecoder(res.Body).Decode(&bvnVerifyRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	return &bvnVerifyRes, nil
}

func (m *MonoClient) GetBVNDetails(data *types.BVNDetailsRequest, sessionID string) (*types.BVNDetailsResponse, *types.BVNBankAccountsResponse, error) {
	path := fmt.Sprintf("/%s/bvn/details", lookupBasePath)
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, nil, utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(jsonBytes)

	headerSessionID := m.req.Header.Get("x-session-id")
	if headerSessionID == "" {
		m.req.Header.Set("x-session-id", sessionID)
	}

	res, err := m.withBody(buf).withPath(path).withMethod("POST").client.Do(m.req)
	if err != nil {
		return nil, nil, utils.RequestError(err)
	}

	if res.StatusCode == http.StatusCreated {
		var bankAccRes types.BVNBankAccountsResponse
		err = json.NewDecoder(res.Body).Decode(&bankAccRes)
		if err != nil {
			return nil, nil, utils.DecodeError(err)
		}
		return nil, &bankAccRes, nil
	}

	if res.StatusCode == http.StatusOK {
		var bvnDetailsRes types.BVNDetailsResponse
		err = json.NewDecoder(res.Body).Decode(&bvnDetailsRes)
		if err != nil {
			return nil, nil, utils.DecodeError(err)
		}

		return &bvnDetailsRes, nil, nil
	}

	return nil, nil, errors.New("unexpected error has occured")
}
