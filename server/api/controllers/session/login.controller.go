package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/mdiazp/rb/server/api"
	"github.com/mdiazp/rb/server/api/controllers"
	dbhandlers "github.com/mdiazp/rb/server/db/handlers"
	"github.com/mdiazp/rb/server/db/models"
)

// LoginController ...
type LoginController interface {
	controllers.BaseController
}

// NewLoginController ...
func NewLoginController(base api.Base) LoginController {
	return &loginController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type loginController struct {
	api.Base
}

func (c *loginController) GetRoute() string {
	return "/session"
}

func (c *loginController) GetMethods() []string {
	return []string{"POST"}
}

// GetAccess ...
func (c *loginController) GetAccess() controllers.Permission {
	return ""
}

// ServeHTTP ...
func (c *loginController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var cred Credentials
	c.ReadJSON(w, r, &cred)

	// Authenticate
	provider := c.GetAuthProvider((api.AuthProvider)(cred.Provider))
	if provider == nil {
		c.WE(w, fmt.Errorf("Unknowed Provider: %s", cred.Provider), 401)
	}

	e := provider.Authenticate(cred.Username, cred.Password)
	if e != nil {
		c.WE(w, fmt.Errorf("Invalid Credentials"), 401)
	}

	//Check User be registered
	user, e := c.DB().RetrieveUserByUsername(cred.Username)
	if e != nil {
		if e == dbhandlers.ErrRecordNotFound {
			c.WE(w, fmt.Errorf("User is not registered"), 401)
		}
		c.WE(w, e, 500)
	}
	if user.Provider != cred.Provider {
		c.WE(w, fmt.Errorf("Incorrect Provider"), 401)
	}

	//Check that user be actived
	if !user.Actived {
		c.WE(w, fmt.Errorf("User is not actived"), 401)
	}

	// Set Claims
	claims := api.Claims{
		Username: cred.Username,
		Provider: cred.Provider,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}

	token, e := c.JWTHandler().GetToken(claims)
	c.WE(w, e, 500)

	// GroupPermissions Loading
	ps := controllers.GetPermissions((controllers.Rol)(user.Rol))

	session := Session{
		User:        *user,
		Token:       token,
		Permissions: ps,
	}
	c.WR(w, 200, session)
}

// Session ...
type Session struct {
	User        models.User
	Token       string
	Permissions controllers.Permissions
}

// Credentials ...
type Credentials struct {
	Username string
	Password string
	Provider string
}
