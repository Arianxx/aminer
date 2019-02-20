package main

type sem struct {
	c chan bool
}

func getSem(concurrency int) *sem {
	return &sem{make(chan bool, concurrency)}
}

func (s *sem) acquire() {
	s.c <- true
}

func (s *sem) release() {
	<-s.c
}

func (s *sem) wait() {
	for i := 0; i < cap(s.c); i++ {
		s.acquire()
	}
}
