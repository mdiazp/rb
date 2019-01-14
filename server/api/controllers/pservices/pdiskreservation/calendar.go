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

	pdrxs := core.GetPDRXs(*l)
	calendar := make([]core.PDRTurnCalendarState, 0)

	state := make([]*models.PDiskReservation, 0)
	stateTotals := make(map[string]int, 0)

	ln := len(pdrxs)

	nowDate := false
	for i := 0; i < ln; i++ {
		date := core.MoveDateToNextWeekDay(pdrxs[i].GetTime(), models.WeekDay(wd))
		if i == 0 && !nowDate {
			date = iniTime
			nowDate = true
		}
		for i < ln && !pdrxs[i].GetTime().After(date) {
			if !pdrxs[i].End() {
				state = append(state, &pdrxs[i].PDiskReservation)
				stateTotals[string(pdrxs[i].DiskCategoryRequest)]++
			} else {
				state = removePDR(state, pdrxs[i].PDiskReservation)
				stateTotals[string(pdrxs[i].DiskCategoryRequest)]--
			}
			i++
		}

		dcrr := make([]core.DiscCategoryRequestReport, 0)
		for key, value := range discTotalsByCategory {
			dcrr = append(
				dcrr,
				core.DiscCategoryRequestReport{
					Category: key,
					DCTotal:  value,
					DCRTotal: stateTotals[key],
				},
			)
		}

		calendar = append(
			calendar,
			core.PDRTurnCalendarState{
				Date:         date,
				PDRs:         state,
				DCRR:         dcrr,
				DCRNullTotal: stateTotals[string(models.DiskCategoryRequestNull)],
				DiscsTotal:   discsTotal,
			},
		)

		i--
	}

	if ln == 0 { //
		dcrr := make([]core.DiscCategoryRequestReport, 0)
		for key, value := range discTotalsByCategory {
			dcrr = append(
				dcrr,
				core.DiscCategoryRequestReport{
					Category: key,
					DCTotal:  value,
					DCRTotal: stateTotals[key],
				},
			)
		}

		calendar = append(
			calendar,
			core.PDRTurnCalendarState{
				Date:         iniTime,
				PDRs:         state,
				DCRR:         dcrr,
				DCRNullTotal: stateTotals[string(models.DiskCategoryRequestNull)],
			},
		)
	}

	c.WE(w, e, 500)
	c.WR(w, 200, calendar)
}

func removePDR(state []*models.PDiskReservation,
	pdr models.PDiskReservation) []*models.PDiskReservation {
	ln := len(state)
	newstate := make([]*models.PDiskReservation, 0)
	for i := 0; i < ln; i++ {
		if state[i].ID != pdr.ID {
			newstate = append(newstate, state[i])
		}
	}
	return newstate
}
