package memory

// 实现Session接口，被Web交互接口调用
type SessionStore struct {
    
}

// 实现Provider接口，由Session接口的实现类型调用存取Session
// 主要实现的是和存储介质交互
type MemoryProvider struct {
    
}