package pdiskreservation

import (
	"fmt"
	"net/http"

	"github.com/mdiazp/rb/server/api"
	"github.com/mdiazp/rb/server/api/controllers"
	"github.com/mdiazp/rb/server/core"
	"github.com/mdiazp/rb/server/db/models"
)

// CreateController ...
type CreateController interface {
	controllers.BaseController
}

// NewCreateController ...
func NewCreateController(base api.Base) CreateController {
	return &createController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type createController struct {
	api.Base
}

func (c *createController) GetRoute() string {
	return "/pdiskreservation"
}

func (c *createController) GetMethods() []string {
	return []string{"POST"}
}

// GetAccess ...
func (c *createController) GetAccess() controllers.Permission {
	return controllers.PermissionCreatePDiskReservation
}

// ServeHTTP ...
func (c *createController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------->  CreatePDiskReservation init")

	var o models.PDiskReservation
	c.ReadJSON(w, r, &o)
	o.ID = 0

	fmt.Println("---------> finish to ReadJSON ")

	c.WE400(w, o.Valid())

	o.InitialTime = core.MoveDateToNextWeekDay(o.InitialTime, o.TurnWeekDay)
	o.FinishTime = core.MoveDateToPreviousWeekDay(o.FinishTime, o.TurnWeekDay)

	c.WE400(w, o.Valid())

	e := c.DB().CreatePDiskReservation(&o)

	c.WE(w, e, 500)
	c.WR(w, 200, o)
}
