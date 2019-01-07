package messenger

import (
	"net/http"

	dbhandlers "github.com/mdiazp/rb/server/db/handlers"

	"github.com/mdiazp/rb/server/api"
)

func readMessengerFilter(c api.Base, w http.ResponseWriter, r *http.Request) *dbhandlers.MessengerFilter {
	f := dbhandlers.MessengerFilter{}
	f.NameSubstr = c.GetQString(w, r, "nameSubstr", false)
	return &f
}

func verificateMessengerExistence(c api.Base, w http.ResponseWriter, id int) {
	_, e := c.DB().RetrieveMessengerByID(id)
	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
}
