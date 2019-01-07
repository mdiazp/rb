package messenger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mdiazp/rb/server/api"
	"github.com/mdiazp/rb/server/api/controllers"
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
	return "/messenger"
}

func (c *createController) GetMethods() []string {
	return []string{"POST"}
}

// GetAccess ...
func (c *createController) GetAccess() controllers.Permission {
	return controllers.PermissionCreateMessenger
}

// ServeHTTP ...
func (c *createController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var o models.Messenger
	c.ReadJSON(w, r, &o)
	o.ID = 0

	c.WE400(w, o.Valid())

	e := c.DB().CreateMessenger(&o)
	if e != nil && strings.Contains(e.Error(),
		fmt.Sprintf(`"%s"`, models.SQLConstrainstMessengerNameUI)) {
		c.WE400(w,
			c.MakeValidationError("Name", "Messenger with same name already exists"),
		)
	}

	c.WE(w, e, 500)
	c.WR(w, 200, o)
}
