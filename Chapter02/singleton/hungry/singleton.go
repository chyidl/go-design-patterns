// this implementation is not thread safe
// 一个类只允许创建一个对象 这个类就是一个单例类
package hungry

// Singleton 饿汉式单例
type singleton struct {
}

// pointer as nil not struct as nil
var instance *singleton

// 饿汉在类加载时候，instance实例已经创建并初始化好
// fail-fast 设计原则(有问题及早暴露)
func init() {
	instance = new(singleton)
}

// 获取实例
func GetInstance() *singleton {
	return instance
}
