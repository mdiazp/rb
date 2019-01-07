package xxx

import (
	"errors"

	"github.com/mdiazp/rb/server/authprovider"
)

type provider struct{}

// Authenticate ...
func (p *provider) Authenticate(username, password string) error {
	if password != "123" {
		return errors.New("Fail Authentication")
	}
	return nil
}

func (p *provider) GetUserRecords(username string) (authprovider.UserRecords, error) {
	return authprovider.UserRecords{
		Username: username,
		Name:     username,
	}, nil
}

// GetProvider ...
func GetProvider() authprovider.Provider {
	return &provider{}
}
