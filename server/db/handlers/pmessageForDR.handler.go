package handlers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mdiazp/rb/server/db/models"
)

// PMessageForDRHandler ...
type PMessageForDRHandler interface {
	CreatePMessageForDR(o *models.PMessageForDR) error
	RetrievePMessageForDRByID(id int) (*models.PMessageForDR, error)
	UpdatePMessageForDR(o *models.PMessageForDR) error
	DeletePMessageForDR(id int) error

	RetrievePMessageForDRList(filter *PMessageForDRFilter,
		orderBy *OrderBy, pag *Paginator) (*[]models.PMessageForDR, error)
	CountPMessageForDRs(filter *PMessageForDRFilter) (count int, e error)
}

// PMessageForDRFilter ...
type PMessageForDRFilter struct {
	PMessageBaseFilter *PMessageBaseFilter
	PDiskReservationID *int
}

/////////////////////////////////////////////////////////////////////////////////////

func makePMessageForDRFilter(db *gorm.DB, filter *PMessageForDRFilter) *gorm.DB {
	if filter == nil {
		return db
	}
	db = makePMessageBaseFilter(db, filter.PMessageBaseFilter, "p_message_for_dr")
	db = makeEqIntFilter(db, filter.PDiskReservationID, "p_disk_reservation_id")

	return db
}

func (h *handler) CreatePMessageForDR(o *models.PMessageForDR) error {
	return h.Create(o).Error
}

func (h *handler) RetrievePMessageForDRByID(id int) (*models.PMessageForDR, error) {
	o := &models.PMessageForDR{}
	e := h.Where("p_message_for_dr.id = ?", id).First(o).Error
	return o, e
}

func (h *handler) UpdatePMessageForDR(o *models.PMessageForDR) error {
	return h.Save(o).Error
}

func (h *handler) DeletePMessageForDR(id int) error {
	return h.Delete(models.PMessageForDR{
		PMessageBase: models.PMessageBase{
			ID: id,
		},
	}).Error
}

func (h *handler) RetrievePMessageForDRList(filter *PMessageForDRFilter,
	orderBy *OrderBy, pag *Paginator) (*[]models.PMessageForDR, error) {
	db := makePMessageForDRFilter(h.DB, filter)
	db = orderByAndPaginator(db, orderBy, pag, (models.PMessageForDR{}).TableName())

	fmt.Println("-------------> ", db.QueryExpr())

	l := make([]models.PMessageForDR, 0)
	e := db.Find(&l).Error

	return &l, e
}

func (h *handler) CountPMessageForDRs(filter *PMessageForDRFilter) (count int, e error) {
	e = makePMessageForDRFilter(h.DB.Model(&models.PMessageForDR{}), filter).Count(&count).Error
	return
}
