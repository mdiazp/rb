package handlers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mdiazp/rb/server/db/models"
)

// PDiskReservationHandler ...
type PDiskReservationHandler interface {
	CreatePDiskReservation(o *models.PDiskReservation) error
	RetrievePDiskReservationByID(id int) (*models.PDiskReservation, error)
	UpdatePDiskReservation(o *models.PDiskReservation) error
	DeletePDiskReservation(id int) error

	RetrievePDiskReservationList(filter *PDiskReservationFilter,
		orderBy *OrderBy, pag *Paginator) (*[]models.PDiskReservation, error)
	CountPDiskReservations(filter *PDiskReservationFilter) (count int, e error)
}

// PDiskReservationFilter ...
type PDiskReservationFilter struct {
	PServiceBaseFilter  *PServiceBaseFilter
	DiskCategoryRequest *models.DiskCategory
	TurnWeekDay         *models.WeekDay
	TurnNum             *models.TurnNum
}

/////////////////////////////////////////////////////////////////////////////////////

func makePDiskReservationFilter(db *gorm.DB, filter *PDiskReservationFilter) *gorm.DB {
	if filter == nil {
		return db
	}
	db = makePServiceBaseFilter(db, filter.PServiceBaseFilter, "p_disk_reservation")

	if filter.DiskCategoryRequest != nil {
		db = db.Where("p_disk_reservation.disk_category_request = ?",
			*(filter.DiskCategoryRequest))
	}
	if filter.TurnWeekDay != nil {
		db = db.Where("p_disk_reservation.turn_week_day = ?",
			*(filter.TurnWeekDay))
	}
	if filter.TurnNum != nil {
		db = db.Where("p_disk_reservation.turn_num = ?",
			*(filter.TurnNum))
	}

	return db
}

func (h *handler) CreatePDiskReservation(o *models.PDiskReservation) error {
	return h.Create(o).Error
}

func (h *handler) RetrievePDiskReservationByID(id int) (*models.PDiskReservation, error) {
	o := &models.PDiskReservation{}
	e := h.Where("p_disk_reservation.id = ?", id).First(o).Error
	return o, e
}

func (h *handler) UpdatePDiskReservation(o *models.PDiskReservation) error {
	return h.Save(o).Error
}

func (h *handler) DeletePDiskReservation(id int) error {
	return h.Delete(models.PDiskReservation{
		PServiceBase: models.PServiceBase{
			ID: id,
		},
	}).Error
}

func (h *handler) RetrievePDiskReservationList(filter *PDiskReservationFilter,
	orderBy *OrderBy, pag *Paginator) (*[]models.PDiskReservation, error) {
	db := h.DB.Model(&models.PDiskReservation{})
	db = addJoinToClient(db, "p_disk_reservation")
	db = makePDiskReservationFilter(db, filter)
	db = orderByAndPaginator(db, orderBy, pag, (models.PDiskReservation{}).TableName())

	l := make([]models.PDiskReservation, 0)

	db = db.
		Select(
			"p_disk_reservation.id, " +
				"p_disk_reservation.client_id, " +
				"client.name, " +
				"p_disk_reservation.initial_time, " +
				"p_disk_reservation.finish_time, " +
				"p_disk_reservation.cost, " +
				"p_disk_reservation.turn_week_day, " +
				"p_disk_reservation.turn_num, " +
				"p_disk_reservation.disk_category_request",
		)

	fmt.Println("	SQL Query = ", db.QueryExpr())

	rows, e := db.Rows()
	if e != nil {
		return nil, e
	}
	defer rows.Close()

	for rows.Next() {
		o := models.PDiskReservation{}
		e = rows.Scan(
			&o.ID,
			&o.ClientID,
			&o.ClientName,
			&o.InitialTime,
			&o.FinishTime,
			&o.Cost,
			&o.TurnWeekDay,
			&o.TurnNum,
			&o.DiskCategoryRequest,
		)
		if e != nil {
			return nil, e
		}
		l = append(l, o)
	}

	return &l, e
}

func (h *handler) CountPDiskReservations(filter *PDiskReservationFilter) (count int, e error) {
	db := addJoinToClient(
		h.DB.Model(&models.PDiskReservation{}),
		"p_disk_reservation",
	)
	e = makePDiskReservationFilter(db, filter).Count(&count).Error
	return
}
