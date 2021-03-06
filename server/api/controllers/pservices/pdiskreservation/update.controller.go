package pdiskreservation

import (
	"net/http"

	"github.com/mdiazp/rb/server/api"
	"github.com/mdiazp/rb/server/api/controllers"
	"github.com/mdiazp/rb/server/core"
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
	return "/pdiskreservation/{pserviceID}"
}

func (c *updateController) GetMethods() []string {
	return []string{"PATCH"}
}

// GetAccess ...
func (c *updateController) GetAccess() controllers.Permission {
	return controllers.PermissionUpdatePDiskReservation
}

// ServeHTTP ...
func (c *updateController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := c.GetPInt(w, r, "pserviceID")
	var o models.PDiskReservation
	c.ReadJSON(w, r, &o)
	o.ID = id

	verificatePDiskReservationExistence(c, w, id)

	c.WE400(w, o.Valid())

	o.InitialTime = core.MoveDateToNextWeekDay(o.InitialTime, o.TurnWeekDay)
	o.FinishTime = core.MoveDateToPreviousWeekDay(o.FinishTime, o.TurnWeekDay)

	c.WE400(w, o.Valid())

	e := c.DB().UpdatePDiskReservation(&o)
	c.WE(w, e, 500)
	c.WR(w, 200, o)
}
