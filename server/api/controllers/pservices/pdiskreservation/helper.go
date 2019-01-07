package pdiskreservation

import (
	"net/http"

	"github.com/mdiazp/rb/server/api/controllers/pservices"
	dbhandlers "github.com/mdiazp/rb/server/db/handlers"
	"github.com/mdiazp/rb/server/db/models"

	"github.com/mdiazp/rb/server/api"
)

func readPDiskReservationFilter(c api.Base, w http.ResponseWriter,
	r *http.Request) *dbhandlers.PDiskReservationFilter {
	f := dbhandlers.PDiskReservationFilter{}

	f.PServiceBaseFilter = pservices.ReadPServiceBaseFilter(c, w, r)

	auxStr := c.GetQString(w, r, "diskCategoryRequest", false)
	if auxStr != nil {
		tmp2 := (models.DiskCategory)(*auxStr)
		f.DiskCategoryRequest = &tmp2
	}

	auxStr = c.GetQString(w, r, "turnWeekDay", false)
	if auxStr != nil {
		tmp2 := (models.WeekDay)(*auxStr)
		f.TurnWeekDay = &tmp2
	}

	auxInt := c.GetQInt(w, r, "turnNum", false)
	if auxInt != nil {
		tmp2 := (models.TurnNum)(*auxInt)
		f.TurnNum = &tmp2
	}

	return &f
}

func verificatePDiskReservationExistence(c api.Base, w http.ResponseWriter, id int) {
	_, e := c.DB().RetrievePDiskReservationByID(id)
	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
}
