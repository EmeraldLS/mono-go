package monogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EmeraldLS/monogo/types"
	"github.com/EmeraldLS/monogo/utils"
)

// This resource returns NIP supported bank coverage.
func (m *MonoClient) BanksListing() (*types.BanksListingResponse, error) {
	path := fmt.Sprintf("/%s/banks", lookupBasePath)
	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var banks types.BanksListingResponse
	err = json.NewDecoder(res.Body).Decode(&banks)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	return &banks, nil
}

// This resource verifies your address via your meter number and house address.
func (m *MonoClient) HouseAddressVerification(body *types.HouseAddressVerificationRequest) (*types.HouseAddressVerificationResponse, error) {
	path := fmt.Sprintf("/%s/address", lookupBasePath)
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(jsonBytes)

	res, err := m.withBody(buf).withMethod("POST").withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}
	var verifRes types.HouseAddressVerificationResponse
	err = json.NewDecoder(res.Body).Decode(&verifRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error occured: %v", verifRes.Message)
	}

	return &verifRes, nil
}

// This resource verifies an international passport document via its passport number and last name.
func (m *MonoClient) LookupInternationalPassport(body *types.InternationalPassportRequest) (*types.InternationalPassportResponse, error) {
	path := fmt.Sprintf("/%s/passport", lookupBasePath)
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(jsonBytes)

	res, err := m.withBody(buf).withMethod("POST").withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var passportRes types.InternationalPassportResponse
	err = json.NewDecoder(res.Body).Decode(&passportRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error occured: %s", passportRes.Message)
	}

	return &passportRes, nil
}

// This resource verifies the tax identification number of an entity.
func (m *MonoClient) LookupTIN(body *types.TINRequest) (*types.TINResponse, error) {
	path := fmt.Sprintf("/%s/tin", lookupBasePath)
	b, err := json.Marshal(body)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(b)

	res, err := m.withPath(path).withBody(buf).withMethod("POST").client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var tinRes types.TINResponse
	err = json.NewDecoder(res.Body).Decode(&tinRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error occured: %s", tinRes.Message)
	}

	return &tinRes, nil

}

// This resource verifies the national identification number of a user.
func (m *MonoClient) LookupNIN(body *types.NINRequest) (*types.NINResponse, error) {
	path := fmt.Sprintf("/%s/nin", lookupBasePath)
	b, err := json.Marshal(body)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(b)

	res, err := m.withBody(buf).withMethod("POST").withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var ninRes types.NINResponse
	err = json.NewDecoder(res.Body).Decode(&ninRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error occured: %s", ninRes.Message)
	}

	return &ninRes, nil
}

// This resource verifies the driver license number of an individual.
func (m *MonoClient) LookupDriverLicense(body *types.DriverLicenseRequest) (*types.DriverLicenseResponse, error) {
	path := fmt.Sprintf("/%s/driver_license", lookupBasePath)
	b, err := json.Marshal(body)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(b)

	res, err := m.withBody(buf).withMethod("POST").withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var licenseRes types.DriverLicenseResponse
	err = json.NewDecoder(res.Body).Decode(&licenseRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error occured: %s", licenseRes.Message)
	}

	return &licenseRes, nil
}

// This resource verifies the account and returns the masked BVN attached to the account number supplied.
func (m *MonoClient) LookupAccountNumber(body *types.AccountNumberRequest) (*types.AccountNumberResponse, error) {
	path := fmt.Sprintf("/%s/account-number", lookupBasePath)
	b, err := json.Marshal(body)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(b)

	res, err := m.withBody(buf).withPath(path).withMethod("POST").client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var accNumRes types.AccountNumberResponse
	err = json.NewDecoder(res.Body).Decode(&accNumRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error occured: %s", accNumRes.Message)
	}

	return &accNumRes, nil
}

/*
This resource enables you to retrieve a user's credit history.

CRC: represents Credit Bureau Limited.

XDS: represents First Central Credit Bureau.

ALL: represents Credit history data from both Bureaus.
*/
func (m *MonoClient) LookupCreditHistory(provider string, body *types.CreditHistoryRequest) (*types.CreditHistoryResponse, error) {
	path := fmt.Sprintf("/%s/credit-history/%s", lookupBasePath, provider)
	b, err := json.Marshal(body)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(b)

	res, err := m.withBody(buf).withMethod("POST").withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var creditHisRes types.CreditHistoryResponse
	err = json.NewDecoder(res.Body).Decode(&creditHisRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error occured: %v", creditHisRes.Message)
	}

	return &creditHisRes, nil
}

// This resource allows you to verify the NIN, BVN and date of birth of your user in one API call for KYC.
func (m *MonoClient) LookupMarshup(body *types.MarshupRequest) (*types.MarshupResponse, error) {
	path := fmt.Sprintf("/%s/marshup", lookupBasePath)
	b, err := json.Marshal(body)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(b)

	res, err := m.withBody(buf).withMethod("POST").withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var marshupRes types.MarshupResponse
	err = json.NewDecoder(res.Body).Decode(&marshupRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unknown error occured: %s", marshupRes.Message)
	}

	return &marshupRes, nil

}
