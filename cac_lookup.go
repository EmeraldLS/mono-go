package monogo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EmeraldLS/monogo/types"
	"github.com/EmeraldLS/monogo/utils"
)

// Use this resource to verify the existence of a business
func (m *MonoClient) LookupBusiness(RCNumber string) (*types.LookupBusinessResponse, error) {
	path := fmt.Sprintf("/%s/cac?search=%s", lookupBasePath, RCNumber)
	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var businessRes types.LookupBusinessResponse
	err = json.NewDecoder(res.Body).Decode(&businessRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error occured: %v", businessRes.Message)
	}

	return &businessRes, nil
}

// Use this resource to retrieve shareholder information of a business
func (m *MonoClient) LookupShareholderDetails(businessID string) (*types.ShareholderDetailsResponse, error) {
	path := fmt.Sprintf("/%s/cac/company/%s", lookupBasePath, businessID)
	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var shareHolderResponse types.ShareholderDetailsResponse
	err = json.NewDecoder(res.Body).Decode(&shareHolderResponse)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error occured: %v", shareHolderResponse.Message)
	}

	return &shareHolderResponse, nil
}

// Use this resource to check previous company address
func (m *MonoClient) LookupPreviousAddress(businessID string) (*types.PreviousBusinessAddressResponse, error) {
	path := fmt.Sprintf("/%s/cac/company/%s", lookupBasePath, businessID)
	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var prevAddressRes types.PreviousBusinessAddressResponse
	if err = json.NewDecoder(res.Body).Decode(&prevAddressRes); err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error occured: %v", prevAddressRes.Message)
	}

	return &prevAddressRes, nil
}

func (m *MonoClient) LookupChangeOfName(businessID string) (*types.ChangeOfNameResponse, error) {
	path := fmt.Sprintf("/%s/cac/company/%s/change-of-name", lookupBasePath, businessID)
	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var nameRes types.ChangeOfNameResponse
	err = json.NewDecoder(res.Body).Decode(&nameRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error occured: %v", nameRes.Message)
	}

	return &nameRes, nil
}

// Use this resource to search for the company secretary
func (m *MonoClient) LookupSecretary(businessID string) (*types.SecretResponse, error) {
	path := fmt.Sprintf("/%s/cac/company/%s/secretary", lookupBasePath, businessID)
	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var secRes types.SecretResponse
	err = json.NewDecoder(res.Body).Decode(&secRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error has occured: %v", secRes.Message)
	}

	return &secRes, nil
}

// Use this resource to search for the company directors
func (m *MonoClient) LookupDirectors(businessID string) (*types.DirectorsResponse, error) {
	path := fmt.Sprintf("/%s/cac/company/%s/directors", lookupBasePath, businessID)
	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var direcRes types.DirectorsResponse
	err = json.NewDecoder(res.Body).Decode(&direcRes)
	if err != nil {
		return nil, utils.DecodeError(err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error has occured: %v", direcRes.Message)
	}

	return &direcRes, nil
}
