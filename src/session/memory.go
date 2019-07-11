package session

import (
	"container/list"
	"sync"
	"time"
)

// 创建一个内存Session存储实例
var memoryProvider = &MemoryProvider{list: list.New()}

// 实现Session接口，被Web交互接口调用
type SessionStore struct {
	sid          string
	timeAccessed time.Time
	value        map[interface{}]interface{}
}

func (ss *SessionStore) Set(key, value interface{}) error {
	ss.value[key] = value
	memoryProvider.SessionUpdate(ss.sid)
	return nil
}

func (ss *SessionStore) Get(key interface{}) interface{} {
	memoryProvider.SessionUpdate(ss.sid)
	if v, ok := ss.value[key]; ok {
		return v
	} else {
		return nil
	}
}

func (ss *SessionStore) Delete(key interface{}) error {
	delete(ss.value, ss.sid)
	memoryProvider.SessionUpdate(ss.sid)
	return nil
}

func (ss *SessionStore) SessionID() string {
	return ss.sid
}

// 实现Provider接口，由Session接口的实现类型调用存取Session
// 主要实现的是和存储介质交互
type MemoryProvider struct {
	lock     sync.Mutex
	sessions map[string]*list.Element //存session
	list     *list.List               //GC清理
}

// 初始化session实例
// 存储到内存
func (p *MemoryProvider) SessionInit(sid string) (Session, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	v := make(map[interface{}]interface{}, 0)
	session := &SessionStore{sid, time.Now(), v}
	element := p.list.PushBack(session)
	p.sessions[sid] = element
	return session, nil
}

// 读取session
// 从内存中读取， 有则返回， 无则SessionInit
func (p *MemoryProvider) SessionRead(sid string) (Session, error) {
	if element, ok := p.sessions[sid]; ok {
		return element.Value.(Session), nil
	} else {
		session, err := p.SessionInit(sid)
		return session, err
	}
}

// 检查内存中是否存在session
// 存在则从map中删除，再从list中删除
func (p *MemoryProvider) SessionDestroy(sid string) error {
	if element, ok := p.sessions[sid]; ok {
		delete(p.sessions, sid)
		p.list.Remove(element)
	}

	return nil
}

// 检查内存中session的最近访问时间+maxLifeTime小于当前时间，回收内存
func (p *MemoryProvider) SessionGC(maxLifeTime int64) {
	p.lock.Lock()
	defer p.lock.Unlock()

	for {
		element := p.list.Back() // 读取list的最后一个元素
		if element == nil {
			return
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxLifeTime) < time.Now().Unix() {
			p.list.Remove(element)
			delete(p.sessions, element.Value.(*SessionStore).SessionID())
		}
	}
}

// 检查session是否在内存，刷新timeAccessed为当前时间，移动到list的最前段
func (p *MemoryProvider) SessionUpdate(sid string) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if element, ok := p.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		p.list.MoveToFront(element)
	}

	return nil
}
