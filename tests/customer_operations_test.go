package tests

import (
	"testing"

	"github.com/EmeraldLS/monogo"
	"github.com/EmeraldLS/monogo/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {

	customerOperation := monogo.New("test_sk_yfswm1978ftt4r7oxhri") // test key. haha

	customer := &types.Customer{
		Email:       "john@m.com",
		FirsNname:   "John",
		LastName:    "Doe",
		AddressLine: "No 1, John Doe Street",
		Phone:       "08012345678",
		Identity: types.Identification{
			Type:   "BVN",
			Number: "12345678901",
		},
	}

	resp, err := customerOperation.CreateCustomer(customer)

	assert.Nil(t, err)

	assert.Equal(t, "success", resp.Status)
}
