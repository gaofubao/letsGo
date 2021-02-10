package synctest

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/petermattis/goid"
)

// 同步原语（或并发原语，是解决并发问题的基础数据结构）的适用场景：
// - 共享资源 解决并发读写资源时的data race问题、Mutex RWMutex
// - 任务编排 解决goroutine之间相互等待或依赖的顺序关系、WaitGroup Channel
// - 消息传递 解决goroutine之间的线程安全的数据交流、Channel

// 互斥锁（排它锁） mutex
// 使用互斥锁来限定临界区只能同时由一个线程所持有
// Mutex和RWMutex实现了Locker接口

func TestMutex1(t *testing.T) {
	var mu sync.Mutex
	var count = 0
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++			// count++不是原子操作
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

// 结构体嵌套
type Counter2 struct {
	sync.Mutex
	Count uint64
}

func TestMutex2(t *testing.T) {
	var counter Counter2
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.Lock()
				counter.Count++
				counter.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count)
}

// 封装锁
type Counter3 struct {
	CounterType int
	Name 		string
	mu 			sync.Mutex
	count 		uint64
}

func (c *Counter3) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter3) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func TestMutex3(t *testing.T) {
	var counter Counter3
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count())
}

// Mutex架构演进
// 1. 初版互斥锁，使用一个flag字段标记是否持有锁
// CAS
// 2. 新的goroutine也能有机会竞争锁
// 3. 新来的和被唤醒的有更多的机会竞争锁
// 4. 解决饥饿问题

// 谁申请，谁释放
type Foo struct {
	mu	  sync.Mutex
	count int
}

func (f *Foo) Bar() {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.count < 1000 {
		f.count += 3
		return
	}

	f.count++
	return
}

// Mutex常见的错误场景
// 1. Lock和Unlock不是成对出现

// 2. Copy已使用的Mutex
// sync包中的同步原语在使用后不能复制，因为Mutex是一个有状态的对象，它的state字段记录了这个锁的状态
type Counter4 struct {
	sync.Mutex
	Count int
}

func foo4(c Counter4) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}

func TestMutex4(t *testing.T)  {
	var c Counter4
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo4(c)		// 复制锁，导致死锁，可以使用go vet xxx.go检查
}

// 3. 重入
// Mutex不是可重入锁（递归锁），因为Mutex没有记录哪个goroutine拥有这把锁
func bar5(l sync.Locker)  {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

func foo5(l sync.Locker)  {
	fmt.Println("in foo")
	l.Lock()
	bar5(l)
	l.Unlock()
}

func TestMutex5(t *testing.T) {
	l := &sync.Mutex{}
	foo5(l)
}

// 如何实现一个可重入锁？
// 1) goroutine id
// 获取goroutine id
// 方式一：简单方式
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine"))[0]
	id ,err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

// 方式二：hacker方式 petermattis/goid
type RecursiveMutex struct {
	sync.Mutex
	owner     int64
	recursion int32
}

func (m *RecursiveMutex) Lock()  {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}

func (m *RecursiveMutex) Unlock()  {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, gid))
	}
	m.recursion--
	if m.recursion != 0 {
		return
	}
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}

// 2) token
type TokenRecursiveMutex struct {
	sync.Mutex
	token     int64
	recursion int32
}

func (m *TokenRecursiveMutex) Lock(token int64)  {
	if atomic.LoadInt64(&m.token) == token {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	atomic.StoreInt64(&m.token, token)
	m.recursion = 1
}

func (m *TokenRecursiveMutex) Unlock(token int64)  {
	if atomic.LoadInt64(&m.token) != token {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.token, token))
	}
	m.recursion--
	if m.recursion != 0 {
		return
	}
	atomic.StoreInt64(&m.token, 0)
	m.Mutex.Unlock()
}

// 4. 死锁
// 死锁产生的必要条件：
// 1) 互斥
// 2) 持有和等待
// 3) 不可剥夺
// 4) 环路等待


// 扩展Mutex功能
// TryLock
// 获取等待者的数量等指标
// 使用Mutex实现一个线程安全的队列
