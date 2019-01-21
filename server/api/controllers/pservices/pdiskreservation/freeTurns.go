package pdiskreservation

import (
	"net/http"

	"github.com/mdiazp/rb/server/api"
	"github.com/mdiazp/rb/server/api/controllers"
	"github.com/mdiazp/rb/server/core"
	dbhandler "github.com/mdiazp/rb/server/db/handlers"
	"github.com/mdiazp/rb/server/db/models"
)

// FreeTurnsContoller ...
type FreeTurnsContoller interface {
	controllers.BaseController
}

// NewFreeTurnsContoller ...
func NewFreeTurnsContoller(base api.Base) FreeTurnsContoller {
	return &freeTurnsContoller{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type freeTurnsContoller struct {
	api.Base
}

func (c *freeTurnsContoller) GetRoute() string {
	return "/pdisk-reservation/free-turns"
}

func (c *freeTurnsContoller) GetMethods() []string {
	return []string{"GET"}
}

// GetAccess ...
func (c *freeTurnsContoller) GetAccess() controllers.Permission {
	return controllers.PermissionFreeTurnsPDiskReservation
}

// ServeHTTP ...
func (c *freeTurnsContoller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dcr := c.GetQString(w, r, "discCategoryRequest", true)
	initialTime := c.GetQTime(w, r, "initialTime", true)
	finishTime := c.GetQTime(w, r, "finishTime", false)
	if finishTime == nil {
		tmp := models.GetMaxTimeValue()
		finishTime = &tmp
	}

	tmpTrue := true
	filter := dbhandler.PDiskReservationFilter{
		PServiceBaseFilter: &dbhandler.PServiceBaseFilter{
			ActivedIntervalTime: &models.IntervalTime{
				InitialTime: *initialTime,
				FinishTime:  *finishTime,
			},
			ActivedClient: &tmpTrue,
		},
	}

	l, e := c.DB().RetrievePDiskReservationList(&filter, nil, nil)
	c.WE(w, e, 500)

	discTotalsByCategory, discsTotal, e := c.DB().GetTotalsByDiscCategory()
	c.WE(w, e, 500)

	calendar := core.GetPDRTurnCalendar(
		wd,
		l,
		iniTime,
		discTotalsByCategory,
		discsTotal,
	)

	c.WE(w, e, 500)
	c.WR(w, 200, *calendar)
}
