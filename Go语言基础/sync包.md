### sync.WaitGroup

不要进行拷贝

### sync.Once 执行一次

### sync.Mutex 锁



### sync.Cond 等待某一个条件完成

### sync.Pool  对象池, 频繁创建和释放某一个结构

**注意:** `sync.pool` 并不保证池中的对象一直可用. 它可以在任何时候清空池, 或者从池中移出对象.
当从池中获得对象时, 需要在使用之前将对象状态重置为默认值, 确保对象处于正确的状态

### sync.Map 并发安全哈希表

### sync.RWMutex 锁读写分离

```go
func RwMutex() {
	// 加读锁
	rwMutext.RLock()
	defer rwMutext.RUnlock()
	// 加写锁
    rwMutext.Lock()
    rwMutext.Unlock()
	
}
```


- 尽量用 sync.RWMutex
- sync.Once 可以保证代码只执行一次, 一般用于解决初始化
- sync.WaitGroup 能用来在多个 goroutine 之间进行同步