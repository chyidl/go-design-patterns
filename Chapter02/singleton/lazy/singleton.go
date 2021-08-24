// this implementation is not thread safe
// 一个类只允许创建一个对象 这个类就是一个单例类
package lazy

import "sync"

// Singleton 懒汉式单例
type singleton struct {
}

// pointer as nil not struct as nil
var instance *singleton
var once = &sync.Once{}

// 获取实例
func GetInstance() *singleton {
	if instance == nil {
		// done uint32
		// m Mutex
		once.Do(func() {
			instance = new(singleton)
		})
	}
	return instance
}
