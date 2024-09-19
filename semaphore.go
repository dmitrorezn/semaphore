package semaphore

type semaphore chan struct{}

func (s *semaphore) acquire() {
	*s <- struct{}{}
}

func (s *semaphore) release() {
	<-*s
}
