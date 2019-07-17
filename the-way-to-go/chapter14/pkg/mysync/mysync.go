package mysync

type Empty interface {}

type semaphore chan Empty

// 互斥锁变量
var Mutex semaphore

func NewSemaphore(n int) semaphore {
	return make(semaphore, n)
}

func NewMutex() semaphore {
	return NewSemaphore(1)
}

func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<- s
	}
}

/* mutexes 实现互斥锁 */
// 向 channel 写一个数据, semaphore 长度为 1 时, 一个 go 程写入后另一个
// go 程就阻塞了
func (s semaphore) Lock() {
	s.P(1)
}
// 从 channel 读一个数据,
func (s semaphore) Unlock() {
	s.V(1)
}

/* signal-wait 实现信号量 */
// 等待 wait 一个信号量（P操作/上锁）：该操作对信号量值进行测试，
// 若其值小于等于0，则等待阻塞。一旦大于 0 则将其减 1
func (s semaphore) Wait(n int) {
	s.P(n)
}
 
// 将一个信号量加 1，并唤醒等待该信号量的任意 go 程，以此挂出信号量
func (s semaphore) Signal() {
	s.V(1)
}

func init() {
	Mutex = NewMutex()
}
