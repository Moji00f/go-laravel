package session

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifeTime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	CookieSecure   string
	SessionType    string
}

func (c *Session) InitSession() *scs.SessionManager {
	var persist, secure bool

	// how log should sessions last?
	minutes, err := strconv.Atoi(c.CookieLifeTime)
	if err != nil {
		minutes = 60
	}

	// should cookie persists?
	if strings.ToLower(c.CookiePersist) == "true" {
		persist = true
	}

	// must coolies be secure?
	if strings.ToLower(c.CookieSecure) == "true" {
		secure = true
	} else {
		secure = false
	}

	// create session
	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = c.CookieName
	session.Cookie.Secure = secure
	session.Cookie.Domain = c.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	// which session store?
	switch strings.ToLower(c.SessionType) {
	case "redis":
	case "mysql", "mariadb":
	case "postgres", "postgresql":
	default:
		//cookie
	}

	return session
}
