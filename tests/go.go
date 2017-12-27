package main

import (
	"fmt"
	"sync"
)

// comment

var mutex sync.RWMutex
var mappy = map[string]string{}

func read(wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Println("read:", mappy["hi"])
	wg.Done()
	mutex.Unlock()
}

func write(wg *sync.WaitGroup, v string) {
	mutex.Lock()
	mappy["hi"] = v
	fmt.Println("wrote:", mappy["hi"])
	wg.Done()
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup

	mappy["hi"] = "hello"

	wg.Add(7)
	go read(&wg)
	go read(&wg)
	go write(&wg, "five")
	go read(&wg)
	go read(&wg)
	go read(&wg)
	go write(&wg, "llo")

	wg.Wait()
}
