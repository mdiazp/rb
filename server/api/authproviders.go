package api

import (
	"github.com/mdiazp/rb/server/authprovider"
	"github.com/mdiazp/rb/server/authprovider/xxx"
)

// AuthProvider ...
type AuthProvider string

const (
	// AuthProviderXXX ...
	AuthProviderXXX AuthProvider = "XXX"
)

// GetAuthproviderNames ...
func (b *base) GetAuthProviderNames() []AuthProvider {
	return []AuthProvider{AuthProviderXXX}
}

// Getauthprovider ...
func (b *base) GetAuthProvider(provider AuthProvider) authprovider.Provider {
	switch provider {
	case AuthProviderXXX:
		if b.GetEnv() == "dev" {
			return xxx.GetProvider()
		}
		return nil
	default:
		return nil
	}
}
