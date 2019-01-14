package main

import (
	"fmt"
	"time"

	"github.com/mdiazp/rb/server/core"
	"github.com/mdiazp/rb/server/db/models"
	"golang.org/x/exp/rand"
)

func getClients() *[]models.Client {
	res, e := db.RetrieveClientList(nil, nil, nil)

	if e != nil {
		panic(e)
	}

	return res
}

var (
	dcs  = models.GetDiskCategories()
	dcrn = models.DiskCategoryRequestNull
	dcrs = append(dcs, dcrn)
	wds  = models.GetWeekDays()
	tns  = models.GetTurnNums()
)

func now() time.Time {
	now := time.Now()
	r, e := time.Parse(
		"2006-01-02",
		fmt.Sprintf(
			"%d-%02d-%02d",
			now.Year(), now.Month(), now.Day(),
		),
	)
	if e != nil {
		panic(e)
	}
	return r
}

func addSeveralPDRs() {
	clients := *getClients()
	costs := []int{0, 25, 50}

	n := 1000

	for i := 0; i < n; i++ {
		iniTime := now()

		iniTime = iniTime.AddDate(
			rand.Int()%3,
			rand.Int()%12+1,
			rand.Int()%28+1,
		)

		endTime := iniTime.AddDate(
			rand.Int()%3,
			rand.Int()%12+1,
			rand.Int()%28+1,
		)

		pdr := models.PDiskReservation{
			PServiceBase: models.PServiceBase{
				ClientID: clients[rand.Int()%len(clients)].ID,
				IntervalTime: models.IntervalTime{
					InitialTime: iniTime,
					FinishTime:  endTime,
				},
				Cost: costs[rand.Int()%len(costs)],
			},
			Turn: models.Turn{
				TurnWeekDay: wds[rand.Int()%len(wds)],
				TurnNum:     tns[rand.Int()%len(tns)],
			},
			DiskCategoryRequest: dcrs[rand.Int()%len(dcrs)],
		}

		pdr.InitialTime = core.MoveDateToNextWeekDay(pdr.InitialTime, pdr.TurnWeekDay)
		pdr.FinishTime = core.MoveDateToPreviousWeekDay(pdr.FinishTime, pdr.TurnWeekDay)

		es := pdr.Valid()
		if es != nil && len(*es) > 0 {
			fmt.Println("Validation Errors:")
			for _, e := range *es {
				fmt.Println(e.PropertyName, ":", e.Error)
			}
			panic(nil)
		}

		db.CreatePDiskReservation(&pdr)

		/*fmt.Printf(
			"%d %s %s\n",
			pdr.ID,
			pdr.InitialTime.Format("2006-01-02"),
			pdr.FinishTime.Format("2006-01-02"),
		)*/
	}
}

func deleteAllPDRs() {
	pdrs, e := db.RetrievePDiskReservationList(nil, nil, nil)
	if e != nil {
		panic(e)
	}
	for _, pdr := range *pdrs {
		db.DeletePDiskReservation(pdr.ID)
	}
}
