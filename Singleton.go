package main

import "sync"

var single *Singleton
var once sync.Once

type Singleton struct {
}

func GetSingletonLazy() *Singleton {
	once.Do(func() {
		single = &Singleton{}
	})
	return single
}
