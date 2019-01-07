package pservices

import (
	"net/http"

	dbhandlers "github.com/mdiazp/rb/server/db/handlers"
	"github.com/mdiazp/rb/server/db/models"

	"github.com/mdiazp/rb/server/api"
)

// ReadPServiceBaseFilter ...
func ReadPServiceBaseFilter(c api.Base, w http.ResponseWriter,
	r *http.Request) *dbhandlers.PServiceBaseFilter {
	f := dbhandlers.PServiceBaseFilter{}

	f.ClientID = c.GetQInt(w, r, "clientID", false)
	f.ActivedClient = c.GetQBool(w, r, "activedClient", false)
	it := c.GetQTime(w, r, "activedInitialTime", false)
	ft := c.GetQTime(w, r, "activedFinishTime", false)
	if it != nil || ft != nil {
		if it == nil {
			tmp := models.GetMinTimeValue()
			it = &tmp
		}
		if ft == nil {
			tmp := models.GetMaxTimeValue()
			ft = &tmp
		}
		f.ActivedIntervalTime = &models.IntervalTime{
			InitialTime: *it,
			FinishTime:  *ft,
		}
	}

	return &f
}
