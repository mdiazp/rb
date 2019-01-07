package pmessagefordr

import (
	"net/http"

	"github.com/mdiazp/rb/server/api"
	"github.com/mdiazp/rb/server/api/controllers"
	dbhandlers "github.com/mdiazp/rb/server/db/handlers"
)

// RetrieveController ...
type RetrieveController interface {
	controllers.BaseController
}

// NewRetrieveController ...
func NewRetrieveController(base api.Base) RetrieveController {
	return &retrieveController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type retrieveController struct {
	api.Base
}

func (c *retrieveController) GetRoute() string {
	return "/pmessagefordr/{pmessageID}"
}

func (c *retrieveController) GetMethods() []string {
	return []string{"GET"}
}

// GetAccess ...
func (c *retrieveController) GetAccess() controllers.Permission {
	return controllers.PermissionRetrievePMessageForDR
}

// ServeHTTP ...
func (c *retrieveController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := c.GetPInt(w, r, "pmessageID")

	o, e := c.DB().RetrievePMessageForDRByID(id)
	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
	c.WR(w, 200, *o)
}
