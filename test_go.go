package main

import (
	"fmt"
	"math/rand"
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

func main() {
	NotRepeat("abcabcdbb")
	var computerFist, userFist string
	computerNum := rand.New(rand.NewSource(123)).Intn(3)
	switch computerNum {
	case 0:
		computerFist = "石头"
	case 1:
		computerFist = "石头"
	case 2:
		computerFist = "布"
	}
	fmt.Println("请出招:")
	fmt.Scan(&userFist)

	switch userFist {
	case "石头":
		if computerFist == "石头" {
			fmt.Println()
		} else if computerFist == "布" {

		} else if computerFist == "剪刀" {

		}
	case "剪刀":
	case "布":
	default:
		fmt.Println("正确的招式")
	}

}
