package domain

import "encoding/json"

// Map represents a map for working with data dynamically.
type Map map[string]any

// AccountFunc create a domain object Account
type AccountFunc func() (Account, error)

// Account executes the function as long as it complies with the AccountFunc signature
func (f AccountFunc) Account() (Account, error) {
	return f()
}

// NewFirstAccount returns an Account
//
// is used for unit testing
func NewFirstAccount() AccountFunc {
	return func() (Account, error) {
		return NewAccount(
			"94343721-6baa-4cd5-a0b4-6c5d0419c02d",
			"JaredNV",
			"Jared Nicolas V",
			"Mitchell",
			"jared.gibson@gmail.com",
			"secret",
			"ISIC",
			"192.168.10.0",
			"true",
			Map{"permissions": []int{1, 2, 3}},
		)
	}
}

// NewTwoAccount returns an Account
//
// is used for unit testing
func NewTwoAccount() AccountFunc {
	return func() (Account, error) {
		return NewAccount(
			"94343721-6baa-4cd5-a0b4-6c5d0419c024",
			"JaredNV",
			"Jared Nicolas V",
			"Mitchell",
			"jared.gibson@gmail.com",
			"secret",
			"ISIC",
			"192.168.10.0",
			"true",
			Map{"permissions": []int{1, 2, 3}},
		)
	}
}

// BinaryMarshalerFunc binary marshal all data types
type BinaryMarshalerFunc func() ([]byte, error)

// MarshalBinary executes the function as long as it complies with the BinaryMarshalerFunc signature
func (f BinaryMarshalerFunc) MarshalBinary() ([]byte, error) {
	return f()
}

// NewMarshalJSON marshal to a binary json the data
func NewMarshalJSON(v any) BinaryMarshalerFunc {
	return func() ([]byte, error) {
		return json.Marshal(v)
	}
}
