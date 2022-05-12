package main

import (
	"fmt"
	"time"
)

func main() {
	//1.定义一个管道，可放10个int类型的数据
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	//2.定义一个管道，可以放5个string类型的数据
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统的方法遍历管道时，如果不关闭会阻塞而导致deadlock

	//在实际的开发中，我们不好确定什么时候该关闭管道！！！！
	//这时可以用select方式可以解决
	//label:
	for {
		select {
		//管道不关闭不会deadlock，会自动到下一个case匹配
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%d\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Printf("取不到数据。。。\n")
			time.Sleep(time.Second)
			return
			//break label
		}
	}
}
