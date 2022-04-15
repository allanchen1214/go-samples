package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func Counter(wg *sync.WaitGroup) {
	var counter int
	for i := 0; i < 300; i++ {
		time.Sleep(time.Second * 1)
		counter++
	}
	wg.Done()
}

func main() {
	flag.Parse()

	//远程获取pprof数据
	go func() {
		log.Println(http.ListenAndServe(":8088", nil))
	}()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Counter(&wg)
	}
	wg.Wait()

	// sleep 1min, 在程序退出之前可以查看性能参数.
	time.Sleep(60 * time.Second)
}
