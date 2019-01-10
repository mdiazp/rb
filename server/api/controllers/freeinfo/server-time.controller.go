package freeinfo

import (
	"net/http"
	"time"

	"github.com/mdiazp/rb/server/api"
	"github.com/mdiazp/rb/server/api/controllers"
)

// ServerTimeController ...
type ServerTimeController interface {
	controllers.BaseController
}

// NewServerTimeController ...
func NewServerTimeController(base api.Base) ServerTimeController {
	return &serverTimeController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type serverTimeController struct {
	api.Base
}

func (c *serverTimeController) GetRoute() string {
	return "/server-time"
}

func (c *serverTimeController) GetMethods() []string {
	return []string{"GET"}
}

// GetAccess ...
func (c *serverTimeController) GetAccess() controllers.Permission {
	return ""
}

// ServeHTTP ...
func (c *serverTimeController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.WR(w, 200, time.Now().Format("2006-01-02T15:04:05"))
}
