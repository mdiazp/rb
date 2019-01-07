package pdiskcopy

import (
	"net/http"

	"github.com/mdiazp/rb/server/api/controllers/pservices"
	dbhandlers "github.com/mdiazp/rb/server/db/handlers"

	"github.com/mdiazp/rb/server/api"
)

func readPDiskCopyFilter(c api.Base, w http.ResponseWriter,
	r *http.Request) *dbhandlers.PDiskCopyFilter {
	f := dbhandlers.PDiskCopyFilter{}

	f.PServiceBaseFilter = pservices.ReadPServiceBaseFilter(c, w, r)

	return &f
}

func verificatePDiskCopyExistence(c api.Base, w http.ResponseWriter, id int) {
	_, e := c.DB().RetrievePDiskCopyByID(id)
	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
}
