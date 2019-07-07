package session

// Session实例的操作接口定义
type Session interface {
	// interface{} 是一个空接口，任意类型实现了空接口，感觉类似与Java的泛型
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}
