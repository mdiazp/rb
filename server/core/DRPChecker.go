package core

import (
	"sort"
	"time"

	"github.com/mdiazp/rb/server/db/models"
)

// CheckPDR ...
func CheckPDR(pdiskReservations []models.PDiskReservation,
	totals []TotDC) ([]PDRFreeTurn, []PDRConflict) {
	ps := make([]pdrX, 0)

	for _, p := range pdiskReservations {
		ps = append(ps, pdrX{
			PDiskReservation: p,
			ini:              true,
		})
		ps = append(ps, pdrX{
			PDiskReservation: p,
			ini:              false,
		})
	}

	sort.Sort(pdrXs(ps))

	nps := len(ps)

	wds := models.GetWeekDays()
	nwd := len(wds)
	wdToInt := make(map[models.WeekDay]int)
	for i, wd := range wds {
		wdToInt[wd] = i
	}

	tns := models.GetTurnNums()
	ntn := len(tns)

	ndc := len(totals)
	dcToInt := make(map[models.DiskCategory]int)
	T := 0
	for i, total := range totals {
		dcToInt[total.DiskCategory] = i
		T += total.Total
	}
	dcToInt[models.DiskCategoryRequestNull] = ndc
	ndc++

	cnt := *init3d(nwd, ntn, ndc)
	mx := *init3d(nwd, ntn, ndc)

	conflicts := make([]PDRConflict, 0)

	for ind, px := range ps {
		i := wdToInt[px.TurnWeekDay]
		j := (int)(px.TurnNum)
		k := dcToInt[px.DiskCategoryRequest]

		if !px.ini {
			cnt[i][j][k]--
			if k != ndc-1 {
				cnt[i][j][ndc-1]--
			}
		} else {
			cnt[i][j][k]++
			mx[i][j][k] = max(mx[i][j][k], cnt[i][j][k])
			if k != ndc-1 {
				cnt[i][j][ndc-1]++
				mx[i][j][ndc-1] = max(mx[i][j][ndc-1], cnt[i][j][ndc-1])
			}

			if ind+1 == nps || px.GetTime().Before(ps[ind+1].GetTime()) ||
				px.TurnNum < ps[ind+1].TurnNum || !ps[ind+1].ini {
				ok := (cnt[i][j][ndc-1] <= T)
				if !ok {
					for k2, tot := range totals {
						if cnt[i][j][k2] > tot.Total {
							ok = false
							break
						}
					}
				}
				if !ok {
					conflicts = append(conflicts, PDRConflict{
						Date: px.InitialTime,
						Turn: px.Turn,
					})
				}
			}
		}
	}

	free := make([]PDRFreeTurn, 0)
	for i, wd := range wds {
		for j, tn := range tns {
			if mx[i][j][ndc] < T {
				free = append(free, PDRFreeTurn{
					Turn: models.Turn{
						TurnWeekDay: wd,
						TurnNum:     tn,
					},
					DCRequest: models.DiskCategoryRequestNull,
				})

				for k, tot := range totals {
					if mx[i][j][k] < tot.Total {
						free = append(free, PDRFreeTurn{
							Turn: models.Turn{
								TurnWeekDay: wd,
								TurnNum:     tn,
							},
							DCRequest: tot.DiskCategory,
						})
					}
				}
			}
		}
	}

	return free, conflicts
}

// PDRFreeTurn ...
type PDRFreeTurn struct {
	models.Turn
	DCRequest models.DiskCategory
}

// PDRConflict ...
type PDRConflict struct {
	Date time.Time
	models.Turn
}

// TotDC ...
type TotDC struct {
	DiskCategory models.DiskCategory
	Total        int
}

func max(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func init3d(n, m, k int) *[][][]int {
	x := make([][][]int, n)
	for i := 0; i < n; i++ {
		x[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			x[i][j] = make([]int, k)
		}
	}
	return &x
}
