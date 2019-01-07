package client

import (
	"net/http"

	dbhandlers "github.com/mdiazp/rb/server/db/handlers"

	"github.com/mdiazp/rb/server/api"
)

func readClientFilter(c api.Base, w http.ResponseWriter,
	r *http.Request) *dbhandlers.ClientFilter {
	f := dbhandlers.ClientFilter{}

	f.IdentificationPrefix = c.GetQString(w, r, "identificationPrefix", false)
	f.NameSubstr = c.GetQString(w, r, "nameSubstr", false)
	f.AddressSubstr = c.GetQString(w, r, "addressSubstr", false)
	f.PhonesSubstr = c.GetQString(w, r, "phonesSubstr", false)
	f.DescriptionSubstr = c.GetQString(w, r, "descriptionSubstr", false)
	f.Actived = c.GetQBool(w, r, "actived", false)

	return &f
}

func verificateClientExistence(c api.Base, w http.ResponseWriter, id int) {
	_, e := c.DB().RetrieveClientByID(id)
	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
}
