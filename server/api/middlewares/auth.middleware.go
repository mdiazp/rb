package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mdiazp/rb/server/api"
	dbhandlers "github.com/mdiazp/rb/server/db/handlers"
)

// AuthHeader ...
const AuthHeader = "AuthToken"

// MustAuth ...
func MustAuth(base api.Base, next http.Handler) http.Handler {
	return &Auth{
		next: next,
		Base: base,
	}
}

// Auth ...
type Auth struct {
	next http.Handler
	api.Base
}

func (c *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get(AuthHeader)
	claims, e := c.JWTHandler().GetClaims(tokenString)
	c.WE(w, e, 401)

	if !claims.VerifyExpiresAt(time.Now().Unix(), false) {
		c.WE(w, fmt.Errorf("Token expired"), 401)
	}

	user, e := c.DB().RetrieveUserByUsername(claims.Username)
	if e != nil {
		if e == dbhandlers.ErrRecordNotFound {
			c.WE(w, fmt.Errorf("User Not Found"), 401)
		}
		c.WE(w, e, 500)
	}

	//Check that user be actived
	if !user.Actived {
		c.WE(w, fmt.Errorf("User is not actived"), 401)
	}

	c.ContextWriteAuthor(r, user)

	c.next.ServeHTTP(w, r)
}
