package freeinfo

import (
	"net/http"

	"github.com/mdiazp/rb/server/api"
	"github.com/mdiazp/rb/server/api/controllers"
	"github.com/mdiazp/rb/server/db/models"
)

// RetrieveDiskCategoriesController ...
type RetrieveDiskCategoriesController interface {
	controllers.BaseController
}

// NewRetrieveDiskCategoriesController ...
func NewRetrieveDiskCategoriesController(base api.Base) RetrieveDiskCategoriesController {
	return &retrieveDiskCategoriesController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type retrieveDiskCategoriesController struct {
	api.Base
}

func (c *retrieveDiskCategoriesController) GetRoute() string {
	return "/diskcategories"
}

func (c *retrieveDiskCategoriesController) GetMethods() []string {
	return []string{"GET"}
}

// GetAccess ...
func (c *retrieveDiskCategoriesController) GetAccess() controllers.Permission {
	return ""
}

// ServeHTTP ...
func (c *retrieveDiskCategoriesController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.WR(w, 200, models.GetDiskCategories())
}
