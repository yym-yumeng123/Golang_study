package main

import "time"

func ExpensiveFunction(key string) string {
	time.Sleep(2 * time.Second)
	return "result for:" + key
}

/**
 * @Doc 入参: f: func
 * 出参: func
 */
func Cache(f func(string) string) func(string) string {
	// 创建一个map
	cache := make(map[string]string)

	return func(key string) string {
		// cache 是否有 key
		if val, ok := cache[key]; ok {
			return val
		}
		val := f(key)
		cache[key] = val
		return val
	}
}

func main() {
	cache1 := Cache(ExpensiveFunction)
	println(cache1("key"))
}
