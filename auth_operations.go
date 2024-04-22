package monogo

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/EmeraldLS/monogo/types"
	"github.com/EmeraldLS/monogo/utils"
)

func (m *MonoClient) Login(loginReq *types.AuthLoginRequest) (*types.AuthLoginResponse, error) {
	path := fmt.Sprintf("/%s/auth", telecomBasePath)
	jsonBytes, err := json.Marshal(loginReq)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(jsonBytes)

	res, err := m.withBody(buf).withMethod("POST").withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var loginRes types.AuthLoginResponse
	err = json.NewDecoder(res.Body).Decode(&loginRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	return &loginRes, nil
}

// This resource is to verify the OTP that is sent to the user's telco phone line.
func (m *MonoClient) VerifyOTP(otpReq *types.VerifyOTPRequest) (*types.VerifyOTPResponse, error) {
	path := fmt.Sprintf("/%s/verify", telecomBasePath)
	jsonBytes, err := json.Marshal(otpReq)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(jsonBytes)

	res, err := m.withBody(buf).withMethod("POST").withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var otpRes types.VerifyOTPResponse
	err = json.NewDecoder(res.Body).Decode(&otpReq)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	return &otpRes, nil
}

func (m *MonoClient) AuthTokenExchange(tokenReq *types.AuthTokenExchangeRequest) (*types.AuthTokenExchangeResponse, error) {
	path := "/account/auth"
	jsonBytes, err := json.Marshal(tokenReq)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(jsonBytes)
	res, err := m.withBody(buf).withMethod("POST").withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var tokenRes types.AuthTokenExchangeResponse
	err = json.NewDecoder(res.Body).Decode(&tokenRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	return &tokenRes, nil
}
