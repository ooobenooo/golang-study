package session

import "sync"

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
        panic(name, "is exists. do not allow register again.")
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

func (manager *Manager) sessionId() string{
    b := make([]byte, 32)
    if _, err := io.ReadFull(rand.Reader, b); err != nil {
        return ""
    }
    
    return base64.URLEncoding.EncodeToString(b)
}