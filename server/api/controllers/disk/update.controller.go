package disk

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
	return "/disk/{diskID}"
}

func (c *updateController) GetMethods() []string {
	return []string{"PATCH"}
}

// GetAccess ...
func (c *updateController) GetAccess() controllers.Permission {
	return controllers.PermissionUpdateDisk
}

// ServeHTTP ...
func (c *updateController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := c.GetPInt(w, r, "diskID")
	var o models.Disk
	c.ReadJSON(w, r, &o)
	o.ID = id

	verificateDiskExistence(c, w, id)

	c.WE400(w, o.Valid())

	e := c.DB().UpdateDisk(&o)
	if e != nil && strings.Contains(e.Error(),
		fmt.Sprintf(`"%s"`, models.SQLConstrainstDiskNameUI)) {
		c.WE400(w,
			c.MakeValidationError("Name", "disk with same name already exists"),
		)
	}
	if e != nil && strings.Contains(e.Error(),
		fmt.Sprintf(`"%s"`, models.SQLConstrainstDiskSerialNumberUI)) {
		c.WE400(w, c.MakeValidationError(
			"SerialNumber", "disk with same serial number already exists"),
		)
	}

	c.WE(w, e, 500)
	c.WR(w, 200, o)
}
