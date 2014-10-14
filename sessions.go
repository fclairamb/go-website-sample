package main

import (
	"github.com/astaxie/beego/session"
	"net/http"
	"net/url"
	"time"
)

const ( // Cookie names
	COOKIE_SESSION_ID = "sid"
	COOKIE_NAME_USER  = "auth_user"
	COOKIE_NAME_AUTH  = "auth_pass"
)

const KEY_SESSION = "thy"

var sessionManager *session.Manager

func init() {
	sessionManager, _ = session.NewManager("memory", `{"cookieName":"`+COOKIE_SESSION_ID+`","gclifetime":36000}`)
	go sessionManager.GC()
}

func GetSession(w http.ResponseWriter, r *http.Request) *ThySession {
	sessionStore := sessionManager.SessionStart(w, r)

	var session *ThySession

	if value := sessionStore.Get(KEY_SESSION); value != nil {
		session = value.(*ThySession)
	} else {
		session = NewThySession(sessionStore)
		sessionStore.Set(KEY_SESSION, session)
		log.Debug("Creating new session: %v", session.ID())

		if err := session.CheckCookies(w, r); err != nil {
			log.Warning("Error: %v", err)
		}

	}

	return session
}

type ThySession struct {
	Session       session.SessionStore
	Authenticated bool
	User          string
}

func NewThySession(store session.SessionStore) *ThySession {
	this := &ThySession{Session: store}

	return this
}

func (this *ThySession) Get(key string) interface{} {
	return this.Session.Get(key)
}

func (this *ThySession) Set(key string, value interface{}) {
	this.Session.Set(key, value)
}

type User struct {
	User string
	Pass string
}

var users = []*User{
	{"florent", "pass1"},
	{"marie", "pass2"},
	{"charles", "pass3"},
	{"louis", "pass4"},
}

func (this *ThySession) Authenticate(username, password string) {
	// This is some dummy authentication
	for _, u := range users {
		if username == u.User && password == u.Pass {
			this.Authenticated = true
			this.User = username
			log.Info("User %s got authenticated !", username)
		}
	}
}

func (this *ThySession) CheckCookies(w http.ResponseWriter, r *http.Request) error {
	var user, auth string
	if cookie, err := r.Cookie(COOKIE_NAME_USER); err != nil {
		return err
	} else {
		user = cookie.Value
	}
	if cookie, err := r.Cookie(COOKIE_NAME_AUTH); err != nil {
		return err
	} else {
		auth = cookie.Value
	}

	log.Debug("user = \"%s\" ; auth = \"%s\"", user, auth)

	for _, u := range users {
		if u.User == user && HashSHA1(u.User+"_SALT") == auth {
			this.Authenticated = true
			this.User = user
			log.Info("User %s got authenticated !", user)
		}
	}
	return nil
}

func (this *ThySession) SaveCookies(w http.ResponseWriter, r *http.Request) {
	expires := time.Now().UTC().Add(time.Duration(time.Hour * 24 * 30 * 12))
	cookies := []*http.Cookie{
		&http.Cookie{
			Name:    COOKIE_NAME_USER,
			Value:   this.User,
			Expires: expires,
		},
		&http.Cookie{
			Name:    COOKIE_NAME_AUTH,
			Value:   HashSHA1(this.User + "_SALT"),
			Expires: expires,
		},
	}
	for _, c := range cookies {
		http.SetCookie(w, c)
	}
}

func (this *ThySession) ID() string {
	return this.Session.SessionID()
}

func (this *ThySession) AuthenticatedIfNotRedirect(w http.ResponseWriter, r *http.Request) bool {
	if !this.Authenticated {
		http.Redirect(w, r, "/login?source="+url.QueryEscape(r.RequestURI), http.StatusFound)
	}
	return this.Authenticated
}

func (this *ThySession) End(w http.ResponseWriter, r *http.Request) {
	this.Session.Delete(KEY_SESSION)
	this.Authenticated = false

	cookies := []*http.Cookie{
		&http.Cookie{
			Name:   COOKIE_NAME_USER,
			Value:  "",
			MaxAge: 0,
		},
		&http.Cookie{
			Name:   COOKIE_NAME_AUTH,
			Value:  "",
			MaxAge: 0,
		},
		&http.Cookie{
			Name:   COOKIE_SESSION_ID,
			Value:  "",
			MaxAge: 0,
		},
	}
	for _, c := range cookies {
		http.SetCookie(w, c)
	}
}

type PageBaseData struct {
	Title   string
	Session *ThySession
}
