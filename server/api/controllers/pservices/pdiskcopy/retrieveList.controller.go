package pdiskcopy

import (
	"net/http"

	"github.com/mdiazp/rb/server/api"
	"github.com/mdiazp/rb/server/api/controllers"
	dbhandlers "github.com/mdiazp/rb/server/db/handlers"
)

// RetrieveListController ...
type RetrieveListController interface {
	controllers.BaseController
}

// NewRetrieveListController ...
func NewRetrieveListController(base api.Base) RetrieveListController {
	return &retrieveListController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type retrieveListController struct {
	api.Base
}

func (c *retrieveListController) GetRoute() string {
	return "/pdiskcopies"
}

func (c *retrieveListController) GetMethods() []string {
	return []string{"GET"}
}

// GetAccess ...
func (c *retrieveListController) GetAccess() controllers.Permission {
	return controllers.PermissionRetrievePDiskCopy
}

// ServeHTTP ...
func (c *retrieveListController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f := readPDiskCopyFilter(c.Base, w, r)
	ob := c.GetQOrderBy(w, r)
	p := c.GetQPaginator(w, r)

	l, e := c.DB().RetrievePDiskCopyList(f, ob, p)

	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
	c.WR(w, 200, *l)
}
