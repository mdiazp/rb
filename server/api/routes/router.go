package routes

import (
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"

	"github.com/mdiazp/rb/server/api"
	"github.com/mdiazp/rb/server/api/controllers"
	"github.com/mdiazp/rb/server/api/controllers/client"
	"github.com/mdiazp/rb/server/api/controllers/disk"
	"github.com/mdiazp/rb/server/api/controllers/freeinfo"
	"github.com/mdiazp/rb/server/api/controllers/messenger"
	"github.com/mdiazp/rb/server/api/controllers/pservices/pdiskcopy"
	"github.com/mdiazp/rb/server/api/controllers/pservices/pdiskreservation"
	"github.com/mdiazp/rb/server/api/controllers/session"
	"github.com/mdiazp/rb/server/api/middlewares"
)

// Router ...
func Router(base api.Base) http.Handler {
	// Middlewares
	logger := middlewares.Logger(base)

	ctrs := []controllers.BaseController{
		session.NewLoginController(base),
		session.NewLogoutController(base),
		session.NewProvidersController(base),

		disk.NewCreateController(base),
		disk.NewRetrieveController(base),
		disk.NewUpdateController(base),
		disk.NewDeleteController(base),
		disk.NewRetrieveListController(base),
		disk.NewCountController(base),

		client.NewCreateController(base),
		client.NewRetrieveController(base),
		client.NewUpdateController(base),
		client.NewDeleteController(base),
		client.NewRetrieveListController(base),
		client.NewCountController(base),

		messenger.NewCreateController(base),
		messenger.NewRetrieveController(base),
		messenger.NewUpdateController(base),
		messenger.NewDeleteController(base),
		messenger.NewRetrieveListController(base),
		messenger.NewCountController(base),

		pdiskreservation.NewCreateController(base),
		pdiskreservation.NewRetrieveController(base),
		pdiskreservation.NewUpdateController(base),
		pdiskreservation.NewDeleteController(base),
		pdiskreservation.NewRetrieveListController(base),
		pdiskreservation.NewCountController(base),

		pdiskcopy.NewCreateController(base),
		pdiskcopy.NewRetrieveController(base),
		pdiskcopy.NewUpdateController(base),
		pdiskcopy.NewDeleteController(base),
		pdiskcopy.NewRetrieveListController(base),
		pdiskcopy.NewCountController(base),

		freeinfo.NewRetrieveDiskCategoriesController(base),
		freeinfo.NewRetrieveTurnNumsController(base),
		freeinfo.NewServerTimeController(base),
	}

	router := mux.NewRouter()

	router.
		PathPrefix("/swagger/").
		Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir(base.PublicFolderPath()))))

	router.Use(logger)

	for _, ctr := range ctrs {
		var h http.Handler = ctr
		if ctr.GetAccess() != "" {
			// h = middlewares.CheckAccessControl(base, ctr)
			// h = middlewares.MustAuth(base, h)
		}
		router.Handle(ctr.GetRoute(), h).Methods(ctr.GetMethods()...)
	}

	h := cors.AllowAll().Handler(router)

	return h
}
