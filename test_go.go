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
	var wg sync.WaitGroup

	dogCh := make(chan struct{}, 1)
	fishCh := make(chan struct{}, 1)
	catCh := make(chan struct{}, 1)

	wg.Add(3)
	go dog(&wg, dogCh, fishCh)
	go fish(&wg, fishCh, catCh)
	go cat(&wg, catCh, dogCh)

	dogCh <- struct{}{}

	wg.Wait()
	fmt.Println("main over")
	//NotRepeat("abcabcdbb")
	//var computerFist, userFist string
	//computerNum := rand.New(rand.NewSource(123)).Intn(3)
	//switch computerNum {
	//case 0:
	//	computerFist = "石头"
	//case 1:
	//	computerFist = "石头"
	//case 2:
	//	computerFist = "布"
	//}
	//fmt.Println("请出招:")
	//fmt.Scan(&userFist)
	//
	//switch userFist {
	//case "石头":
	//	if computerFist == "石头" {
	//		fmt.Println()
	//	} else if computerFist == "布" {
	//
	//	} else if computerFist == "剪刀" {
	//
	//	}
	//case "剪刀":
	//	if computerFist == "石头" {
	//		fmt.Println()
	//	} else if computerFist == "布" {
	//
	//	} else if computerFist == "剪刀" {
	//
	//	}
	//case "布":
	//	if computerFist == "石头" {
	//		fmt.Println()
	//	} else if computerFist == "布" {
	//
	//	} else if computerFist == "剪刀" {
	//
	//	}
	//default:
	//	fmt.Println("正确的招式")
	//}

}

func dog(wg *sync.WaitGroup, dogCh, fishCh chan struct{}) {
	for {
		if dogCount >= int64(100) {
			wg.Done()
			return
		}
		<-dogCh
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
