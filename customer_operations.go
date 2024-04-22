package monogo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/EmeraldLS/monogo/types"
	"github.com/EmeraldLS/monogo/utils"
)

/*
The customer creation endpoint allows businesses to seamlessly onboard new users by capturing essential user details and performing Know Your Customer (KYC) checks. This process ensures compliance with regulatory requirements and facilitates the efficient addition of users to your platform.
*/
func (m *MonoClient) CreateCustomer(requestBody *types.Customer) (*types.CustomerResponse, error) {
	jsonBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error occured marsahling request body: %v", err)
	}

	res, err := m.withBody(bytes.NewBuffer(jsonBytes)).withPath(customerBasePath).withMethod("POST").client.Do(m.req)

	if err != nil {
		return nil, utils.RequestError(err)
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Printf("error occured closing response body: %v\n", err)
		}
	}(m.req.Body)

	if res.StatusCode != http.StatusCreated {
		fmt.Println("Status = ", res.Status)
		return nil, errors.New(res.Status)
	}

	var customerResponse types.CustomerResponse
	err = json.NewDecoder(res.Body).Decode(&customerResponse)

	fmt.Println("Customer Response = ", customerResponse)

	if err != nil {
		return nil, utils.DecodeError(err)
	}

	return &customerResponse, nil
}

/*
This endpoint allows you to fetch details of a specific customer within your application ecosystem. By providing the customer ID, you can access their information, including personal details and transaction history.
*/
func (m *MonoClient) GetCustomer(customerID string) (*types.CustomerMininalResponse, error) {

	res, err := m.withPath(customerBasePath + "/" + customerID).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	var customerResponse types.CustomerMininalResponse

	if err := json.NewDecoder(res.Body).Decode(&customerResponse); err != nil {
		return nil, utils.DecodeError(err)
	}

	return &customerResponse, nil
}

// This resource is to retrieve information about all customers.
func (m *MonoClient) GetAllCustomers() (*types.CustomersMinimalResponse, error) {

	res, err := m.client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var customersResponse types.CustomersMinimalResponse

	if err := json.NewDecoder(res.Body).Decode(&customersResponse); err != nil {
		return nil, utils.DecodeError(err)
	}

	return &customersResponse, nil
}

// This resource retrieves all transactions performed by a customer via our payments products
func (m *MonoClient) GetCustomerTransactions(customerID, period, accountID string, perPage, page int) (*types.CustomerTransactionsResponse, error) {
	path := fmt.Sprintf("%s/%s/transactions?per_page=%d&period=%s&page=%d&account=%s", customerBasePath, customerID, perPage, period, page, accountID)

	res, err := m.withPath(path).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var transactionsResponse *types.CustomerTransactionsResponse

	if err = json.NewDecoder(res.Body).Decode(&transactionsResponse); err != nil {
		return nil, utils.DecodeError(err)
	}

	return transactionsResponse, nil
}

// This resource helps to filter and retrieve information about a single customer
// TODO: Will come back to this function, something isnt right about it
func (m *MonoClient) FilterCustomer(customerID, page, phone, identity, email string) (*types.CustomerMininalResponse, error) {
	url := fmt.Sprintf("/%s/%s/%s/%s/%s/%s", customerBasePath, customerID, page, phone, identity, email)

	res, err := m.withPath(url).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var customerResponse types.CustomerMininalResponse
	if err := json.NewDecoder(res.Body).Decode(&customerResponse); err != nil {
		return nil, utils.DecodeError(err)
	}

	return &customerResponse, nil

}

// This resource provides all accounts linked to your business.
func (m *MonoClient) LinkedAccounts(customerId string) (*types.LinkedAccountsResponse, error) {

	res, err := m.withPath("/accounts/" + customerId).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var linkedAccountsResponse types.LinkedAccountsResponse
	if err := json.NewDecoder(res.Body).Decode(&linkedAccountsResponse); err != nil {
		return nil, utils.DecodeError(err)
	}
	return &linkedAccountsResponse, nil
}

func (m *MonoClient) UpdateCustomer(customerId string, customer *types.CustomerUpdateRequest) (*types.CustomerUpdateResponse, error) {
	path := fmt.Sprintf("/%s/%s", customerBasePath, customerId)

	jsonBytes, err := json.Marshal(customer)
	if err != nil {
		return nil, fmt.Errorf("error occured marsahling request body: %v", err)
	}

	buf := bytes.NewBuffer(jsonBytes)

	res, err := m.withMethod("PATCH").withPath(path).withBody(buf).client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}
	var updateResponse types.CustomerUpdateResponse
	if err = json.NewDecoder(res.Body).Decode(&updateResponse); err != nil {
		return nil, utils.DecodeError(err)
	}

	return &updateResponse, nil
}

func (m *MonoClient) DeleteCustomer(customerID string) (*types.CustomerUpdateResponse, error) {
	urlPath := fmt.Sprintf("/%s/%s", customerBasePath, customerID)

	res, err := m.withPath(urlPath).withMethod("DELETE").client.Do(m.req)
	if err != nil {
		return nil, utils.RequestError(err)
	}

	var deleteResponse types.CustomerUpdateResponse
	if err := json.NewDecoder(res.Body).Decode(&deleteResponse); err != nil {
		return nil, utils.DecodeError(err)
	}

	return &deleteResponse, nil
}
