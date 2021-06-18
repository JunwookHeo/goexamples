// https://github.com/konradreiche/concurrent-non-blocking-cache

package main

import (
	"fmt"
)

type Cache struct {
	requests chan request
}

type request struct {
	key      string
	response chan result
}

type entry struct {
	res   result
	ready chan struct{}
}

type result struct {
	value []byte
	err   error
}

type Func func(key string) ([]byte, error)

func NewCache(f Func) *Cache {
	cache := &Cache{requests: make(chan request)}
	go cache.server(f)
	return cache
}

func (c *Cache) Get(key string) ([]byte, error) {
	response := make(chan result)
	c.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (c *Cache) server(f Func) {
	cache := make(map[string]*entry)
	for req := range c.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}

func func1(str string) func() string {
	return func() string {
		str2 := str + "_something"
		return str2
	}
}
func main() {
	f := func1("test2")
	str := f()
	fmt.Println(str)
}
