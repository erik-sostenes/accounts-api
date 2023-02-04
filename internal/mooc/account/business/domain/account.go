package domain

// Account(Domain Object) represents a user account
type Account struct {
	accountId       AccountId
	accountUserName AccountUserName
	accountName     AccountName
	accountLastName AccountLastName
	accountEmail    AccountEmail
	accountPassword AccountPassword
	accountCareer   AccountCareer
	accountIP       AccountIP
	accountActive   AccountActive
	accountDetails  AccountDetails
}

// NewAccount takes primitive values and converts them into value objects that make up an Account
func NewAccount(id, username, name, lastName, email, password, career, ip, active string,
	details Map,
) (Account, error) {

	accounId, err := NewAccountId(id)
	if err != nil {
		return Account{}, err
	}

	accountUserName, err := NewAccountUserName(username)
	if err != nil {
		return Account{}, err
	}

	accountName, err := NewAccountName(name)
	if err != nil {
		return Account{}, err
	}

	accountLastName, err := NewAccountLastName(lastName)
	if err != nil {
		return Account{}, err
	}

	accountEmail, err := NewAccountEmail(email)
	if err != nil {
		return Account{}, err
	}

	accountPassword, err := NewAccountPassword(password)
	if err != nil {
		return Account{}, err
	}

	accountCareer, err := NewAccountCareer(career)
	if err != nil {
		return Account{}, err
	}

	accountIP, err := NewAccountIP(ip)
	if err != nil {
		return Account{}, err
	}

	accountActive, err := NewAccounActive(active)
	if err != nil {
		return Account{}, err
	}

	accountDetails, err := NewAccountDetails(details)
	if err != nil {
		return Account{}, err
	}

	return Account{
		accounId,
		accountUserName,
		accountName,
		accountLastName,
		accountEmail,
		accountPassword,
		accountCareer,
		accountIP,
		accountActive,
		accountDetails,
	}, nil
}

func (a *Account) AccountId() AccountId {
	return a.accountId
}

func (a *Account) AccountUserName() AccountUserName {
	return a.accountUserName
}

func (a *Account) AccountName() AccountName {
	return a.accountName
}

func (a *Account) AccountLastName() AccountLastName {
	return a.accountLastName
}

func (a *Account) AccountEmail() AccountEmail {
	return a.accountEmail
}

func (a *Account) AccountPassword() AccountPassword {
	return a.accountPassword
}

func (a *Account) AccountCareer() AccountCareer {
	return a.accountCareer
}

func (a *Account) AccountIP() AccountIP {
	return a.accountIP
}

func (a *Account) AccountActive() AccountActive {
	return a.accountActive
}

func (a *Account) AccountDetails() AccountDetails {
	return a.accountDetails
}
