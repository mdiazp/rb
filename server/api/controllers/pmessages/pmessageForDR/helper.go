package pmessagefordr

import (
	"net/http"

	"github.com/mdiazp/rb/server/api/controllers/pmessages"
	dbhandlers "github.com/mdiazp/rb/server/db/handlers"

	"github.com/mdiazp/rb/server/api"
)

func readPMessageForDRFilter(c api.Base, w http.ResponseWriter,
	r *http.Request) *dbhandlers.PMessageForDRFilter {
	f := dbhandlers.PMessageForDRFilter{}

	f.PMessageBaseFilter = pmessages.ReadPMessageBaseFilter(c, w, r)

	f.PDiskReservationID = c.GetQInt(w, e, "pDiskReservationID")

	return &f
}

func verificatePMessageForDRExistence(c api.Base, w http.ResponseWriter, id int) {
	_, e := c.DB().RetrievePMessageForDRByID(id)
	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
}
