package messenger

import (
	"fmt"
	"net/http"
	"strings"

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
	return "/messenger/{messengerID}"
}

func (c *updateController) GetMethods() []string {
	return []string{"PATCH"}
}

// GetAccess ...
func (c *updateController) GetAccess() controllers.Permission {
	return controllers.PermissionUpdateMessenger
}

// ServeHTTP ...
func (c *updateController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := c.GetPInt(w, r, "messengerID")
	var o models.Messenger
	c.ReadJSON(w, r, &o)
	o.ID = id

	verificateMessengerExistence(c, w, id)

	c.WE400(w, o.Valid())

	e := c.DB().UpdateMessenger(&o)
	if e != nil && strings.Contains(e.Error(),
		fmt.Sprintf(`"%s"`, models.SQLConstrainstMessengerNameUI)) {
		c.WE400(w,
			c.MakeValidationError("Name", "Messenger with same name already exists"),
		)
	}

	c.WE(w, e, 500)
	c.WR(w, 200, o)
}
