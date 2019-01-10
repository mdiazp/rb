package freeinfo

import (
	"net/http"

	"github.com/mdiazp/rb/server/api"
	"github.com/mdiazp/rb/server/api/controllers"
	"github.com/mdiazp/rb/server/db/models"
)

// RetrieveTurnNumsController ...
type RetrieveTurnNumsController interface {
	controllers.BaseController
}

// NewRetrieveTurnNumsController ...
func NewRetrieveTurnNumsController(base api.Base) RetrieveTurnNumsController {
	return &retrieveTurnNumsController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type retrieveTurnNumsController struct {
	api.Base
}

func (c *retrieveTurnNumsController) GetRoute() string {
	return "/turn-nums"
}

func (c *retrieveTurnNumsController) GetMethods() []string {
	return []string{"GET"}
}

// GetAccess ...
func (c *retrieveTurnNumsController) GetAccess() controllers.Permission {
	return ""
}

// ServeHTTP ...
func (c *retrieveTurnNumsController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.WR(w, 200, models.GetTurnNums())
}
