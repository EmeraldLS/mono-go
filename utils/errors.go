package utils

import "fmt"

// func TestError(err error, data interface{}) (interface{}, error) {
// 	if err != nil {
// 		return nil, fmt.Errorf("error occured marshalling request body: %v", err)
// 	}

// 	return data, nil
// }

func MarshalError(err error) error {
	return fmt.Errorf("error occured marshalling request body: %v", err)
}

func CreateRequestError(err error) error {
	return fmt.Errorf("error occured creating request: %v", err)
}

func RequestError(err error) error {
	return fmt.Errorf("error occured sending request: %v", err)
}

func DecodeError(err error) error {
	return fmt.Errorf("error occured receiving response: %v", err)
}
