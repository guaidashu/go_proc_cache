package go_proc_cache

import (
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

// var ProcCache *cache.Cache
var ProcCache *Group

// 初始化进程内缓存防击透工具
// goroutinePoolNumber 协程池数量
func InitProcCache(goroutinePoolNumber int) {
	InitGoSafePool(goroutinePoolNumber)
	// ProcCache = cache.New(5*time.Minute, 10*time.Minute)
	ProcCache = &Group{
		cache: cache.New(5*time.Minute, 10*time.Minute),
		lock:  sync.Mutex{},
		data:  make(map[string]*cacheData),
	}
}
