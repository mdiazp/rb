package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/mdiazp/rb/server/authprovider"
	dbH "github.com/mdiazp/rb/server/db/handlers"
	"github.com/mdiazp/rb/server/db/models"
)

// Base ...
type Base interface {
	DB() dbH.Handler
	Logger() *log.Logger

	JWTHandler() JWTHandler
	PublicFolderPath() string
	GetEnv() string

	GetAuthProvider(provider AuthProvider) authprovider.Provider
	GetAuthProviderNames() []AuthProvider

	ReadJSON(w http.ResponseWriter, r *http.Request, objs ...interface{})

	GetPInt(w http.ResponseWriter, r *http.Request, vname string) int
	GetPString(w http.ResponseWriter, r *http.Request, vname string) string

	GetQInt(w http.ResponseWriter, r *http.Request, vname string, required ...bool) *int
	GetQBool(w http.ResponseWriter, r *http.Request, vname string, required ...bool) *bool
	GetQString(w http.ResponseWriter, r *http.Request, vname string, required ...bool) *string
	GetQTime(w http.ResponseWriter, r *http.Request, vname string, required ...bool) *time.Time
	GetQPaginator(w http.ResponseWriter, r *http.Request) *dbH.Paginator
	GetQOrderBy(w http.ResponseWriter, r *http.Request) *dbH.OrderBy

	ContextWriteAuthor(r *http.Request, author *models.User)
	ContextReadAuthor(w http.ResponseWriter, r *http.Request, required ...bool) *models.User

	MakeValidationError(propertyName string, e string,
		es ...*[]models.ValidationError) *[]models.ValidationError
	WE400(w http.ResponseWriter, es *[]models.ValidationError)
	WE(w http.ResponseWriter, e error, status int)
	WR(w http.ResponseWriter, status int, body interface{})
}

// NewBase ...
func NewBase(db dbH.Handler, logFile *os.File, jwth JWTHandler,
	publicFolderPath string, env string) Base {
	return &base{
		db:               db,
		logger:           NewLogger(logFile),
		publicFolderPath: publicFolderPath,
		jwth:             jwth,
		env:              env,
	}
}

///////////////////////////////////////////////////////////////////////////

type base struct {
	db               dbH.Handler
	logger           *log.Logger
	jwth             JWTHandler
	publicFolderPath string
	env              string
}

func (b *base) DB() dbH.Handler {
	return b.db
}

func (b *base) Logger() *log.Logger {
	return b.logger
}

func (b *base) JWTHandler() JWTHandler {
	return b.jwth
}

func (b *base) PublicFolderPath() string {
	return b.publicFolderPath
}

func (b *base) GetEnv() string {
	return b.env
}

func (b *base) ReadJSON(w http.ResponseWriter, r *http.Request, objs ...interface{}) {
	decoder := json.NewDecoder(r.Body)
	for _, obj := range objs {
		e := decoder.Decode(obj)
		if e != nil {
			b.WE400(w, b.MakeValidationError("RequestBody", e.Error()))
		}
	}
}

func (b *base) GetPInt(w http.ResponseWriter, r *http.Request, vname string) int {
	vs := b.GetPString(w, r, vname)
	v, e := strconv.Atoi(vs)
	if e != nil {
		b.WE400(w, b.MakeValidationError(vname, e.Error()))
	}
	return v
}

func (b *base) GetPString(w http.ResponseWriter, r *http.Request, vname string) string {
	vars := mux.Vars(r)
	v, ok := vars[vname]
	if !ok {
		b.WE400(w, b.MakeValidationError(vname, "Required field is missing"))
	}
	return v
}

func (b *base) GetQInt(w http.ResponseWriter, r *http.Request, vname string, required ...bool) *int {
	vs := b.GetQString(w, r, vname, required...)
	if vs == nil {
		return nil
	}

	v, e := strconv.Atoi(*vs)
	if e != nil {
		b.WE400(w, b.MakeValidationError(vname, "Value must be an integer"))
	}
	return &v
}

func (b *base) GetQBool(w http.ResponseWriter, r *http.Request,
	vname string, required ...bool) *bool {
	vs := b.GetQString(w, r, vname, required...)
	if vs == nil {
		return nil
	}

	v, e := strconv.ParseBool(*vs)
	if e != nil {
		b.WE400(w, b.MakeValidationError(vname, "Value must be a bool"))
	}
	return &v
}

func (b *base) GetQTime(w http.ResponseWriter, r *http.Request,
	vname string, required ...bool) *time.Time {
	vs := b.GetQString(w, r, vname, required...)
	if vs == nil {
		return nil
	}

	v, e := time.Parse("2006-01-02", *vs)
	if e != nil {
		b.WE400(w, b.MakeValidationError(vname, "Value Format must be 2006-01-02"))
	}
	return &v
}

func (b *base) GetQString(w http.ResponseWriter, r *http.Request, vname string, required ...bool) *string {
	v := r.URL.Query().Get(vname)
	req := false
	if len(required) > 0 {
		req = required[0]
	}
	if v == "" && req {
		b.WE400(w, b.MakeValidationError(vname, "Required field is missing"))
	}
	if v == "" {
		return nil
	}
	return &v
}

func (b *base) GetQPaginator(w http.ResponseWriter, r *http.Request) *dbH.Paginator {
	p := dbH.Paginator{}
	limit := b.GetQInt(w, r, "limit", false)
	offset := b.GetQInt(w, r, "offset", false)
	if limit == nil || offset == nil {
		return nil
	}
	p.Limit = *limit
	p.Offset = *offset
	return &p
}

func (b *base) GetQOrderBy(w http.ResponseWriter, r *http.Request) *dbH.OrderBy {
	ob := dbH.OrderBy{}
	by := b.GetQString(w, r, "orderby", false)
	desc := b.GetQBool(w, r, "desc", false)
	if by == nil {
		return nil
	}
	if desc == nil {
		tmp := false
		desc = &tmp
	}
	ob.By = *by
	ob.DESC = *desc
	return &ob
}

func (b *base) MakeValidationError(propertyName string, e string,
	eS ...*[]models.ValidationError) *[]models.ValidationError {
	var es *[]models.ValidationError
	if len(eS) > 0 {
		es = eS[0]
	} else {
		es = &[]models.ValidationError{}
	}
	*es = append(*es, models.ValidationError{
		PropertyName: propertyName,
		Error:        e,
	})
	return es
}

func (b *base) WE400(w http.ResponseWriter, es *[]models.ValidationError) {
	if es == nil || len(*es) == 0 {
		return
	}
	b.whs(w, 400)
	body, _ := json.Marshal(*es)
	w.Write(body)

	panic(
		Error{
			Status:   400,
			Location: WAI(2),
			error:    fmt.Errorf("Bad Request"),
		},
	)
}

func (b *base) WE(w http.ResponseWriter, e error, status int) {
	if e == nil {
		return
	}
	b.whs(w, status)
	body, _ := json.Marshal(fmt.Sprintf("%s", e.Error()))
	w.Write(body)

	panic(
		Error{
			Status:   status,
			Location: WAI(2),
			error:    e,
		},
	)
}

func (b *base) WR(w http.ResponseWriter, status int, body interface{}) {
	bod, e := json.Marshal(body)
	if e != nil {
		b.WE(w, e, 500)
	}
	b.whs(w, status)
	w.Write(bod)
}

func (b *base) whs(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
}
