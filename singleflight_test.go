package go_proc_cache_test

import (
	"fmt"
	"github.com/guaidashu/go_proc_cache"
	"testing"
	"time"
)

func TestGroup_Get(t *testing.T) {
	go_proc_cache.InitProcCache(100)
	g := go_proc_cache.ProcCache

	// data, ok := g.Get("test", func() (interface{}, error) {
	// 	return "我是测试", nil
	// }, time.Second*2)
	// fmt.Println("ok ==>", ok, " data ====>", data)

	for i := 0; i < 1000; i++ {
		go func() {
			data, ok := g.Get("test", func() (interface{}, error) {
				return "我是测试", nil
			}, time.Second*2)
			fmt.Println("ok ==>", ok, " data ====>", data)
			// data, ok = g.Get("test", func() (interface{}, error) {
			// 	fmt.Println("走到这里了吗")
			// 	return "1111111", nil
			// })
			// fmt.Println("ok ==>", ok, " data ====>", data)
		}()
	}

	data, ok := g.Get("test_empty", func() (interface{}, error) {
		fmt.Println("空数据测试")
		return nil, nil
	})
	fmt.Println("ok ==>", ok, " data ====>", data)

	data, ok = g.Get("test_empty", func() (interface{}, error) {
		fmt.Println("空数据测试")
		return "空数据测试第二次尝试有数据没", nil
	})
	fmt.Println("ok ==>", ok, " data ====>", data)

	time.Sleep(time.Second * 3)
}

func TestIsNil(t *testing.T) {
	data := "1212"
	testIsNil(data)

	var data2 *int
	testIsNil(data2)

	type dataInterface interface {
	}
	var data3 dataInterface
	testIsNil(data3)

	var data4 map[string]string
	testIsNil(data4)

	var data5 *bool
	testIsNil(data5)

	testIsNil(100)
}

func testIsNil(data interface{}) {
	fmt.Println(go_proc_cache.IsNil(data))
}

