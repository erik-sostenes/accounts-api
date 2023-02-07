package domain

// AccountDetails(Value Object) represent the account details
type AccountDetails struct {
	value Map
}

func NewAccountDetails(value Map) (AccountDetails, error) {
	return AccountDetails{
		value: value,
	}, nil
}

func (a AccountDetails) Value() Map {
	return a.value
}
