package core

import (
	"sort"
	"time"

	"github.com/mdiazp/rb/server/db/models"
)

// GetPDRXs ...
func GetPDRXs(pdrs []models.PDiskReservation) PDRXs {
	pdrxs := make([]pdrX, 0)

	for _, p := range pdrs {
		pdrxs = append(pdrxs, pdrX{
			PDiskReservation: p,
			ini:              true,
		})
		pdrxs = append(pdrxs, pdrX{
			PDiskReservation: p,
			ini:              false,
		})
	}

	sort.Sort(pdrXs(pdrxs))

	return pdrxs
}

// PDRX ...
type PDRX pdrX

// PDRXs ...
type PDRXs pdrXs

type pdrX struct {
	models.PDiskReservation
	ini bool
}

func (x *pdrX) End() bool {
	return !x.ini
}

func (x *pdrX) GetTime() time.Time {
	if x.ini {
		return x.InitialTime
	}
	return x.FinishTime.AddDate(0, 0, 1)
}

// pdrXs ...
type pdrXs []pdrX

// Len ...
func (pdrXs pdrXs) Len() int {
	return len(pdrXs)
}

// Swap ...
func (pdrXs pdrXs) Swap(i, j int) {
	pdrXs[i], pdrXs[j] = pdrXs[j], pdrXs[i]
}

// Less ...
func (pdrXs pdrXs) Less(i, j int) bool {
	if !pdrXs[i].GetTime().Equal(pdrXs[j].GetTime()) {
		return pdrXs[i].GetTime().Before(pdrXs[j].GetTime())
	}
	if pdrXs[i].TurnNum != pdrXs[j].TurnNum {
		return pdrXs[i].TurnNum < pdrXs[j].TurnNum
	}
	return pdrXs[i].ini
}
