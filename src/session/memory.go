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
    sid string
    timeAccessed time.Time
    value map[interface{}]interface{}
}

// 实现Provider接口，由Session接口的实现类型调用存取Session
// 主要实现的是和存储介质交互
type MemoryProvider struct {
    lock sync.Mutex
    sessions map[string]*list.Element //存session
    list *list.List //GC清理
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
