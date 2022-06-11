package session

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type Provider struct {
	lock     sync.Mutex               // lock
	Sessions map[string]*list.Element // save in memory
	list     *list.List               // gc
}

func (pder *Provider) SessionInit(sid string) (*SessionStore, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := pder.list.PushBack(newsess)
	if _, ok := pder.Sessions[sid]; ok {
		return nil, fmt.Errorf("Session with SID %s Alreay Exist", sid)
	}
	pder.Sessions[sid] = element
	return newsess, nil
}

func (pder *Provider) SessionRead(sid string) (*SessionStore, error) {
	if element, ok := pder.Sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	}
	return nil, fmt.Errorf("Invalid Session With SID %s", sid)
}

func (pder *Provider) SessionDestroy(sid string) error {
	if element, ok := pder.Sessions[sid]; ok {
		delete(pder.Sessions, sid)
		pder.list.Remove(element)
		return nil
	}
	return fmt.Errorf("Invalid Session With SID %s", sid)
}

func (pder *Provider) SessionTimout(maxlifetime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	for {
		element := pder.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.Sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.Sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return fmt.Errorf("Invalid Session With SID %s", sid)
}
