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
