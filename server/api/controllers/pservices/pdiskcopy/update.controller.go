package pdiskcopy

import (
	"net/http"

	"github.com/mdiazp/rb/server/api"
	"github.com/mdiazp/rb/server/api/controllers"
	"github.com/mdiazp/rb/server/db/models"
)

// UpdateController ...
type UpdateController interface {
	controllers.BaseController
}

// NewUpdateController ...
func NewUpdateController(base api.Base) UpdateController {
	return &updateController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type updateController struct {
	api.Base
}

func (c *updateController) GetRoute() string {
	return "/pdiskcopy/{pserviceID}"
}

func (c *updateController) GetMethods() []string {
	return []string{"PATCH"}
}

// GetAccess ...
func (c *updateController) GetAccess() controllers.Permission {
	return controllers.PermissionUpdatePDiskCopy
}

// ServeHTTP ...
func (c *updateController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := c.GetPInt(w, r, "pserviceID")
	var o models.PDiskCopy
	c.ReadJSON(w, r, &o)
	o.ID = id

	verificatePDiskCopyExistence(c, w, id)

	c.WE400(w, o.Valid())

	e := c.DB().UpdatePDiskCopy(&o)
	c.WE(w, e, 500)
	c.WR(w, 200, o)
}
