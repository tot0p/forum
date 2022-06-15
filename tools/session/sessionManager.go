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

//Method to get the number of sessions on the website
func (s SessionManager) GetNBSession() int {
	return len(s.Provider.Sessions)
}

//Function to generate a new session manager
func NewManager(provideName, CookieName string, maxlifetime int64) (*SessionManager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("Session: Unknown Provide %q (Forgotten Import?)", provideName)
	}
	return &SessionManager{Provider: provider, CookieName: CookieName, Maxlifetime: maxlifetime}, nil
}

var provides = make(map[string]Provider)

//Function to register a session
func Register(name string, provider *Provider) error {
	if _, dup := provides[name]; dup {
		return fmt.Errorf("Session: Register Called Twice For Provider " + name)
	}
	provides[name] = *provider
	return nil
}

//Method to start a session
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

//Method to check if a session exist using a sid
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

//Method to delete a session using a sid
func (manager *SessionManager) SessionDestroy(SID string) error {
	manager.Provider.SessionDestroy(SID)
	return nil
}

//Method to generate a session id
func (manager *SessionManager) GenerateSID() string {
	return uuid.NewString()
}

//Method to time out a session
func (manager *SessionManager) Timout() {
	manager.Provider.SessionTimout(manager.Maxlifetime)
	time.AfterFunc(time.Duration(manager.Maxlifetime), func() { manager.Timout() })
}
