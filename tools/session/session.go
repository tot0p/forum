package session

import (
	"container/list"
	"fmt"
	"log"
	"time"
)

var pder = &Provider{list: list.New()}

type SessionStore struct {
	sid          string                      // unique session id
	timeAccessed time.Time                   // last access time
	value        map[interface{}]interface{} // session value stored inside
}

//Method to set a session
func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	err := pder.SessionUpdate(st.sid)
	if err != nil {
		return fmt.Errorf("Invalid Session With SID %s", st.sid)
	}
	return nil
}

//Method to get a session
func (st *SessionStore) Get(key interface{}) (interface{}, error) {
	r := pder.SessionUpdate(st.sid)
	if r != nil {
		return nil, fmt.Errorf("Invalid Session With SID %s", st.sid)
	}
	if v, ok := st.value[key]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("Getter Key \"%s\" Does Not Exist", key)
}

//Method to delete a session
func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	err := pder.SessionUpdate(st.sid)
	if err != nil {
		return fmt.Errorf("Invalid Session with SID %s", st.sid)
	}
	return nil
}

//Method which return the session id
func (st *SessionStore) SessionID() string {
	fmt.Println(st.sid)
	return st.sid
}

//Function to init a session
func init() {
	pder.Sessions = make(map[string]*list.Element, 0)
	Register("memory", pder)
	var err error
	GlobalSessions, err = NewManager("memory", "SID", 6000)
	if err != nil {
		log.Fatal(err)
	}
	go GlobalSessions.Timout()
}
