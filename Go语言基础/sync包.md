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