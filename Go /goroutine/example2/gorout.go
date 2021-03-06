package main

import (
	"fmt"
	"sync"
	"time"
)

var over = make(chan bool)

const MAXConCurrency = 3

//var sem = make(chan int, 4) //控制并发任务数
var sem = make(chan bool, MAXConCurrency) //控制并发任务数

var maxCount = 6

func Worker(i int) bool {

	sem <- true
	defer func() {
		<-sem
	}()

	// 模拟出错处理
	if i == 5 {
		return false
	}
	fmt.Printf("now:%v num:%v\n", time.Now().Format("04:05"), i)
	time.Sleep(1 * time.Second)
	return true
}

func main() {
	//wg := &sync.WaitGroup{}
	var wg sync.WaitGroup
	for i := 1; i <= maxCount; i++ {
		wg.Add(1)
		fmt.Printf("for num:%v\n", i)
		go func(i int) {
			defer wg.Done()
			for x := 1; x <= 3; x++ {
				if Worker(i) {
					break
				} else {
					fmt.Printf("retry :%v\n", x)
				}
			}
		}(i)
	}
	wg.Wait() //等待所有goroutine退出
}
