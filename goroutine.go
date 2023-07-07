package go_proc_cache

import (
	"github.com/panjf2000/ants/v2"
	"log"
)

var routinePool *ants.PoolWithFunc

// 初始化安全协程池
func InitGoSafePool(poolNumber int) {
	routinePool, _ = ants.NewPoolWithFunc(poolNumber, runSafe, ants.WithNonblocking(true))
	for i := 0; i < poolNumber; i++ {
		_ = routinePool.Invoke(func() {})
	}
}

// 使用另外一个协程运行函数fn，当fn函数panic后会被recovers
func GoSafe(fn func()) {
	err := routinePool.Invoke(fn)
	if err != nil {
		log.Println("启动协程任务失败：", err)
	}
}

func runSafe(fn interface{}) {
	defer func() {
		if e := recover(); e != nil {
			log.Println("goroutine panic: ", e)
		}
	}()

	fn.(func())()
}

// 并行执行多个函数，当这些函数某些返回error，Mr函数会随机返回其中一个error
func Mr(fns ...func() error) error {
	var (
		funcLen    = len(fns)
		resultChan = make(chan error, funcLen)
		ch         = make(chan error)
		errCount   int
	)

	for _, fn := range fns {
		GoSafe(func() {
			resultChan <- fn()
		})
	}

	GoSafe(func() {
		var isReturn bool
		for err := range resultChan {
			errCount++
			if errCount >= funcLen {
				if !isReturn {
					ch <- err
					close(ch)
				}
				break
			}

			if !isReturn && err != nil {
				isReturn = true
				ch <- err
				close(ch)
				continue
			}
		}
		close(resultChan)
	})

	return <-ch
}
