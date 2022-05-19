package main

import (
	"fmt"
	"reflect"
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

func main() {
	goods := make([]interface{}, 0)
	goods = append(goods, Cat{Name: "cat test", Type: "type test", Tag: "tag test"})
	goods = append(goods, Car{Name: "car test", Price: 2300, Tag: "car tag test"})
	fmt.Printf("goods : %s", goods)
	for _, item := range goods {
		//fmt.Print(reflect.ValueOf(item))
		s := reflect.ValueOf(&item)
		typeOfT := s.Type()
		fmt.Printf("typeOFT:%s", typeOfT)
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			fmt.Printf("%d: %s %s = %v\n", i,
				typeOfT.Field(i).Name, f.Type(), f.Interface())
		}
		//fmt.Printf(itemValue.Type())
		//name := itemValue.FieldByName("Name")
		//fmt.Print(name.Interface().(string))
		//fmt.Printf(item.(string))
		//fmt.Println("item", item.Name)
	}

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
