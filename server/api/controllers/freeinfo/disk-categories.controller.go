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
	return "/disc-categories-info"
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
	x := DiscCategories{
		DiscCategories:          models.GetDiskCategories(),
		DiscCategoryRequestNull: models.DiskCategoryRequestNull,
	}
	c.WR(w, 200, x)

}

// DiscCategories ...
type DiscCategories struct {
	DiscCategories          []models.DiskCategory
	DiscCategoryRequestNull models.DiskCategory
}
