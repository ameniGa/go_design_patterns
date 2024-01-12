package singleton

import (
	"log"
	"sync"
)

type Singleton interface {
	AddOne() int
}

// should be unexposed
type singleton struct {
	count int
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

// should be unexposed
var instance *singleton

var once sync.Once

func GetInstance() Singleton {
	if instance == nil {
		log.Printf("instance not created yet")
		once.Do(
			func() {
				log.Printf("create instance once")
				instance = &singleton{}
			})
	}
	return instance
}
