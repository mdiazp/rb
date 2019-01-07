package pdiskreservation

import (
	"net/http"

	"github.com/mdiazp/rb/server/api"
	"github.com/mdiazp/rb/server/api/controllers"
)

// DeleteController ...
type DeleteController interface {
	controllers.BaseController
}

// NewDeleteController ...
func NewDeleteController(base api.Base) DeleteController {
	return &deleteController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type deleteController struct {
	api.Base
}

func (c *deleteController) GetRoute() string {
	return "/pdiskreservation/{pserviceID}"
}

func (c *deleteController) GetMethods() []string {
	return []string{"DELETE"}
}

// GetAccess ...
func (c *deleteController) GetAccess() controllers.Permission {
	return controllers.PermissionDeletePDiskReservation
}

// ServeHTTP ...
func (c *deleteController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := c.GetPInt(w, r, "pserviceID")

	verificatePDiskReservationExistence(c, w, id)

	e := c.DB().DeletePDiskReservation(id)
	c.WE(w, e, 500)
	c.WR(w, 204, "PDiskReservation was deleted successfully")
}
