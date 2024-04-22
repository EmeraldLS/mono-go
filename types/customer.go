package types

/*
Package types provides data structures for managing customer information.
*/

// The Customer struct represents a customer and their details, including email, first name, last name, address, phone number, and identification information.
type Customer struct {
	Email       string         `json:"email"`
	FirsNname   string         `json:"firstName"`
	LastName    string         `json:"lastName"`
	AddressLine string         `json:"address"`
	Phone       string         `json:"phone"`
	Identity    Identification `json:"identity"`
}

type CustomerMininalData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// The Identification struct represents the type and number of identification documents.
type Identification struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

// The CustomerResponse struct represents the response received when interacting with the customer API, including the status, message, and customer data.
type CustomerResponse struct {
	Status  string               `json:"status"`
	Message string               `json:"Message"`
	Data    CustomerResponseData `json:"data"`
}

// The CustomerResponseData struct represents the customer data returned in the response, including the customer ID, name, email, address, identification number, identification type, and BVN.
type CustomerResponseData struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Email              string `json:"email"`
	AddressLine1       string `json:"address_line_1"`
	AddressLine2       string `json:"address_line_2"`
	IdentificationNo   string `json:"identification_no"`
	IdentificationType string `json:"identification_type"`
	BVN                string `json:"bvn"`
}

// The CustomerMinimalResponse struct represents a minimal response received when interacting with the customer API, including the status, message, and minimal customer data.
type CustomerMininalResponse struct {
	Status  string                      `json:"status"`
	Message string                      `json:"message"`
	Data    CustomerMininalResponseData `json:"data"`
}

// The CustomersMinimalResponse struct represents a list of minimal customer responses received when interacting with the customer API, including the status, message, and an array of minimal customer data.
type CustomersMinimalResponse struct {
	Status  string                         `json:"status"`
	Message string                         `json:"message"`
	Data    []*CustomerMininalResponseData `json:"data"`
}

// The CustomerMinimalResponseData struct represents the minimal customer data returned in the response, including the customer ID, name, address, identification number, identification type, and BVN.
type CustomerMininalResponseData struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Address            string `json:"address"`
	IdentificationNo   string `json:"identification_no"`
	IdentificationType string `json:"identification_type"`
	BVN                string `json:"bvn"`
}

type CustomerUpdateRequest struct {
	Identity Identification `json:"identity"`
	Address  string         `json:"address"`
	Phone    string         `json:"phone"`
}

type CustomerUpdateResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    struct{} `json:"data"`
}
