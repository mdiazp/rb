package core

import (
	"time"

	"github.com/mdiazp/rb/server/db/models"
)

// GetPDRTurnCalendar ...
func GetPDRTurnCalendar(wd models.WeekDay, l *[]models.PDiskReservation,
	iniTime time.Time, discTotalsByCategory map[string]int,
	discsTotal int) *PDRTurnCalendar {

	pdrxs := GetPDRXs(*l)
	states := make([]PDRTurnCalendarState, 0)
	totalWrongStates := 0

	state := make([]*models.PDiskReservation, 0)
	stateTotals := make(map[string]int, 0)

	ln := len(pdrxs)

	if ln == 0 {
		calendarState := makePDRTurnCalendarState(
			iniTime,
			state,
			stateTotals,
			discTotalsByCategory,
			discsTotal,
		)
		if calendarState.Wrong {
			totalWrongStates++
		}
		states = append(
			states,
			*calendarState,
		)
	} else {
		nowDate := false
		for i := 0; i < ln; i++ {
			date := MoveDateToNextWeekDay(pdrxs[i].GetTime(), models.WeekDay(wd))
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

			calendarState := makePDRTurnCalendarState(
				date,
				state,
				stateTotals,
				discTotalsByCategory,
				discsTotal,
			)
			if calendarState.Wrong {
				totalWrongStates++
			}
			states = append(
				states,
				*calendarState,
			)

			i--
		}
	}

	return &PDRTurnCalendar{
		TotalWrongStates: totalWrongStates,
		States:           states,
	}
}

func makePDRTurnCalendarState(iniTime time.Time,
	state []*models.PDiskReservation, stateTotals map[string]int,
	discTotalsByCategory map[string]int, discsTotal int) *PDRTurnCalendarState {

	good := true

	dcrr := make([]DiscCategoryRequestReport, 0)

	// iterate over dcs to avoid random iterations over map
	dcs := models.GetDiskCategories()
	for _, dc := range dcs {
		key, value := string(dc), discTotalsByCategory[string(dc)]
		dcrr = append(
			dcrr,
			DiscCategoryRequestReport{
				Category: key,
				DCTotal:  value,
				DCRTotal: stateTotals[key],
			},
		)
		good = (good && stateTotals[key] <= value)
	}

	good = (good && len(state) <= discsTotal)

	return &PDRTurnCalendarState{
		Date:         iniTime,
		PDRs:         state,
		DCRR:         dcrr,
		DCRNullTotal: stateTotals[string(models.DiskCategoryRequestNull)],
		DiscsTotal:   discsTotal,
		Wrong:        !good,
	}
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

// PDRTurnCalendar ...
type PDRTurnCalendar struct {
	TotalWrongStates int
	States           []PDRTurnCalendarState
}

// PDRTurnCalendarState ...
type PDRTurnCalendarState struct {
	Date         time.Time
	PDRs         []*models.PDiskReservation
	DCRR         []DiscCategoryRequestReport
	DCRNullTotal int
	DiscsTotal   int
	Wrong        bool
}

// DiscCategoryRequestReport ...
type DiscCategoryRequestReport struct {
	Category string
	DCTotal  int
	DCRTotal int
}
