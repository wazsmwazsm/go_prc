package mysync

type Empty interface {}

/* signal-wait 信号量 */
type Semaphore chan Empty

func NewSemaphore(n int) Semaphore {
	return make(Semaphore, n)
}

// wait
func (s Semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}
// 释放
func (s Semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<- s
	}
}

// 等待 wait 一个信号量（P操作/上锁）：该操作对信号量值进行测试，
// 若其值小于等于0，则等待阻塞。一旦大于 0 则将其减 1
func (s Semaphore) Wait(n int) {
	s.V(n)
}
 
// 将一个信号量加 1，并唤醒等待该信号量的任意 go 程，以此挂出信号量
func (s Semaphore) Signal() {
	s.P(1)
}


/* mutexes 互斥锁 */
type Mutex struct {
	se Semaphore
}

func NewMutex() *Mutex {
	mutex := new(Mutex)
	mutex.se = make(Semaphore, 1)
	return mutex
}

// 向 channel 写一个数据, semaphore 长度为 1 时, 一个 go 程写入后另一个
// go 程就阻塞了
func (m *Mutex) Lock() {
	m.se <- new(Empty)
}
// 从 channel 读一个数据,
func (m *Mutex) Unlock() {
	<- m.se
}
