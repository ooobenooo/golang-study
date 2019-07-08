package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var GlobalSessions *Manager

func init() {
	GlobalSessions, _ = NewManager("memory", "gSessionId", 3600)
}

// go没有标准Session，需自己实现或用第三方库
// 定义一个Session管理器的数据结构
type Manager struct {
	cookieName  string     // session的名称，传递给客户端
	lock        sync.Mutex // 锁，多线程并发保护session manager
	provider    Provider   // Session可以存储在任意计算机的存储介质，生命周期抽象管理接口
	maxLifeTime int64
}

var providers = make(map[string]Provider)

// 注册Provider
func RegisterProvider(name string, provider Provider) {
	if _, exist := providers[name]; exist {
		panic("provider is exists. do not allow register again.")
	}
	providers[name] = provider
}

// 创建新的Session Manager，调用方维护一个Manager的全局变量实例
func NewManager(providerName string, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provider %q", providerName)
	}

	return &Manager{cookieName: cookieName, provider: provider, maxLifeTime: maxLifeTime}, nil
}

// 生产sessionId
func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)
}

// 这个函数用于检查访问用户是否已经有session,没有则分配一个session给用户
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName) // 从请求中获取cookie

	// 如果从请求中没有获得cookie, 表示是一个新的请求，分配session
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()                     // 随机生产sessionId
		session, _ = manager.provider.SessionInit(sid) // 利用provider创建session对象
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid),
			Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
		http.SetCookie(w, &cookie) // 设置cookie到response中，由浏览器保存
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid) //读取session
	}

	return session
}

// 销毁session并设置cookie过期
func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()

		// 销毁session
		manager.provider.SessionDestroy(cookie.Value)
		// cookie设置过期
		expiration := time.Now()

		// go 技巧，cookie变量之前已经被声明，在这里可以重新被声明，在
		// 一个作用域中变量名可以复用
		var cookie = http.Cookie{Name: manager.cookieName, Path: "/",
			HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}
