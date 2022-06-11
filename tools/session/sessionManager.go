package session

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
)

var GlobalSessions *SessionManager

type SessionManager struct {
	CookieName  string //private cookiename
	Provider    Provider
	Maxlifetime int64
}

func (s SessionManager) GetNBSession() int {
	return len(s.Provider.Sessions)
}

func NewManager(provideName, CookieName string, maxlifetime int64) (*SessionManager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("Session: Unknown Provide %q (Forgotten Import?)", provideName)
	}
	return &SessionManager{Provider: provider, CookieName: CookieName, Maxlifetime: maxlifetime}, nil
}

var provides = make(map[string]Provider)

func Register(name string, provider *Provider) error {
	if _, dup := provides[name]; dup {
		return fmt.Errorf("Session: Register Called Twice For Provider " + name)
	}
	provides[name] = *provider
	return nil
}

func (manager *SessionManager) SessionStart(w http.ResponseWriter, r *http.Request) (session *SessionStore) {
	cookie, err := r.Cookie(manager.CookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.GenerateSID()
		session, _ = manager.Provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.CookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.Maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.Provider.SessionRead(sid)
	}
	return
}

func (manager *SessionManager) SessionExist(SID string) bool {
	if SID == "" {
		return false
	} else {
		if _, ok := manager.Provider.Sessions[SID]; ok {
			return true
		}
		return false
	}
}

func (manager *SessionManager) SessionDestroy(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(manager.CookieName)
	if err != nil || cookie.Value == "" {
		return fmt.Errorf("Cant Delete A Nil Session")
	} else {
		manager.Provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name: manager.CookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
	return nil
}

func (manager *SessionManager) GenerateSID() string {
	return uuid.NewString()
}

func (manager *SessionManager) Timout() {
	manager.Provider.SessionTimout(manager.Maxlifetime)
	time.AfterFunc(time.Duration(manager.Maxlifetime), func() { manager.Timout() })
}
