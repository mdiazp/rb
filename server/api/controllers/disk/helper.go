package disk

import (
	"net/http"

	dbhandlers "github.com/mdiazp/rb/server/db/handlers"
	"github.com/mdiazp/rb/server/db/models"

	"github.com/mdiazp/rb/server/api"
)

func readDiskFilter(c api.Base, w http.ResponseWriter, r *http.Request) *dbhandlers.DiskFilter {
	f := dbhandlers.DiskFilter{}
	auxInt := c.GetQInt(w, r, "capacity", false)
	if auxInt != nil {
		tmp2 := (models.GB)(*auxInt)
		f.Capacity = &tmp2
	}

	auxStr := c.GetQString(w, r, "category", false)
	if auxStr != nil {
		tmp2 := (models.DiskCategory)(*auxStr)
		f.Category = &tmp2
	}

	f.Actived = c.GetQBool(w, r, "actived", false)

	return &f
}

func verificateDiskExistence(c api.Base, w http.ResponseWriter, id int) {
	_, e := c.DB().RetrieveDiskByID(id)
	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
}
