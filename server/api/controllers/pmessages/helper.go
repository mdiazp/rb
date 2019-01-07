package pmessages

import (
	"net/http"

	dbhandlers "github.com/mdiazp/rb/server/db/handlers"
	"github.com/mdiazp/rb/server/db/models"

	"github.com/mdiazp/rb/server/api"
)

// ReadPMessageBaseFilter ...
func ReadPMessageBaseFilter(c api.Base, w http.ResponseWriter,
	r *http.Request) *dbhandlers.PMessageBaseFilter {
	f := dbhandlers.PMessageBaseFilter{}

	f.MessengerID = c.GetQInt(w, r, "messengerID", false)

	auxStr := c.GetQString(w, r, "turnWeekDay", false)
	if auxStr != nil {
		tmp2 := (models.WeekDay)(*auxStr)
		f.TurnWeekDay = &tmp2
	}
	auxInt := c.GetQInt(w, r, "turnNum", false)
	if auxInt != nil {
		tmp2 := (models.TurnNum)(*auxInt)
		f.TurnNum = &tmp2
	}

	auxStr = c.GetQString(w, r, "type", false)
	if auxStr != nil {
		tmp2 := (models.MessageType)(*auxStr)
		f.Type = &tmp2
	}

	f.NotesSubstr = c.GetQString(w, r, "notesSubstr", false)

	return &f
}
