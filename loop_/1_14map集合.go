package main

import (
	"fmt"
	"sync"
	"time"
)

func receiveMap(param map[string]int) {
	fmt.Println("before modify, in receiveMap func: param[\"a\"] = ", param["a"])
	param["a"] = 2
	param["b"] = 3
}
func main() {
	//var m1 map[string]string
	//fmt.Println("m1 length:", len(m1))
	//
	//m2 := make(map[string]string)
	//fmt.Println("m2 length:", len(m2))
	//fmt.Println("m2 =", m2)
	//
	//m3 := make(map[string]string, 10)
	//fmt.Println("m3 length:", len(m3))
	//fmt.Println("m3 =", m3)
	//
	//m4 := map[string]string{}
	//fmt.Println("m4 length:", len(m4))
	//fmt.Println("m4 =", m4)
	//
	//m5 := map[string]string{
	//	"key1": "value1",
	//	"key2": "value2",
	//}
	//fmt.Println("m5 length:", len(m5))
	//fmt.Println("m5 =", m5)
	//m := make(map[string]int, 10)
	//
	//m["1"] = int(1)
	//m["2"] = int(2)
	//m["3"] = int(3)
	//m["4"] = int(4)
	//m["5"] = int(5)
	//m["6"] = int(6)
	//
	//// 获取元素
	//value1 := m["1"]
	//fmt.Println("m[\"1\"] =", value1)
	//
	//value1, exist := m["1"]
	//fmt.Println("m[\"1\"] =", value1, ", exist =", exist)
	//
	//valueUnexist, exist := m["10"]
	//fmt.Println("m[\"10\"] =", valueUnexist, ", exist =", exist)
	//
	//// 修改值
	//fmt.Println("before modify, m[\"2\"] =", m["2"])
	//m["2"] = 20
	//fmt.Println("after modify, m[\"2\"] =", m["2"])
	//
	//// 获取map的长度
	//fmt.Println("before add, len(m) =", len(m))
	//m["10"] = 10
	//fmt.Println("after add, len(m) =", len(m))
	//
	//// 遍历map集合main
	//for key, value := range m {
	//	fmt.Println("iterate map, m[", key, "] =", value)
	//}
	//
	//// 使用内置函数删除指定的key
	//_, exist_10 := m["10"]
	//fmt.Println("before delete, exist 10: ", exist_10)
	//delete(m, "10")
	//_, exist_10 = m["10"]
	//fmt.Println("after delete, exist 10: ", exist_10)
	//
	//// 在遍历时，删除map中的key
	//for key := range m {
	//	fmt.Println("iterate map, will delete key:", key)
	//	delete(m, key)
	//}
	//fmt.Println("m = ", m)
	//m := make(map[string]int)
	//m["a"] = 1
	//receiveMap(m)
	//fmt.Println("m =", m)
	//2、并发的时候使用map集合
	//m := make(map[string]int)
	//
	//go func() {
	//	for {
	//		m["a"]++
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		m["a"]++
	//		fmt.Println(m["a"])
	//	}
	//}()
	//
	//select {
	//case <-time.After(time.Second * 5):
	//	fmt.Println("timeout, stopping")
	//}
	m := make(map[string]int)
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(2)

	go func() {
		for {
			lock.Lock()
			m["a"]++
			lock.Unlock()
		}
	}()

	go func() {
		for {
			lock.Lock()
			m["a"]++
			fmt.Println(m["a"])
			lock.Unlock()
		}
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("timeout, stopping")
	}
}
