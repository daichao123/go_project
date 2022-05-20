package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func NotRepeat(str string) int {
	return 0
}

// IsListDuplicated 返回false 则没有重复 返回true 则重复
func IsListDuplicated(list []string) bool {

	tmpMap := make(map[string]int)

	for _, value := range list {
		tmpMap[value] = 1
	}

	var keys []interface{}
	for k := range tmpMap {
		keys = append(keys, k)
	}
	if len(keys) != len(list) {
		return true
	}
	return false
}

type Cat struct {
	Name string
	Type string
	Tag  string
}

type Car struct {
	Name  string
	Price float64
	Tag   string
}

var dogCount int64
var fishCount int64
var catCount int64

func main() {

	//3个函数分别打印dog、fish、cat，要求每个函数都要起一个goroutine，按照dog、fish、cat顺序打印在屏幕上10次。

	//使用等待组保持主协程和子协程同步完成
	var wg sync.WaitGroup

	//申明三个Channel 使用空结构体的原因是占用的内存空间大小为0，同时对于空结构体的组合，占用空间大小也为0：
	dogCh := make(chan struct{}, 1)
	fishCh := make(chan struct{}, 1)
	catCh := make(chan struct{}, 1)

	wg.Add(3)
	//启用三个协程
	go dog(&wg, dogCh, fishCh)
	go fish(&wg, fishCh, catCh)
	go cat(&wg, catCh, dogCh)

	//向一个channel 写入数据
	dogCh <- struct{}{}

	//等待子协程完成任务
	wg.Wait()
	fmt.Println("main over")
}

func dog(wg *sync.WaitGroup, dogCh, fishCh chan struct{}) {
	for {
		if dogCount >= int64(100) {
			wg.Done()
			return
		}
		//输入channel 里面的数据
		//如果此channel 里面没有数据 则会一直阻塞
		<-dogCh
		//原子操作
		atomic.AddInt64(&dogCount, 1)
		fmt.Println("dog")
		fishCh <- struct{}{}
	}
}

func fish(wg *sync.WaitGroup, fishCh, catCh chan struct{}) {
	for {
		if fishCount >= 100 {
			wg.Done()
			return
		}
		<-fishCh
		atomic.AddInt64(&fishCount, 1)
		fmt.Println("fish")
		catCh <- struct{}{}
	}
}

func cat(wg *sync.WaitGroup, catCh, dogCh chan struct{}) {
	for {
		if fishCount >= 100 {
			wg.Done()
			return
		}
		<-catCh
		atomic.AddInt64(&catCount, 1)
		fmt.Println("cat")
		dogCh <- struct{}{}
	}
}
