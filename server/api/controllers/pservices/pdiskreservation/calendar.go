package pdiskreservation

import (
	"net/http"
	"time"

	"github.com/mdiazp/rb/server/api"
	"github.com/mdiazp/rb/server/api/controllers"
	"github.com/mdiazp/rb/server/core"
	dbhandler "github.com/mdiazp/rb/server/db/handlers"
	"github.com/mdiazp/rb/server/db/models"
)

// CalendarController ...
type CalendarController interface {
	controllers.BaseController
}

// NewCalendarController ...
func NewCalendarController(base api.Base) CalendarController {
	return &calendarController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type calendarController struct {
	api.Base
}

func (c *calendarController) GetRoute() string {
	return "/pdisk-reservation/calendar/{weekDay}/{turnNum}"
}

func (c *calendarController) GetMethods() []string {
	return []string{"GET"}
}

// GetAccess ...
func (c *calendarController) GetAccess() controllers.Permission {
	return controllers.PermissionRetrievePDiskReservation
}

// ServeHTTP ...
func (c *calendarController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wd := (models.WeekDay)(c.GetPString(w, r, "weekDay"))
	tn := (models.TurnNum)(c.GetPInt(w, r, "turnNum"))

	iniTime := core.MoveDateToNextWeekDay(time.Now(), models.WeekDay(wd))
	endTime := models.GetMaxTimeValue()

	tmpTrue := true
	filter := dbhandler.PDiskReservationFilter{
		PServiceBaseFilter: &dbhandler.PServiceBaseFilter{
			ActivedIntervalTime: &models.IntervalTime{
				InitialTime: iniTime,
				FinishTime:  endTime,
			},
			ActivedClient: &tmpTrue,
		},
		TurnWeekDay: &wd,
		TurnNum:     &tn,
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
