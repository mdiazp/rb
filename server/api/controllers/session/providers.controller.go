package session

import (
	"net/http"

	"github.com/mdiazp/rb/server/api"
	"github.com/mdiazp/rb/server/api/controllers"
)

// ProvidersController ...
type ProvidersController interface {
	controllers.BaseController
}

// NewProvidersController ...
func NewProvidersController(base api.Base) ProvidersController {
	return &providersController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type providersController struct {
	api.Base
}

func (c *providersController) GetRoute() string {
	return "/authproviders"
}

func (c *providersController) GetMethods() []string {
	return []string{"GET"}
}

// GetAccess ...
func (c *providersController) GetAccess() controllers.Permission {
	return ""
}

// ServeHTTP ...
func (c *providersController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.WR(w, 200, c.GetAuthProviderNames())
}
