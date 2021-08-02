package main

import (
	"fmt"
	"time"
)

//线程，抢占式多任务处理，操作系统层面多任务
//协程，非抢占式多任务处理，由协程主动交出控制权，编译器/解释器/虚拟机 层面多任务， 多个协程可在一个线程上运行

func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Println("goroutine from ", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}

//func main() {
//	var a [10]int
//	for i := 0; i < 10; i++ {
//		go func(i int) {
//			for {
//				a[i]++
//				//runtime.Gosched()
//			}
//		}(i)
//	}
//	time.Sleep(time.Millisecond)
//	fmt.Println(a)
//}

//func main() {
//	var a [10]int
//	for i := 0; i < 10; i++ {
//		go func(i int) {
//			for {
//				a[i]++
//				//runtime.Gosched()
//			}
//		}(i)
//	}
//	time.Sleep(time.Millisecond)
//	fmt.Println(a)
//}
