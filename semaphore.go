package semaphore

type Semaphore chan struct{}

func (s *Semaphore) acquire() {
	*s <- struct{}{}
}

func (s *Semaphore) release() {
	<-*s
}
