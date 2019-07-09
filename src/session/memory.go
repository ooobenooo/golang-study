package session

import (
    "container/list"
    "sync"
    "time"
)

// 实现Session接口，被Web交互接口调用
type SessionStore struct {
    sid string,
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